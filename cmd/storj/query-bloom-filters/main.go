// usage: query-bloom-filters [-q] path/to/bloom/filters/… path/to/a/list.csv
//
// list.csv is
//
// nodeID1,pieceID1
// nodeID1,pieceID2
// nodeID2,pieceID3
// …

package main

import (
	"archive/zip"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"storj.io/common/storj"
	"storj.io/storj/satellite/gc/sender"
	"storj.io/storj/satellite/internalpb"
	"storj.io/storj/shared/bloomfilter"
)

func main() {
	quiet := flag.Bool("q", false, "quiet (only report if a piece is missing)")
	flag.Parse()

	if n := flag.NArg(); n != 2 {
		log.Fatalf("not enough/too many args (has %d, needs 2)", n)
	}

	nodesPieces, err := loadNodesPieces(flag.Arg(1))
	if err != nil {
		log.Panicf("loadPieceIDs: %v", err)
	}

	bloomFiltersDir := flag.Arg(0)

	entries, err := os.ReadDir(bloomFiltersDir)
	if err != nil {
		log.Panicf("ReadDir: %v", err)
	}

	retainInfoLookup := make(map[storj.NodeID]*internalpb.RetainInfo)

	for i, e := range entries {
		if e.IsDir() {
			continue
		}

		fileInfo, err := e.Info()
		if err != nil {
			log.Panicf("Info: %v", err)
		}

		path := bloomFiltersDir + fileInfo.Name()

		retainInfos, err := loadRetainInfos(path, fileInfo.Size())
		if err != nil {
			log.Panicf("loadRetainInfos: %v", err)
		}

		for _, info := range retainInfos {
			if _, ok := retainInfoLookup[info.StorageNodeId]; ok {
				log.Panicf("duplicate RetainInfo for %s", info.StorageNodeId)
			}
			if _, ok := nodesPieces[info.StorageNodeId]; ok {
				retainInfoLookup[info.StorageNodeId] = info
			}
		}

		log.Printf("%.0f%%: loaded %s", float32(i+1)/float32(len(entries))*100, path)
	}

	for nid, pieces := range nodesPieces {
		for _, pid := range pieces {
			checkFilter(retainInfoLookup, nid, pid, *quiet)
		}
	}
}

func checkFilter(lookup map[storj.NodeID]*internalpb.RetainInfo, nid storj.NodeID, pid storj.PieceID, quiet bool) {
	info, ok := lookup[nid]
	if !ok {
		log.Printf("bloom filter for %s not found", nid)
		return
	}

	f, err := bloomfilter.NewFromBytes(info.Filter)
	if err != nil {
		log.Printf("bloom filter for %s cannot be loaded: %v", nid, err)
		return
	}

	if f.Contains(pid) {
		if !quiet {
			log.Printf("bloom filter (fill=%.2f, size=%d) for %s (piece count=%d) contains checked piece", f.FillRate(), f.Size(), nid, info.PieceCount)
		}
		return
	}
	log.Printf("bloom filter for %s doesn't have %s", nid, pid)
}

func loadRetainInfos(path string, size int64) ([]*internalpb.RetainInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r, err := zip.NewReader(f, size)
	if err != nil {
		return nil, err
	}

	var infos []*internalpb.RetainInfo
	for _, rf := range r.File {
		i, err := sender.UnpackZipEntry(rf)
		if err != nil {
			return nil, fmt.Errorf("couldn't unpack %q: %w", rf.Name, err)
		}
		infos = append(infos, i)
	}
	return infos, nil
}

func loadNodesPieces(path string) (map[storj.NodeID][]storj.PieceID, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	ret := make(map[storj.NodeID][]storj.PieceID)
	for i := 0; ; i++ {
		records, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		nid, err := storj.NodeIDFromString(records[0])
		if err != nil {
			return nil, fmt.Errorf("couldn't decode %q: %w", nid, err)
		}
		pid, err := storj.PieceIDFromString(records[1])
		if err != nil {
			return nil, fmt.Errorf("couldn't decode %q: %w", pid, err)
		}
		ret[nid] = append(ret[nid], pid)
	}

	return ret, nil
}
