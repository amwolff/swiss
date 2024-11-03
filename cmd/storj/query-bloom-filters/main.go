// usage: query-bloom-filters [-l=path/to/pieces/to/check] [-q] path/to/bloom/filters/…

package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"storj.io/common/storj"
	"storj.io/storj/satellite/gc/sender"
	"storj.io/storj/satellite/internalpb"
	"storj.io/storj/shared/bloomfilter"
)

func main() {
	list := flag.String("l", "", "path to pieces to check")
	quiet := flag.Bool("q", false, "quiet (only report if a piece is missing)")
	flag.Parse()

	if *list == "" && *quiet {
		log.Fatalf("quiet can only be used with list")
	}
	if n := flag.NArg(); n != 1 {
		log.Fatalf("not enough/too many args (has %d, needs 1)", n)
	}

	pieceIDs, err := loadPieceIDs(*list)
	if err != nil {
		log.Panicf("loadPieceIDs: %v", err)
	}

	bloomFiltersDir := flag.Arg(0)

	entries, err := os.ReadDir(bloomFiltersDir)
	if err != nil {
		log.Panicf("ReadDir: %v", err)
	}

	var retainInfos []*internalpb.RetainInfo
	for i, e := range entries {
		if e.IsDir() {
			continue
		}

		fileInfo, err := e.Info()
		if err != nil {
			log.Panicf("Info: %v", err)
		}

		path := bloomFiltersDir + fileInfo.Name()

		infos, err := loadRetainInfos(path, fileInfo.Size())
		if err != nil {
			log.Panicf("loadRetainInfos: %v", err)
		}

		retainInfos = append(retainInfos, infos...)

		log.Printf("%.0f%%: loaded %s", float32(i)/float32(len(entries)-1)*100, path)
	}

	if len(pieceIDs) > 0 {
		fmt.Printf("---\nchecking against the list first…\n")
		for _, p := range pieceIDs {
			checkFilters(retainInfos, p, *quiet)
		}
		if *quiet {
			return
		}
		fmt.Printf("---\nswitching to interactive mode…\n")
	}

	for {
		fmt.Printf("(q to exit) > ")

		var s string
		if _, err := fmt.Scanln(&s); err != nil {
			fmt.Println("couldn't scan; q to exit")
			continue
		}

		if s == "q" {
			return
		}

		pid, err := storj.PieceIDFromString(s)
		if err != nil {
			fmt.Printf("couldn't decode %q: %v\n", s, err)
			continue
		}

		checkFilters(retainInfos, pid, false)
	}
}

func checkFilters(infos []*internalpb.RetainInfo, pid storj.PieceID, quiet bool) {
	for _, info := range infos {
		f, err := bloomfilter.NewFromBytes(info.Filter)
		if err != nil {
			log.Panicf("NewFromBytes: %v", err)
		}
		if f.Contains(pid) {
			if !quiet {
				fmt.Printf("BF (fill=%.2f, size=%d) for %s (piece count=%d) contains this piece\n", f.FillRate(), f.Size(), info.StorageNodeId, info.PieceCount)
			}
			return
		}
	}
	fmt.Printf("%s not found\n", pid)
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

func loadPieceIDs(path string) ([]storj.PieceID, error) {
	if path == "" {
		return nil, nil
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var ids []storj.PieceID
	for _, s := range strings.Split(string(b), "\n") {
		if s == "" {
			continue
		}
		id, err := storj.PieceIDFromString(s)
		if err != nil {
			return nil, fmt.Errorf("couldn't decode %q: %w", id, err)
		}
		ids = append(ids, id)
	}
	return ids, nil
}
