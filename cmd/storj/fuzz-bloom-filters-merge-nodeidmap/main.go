// usage: fuzz-bloom-filters-merge-nodeidmap [-c=(piece)count] [-n=(nodes)count] [-i=iterations]
package main

import (
	"flag"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"

	"storj.io/common/storj"
	"storj.io/common/testrand"
	"storj.io/storj/shared/bloomfilter"
	"storj.io/storj/shared/nodeidmap"
)

type retainInfo struct {
	filter *bloomfilter.Filter
	count  int
}

type nodeWithSeed struct {
	node storj.NodeID
	seed byte
}

func main() {
	count := flag.Int("c", 13535192, "piece count per filter")
	nodes := flag.Int("n", 26562, "number of nodes")
	iterations := flag.Int("i", 100, "iterations")
	flag.Parse()

	var nodeIDs []nodeWithSeed
	for i := 0; i < *nodes; i++ {
		nodeIDs = append(nodeIDs, nodeWithSeed{
			node: testrand.NodeID(),
			seed: bloomfilter.GenerateSeed(),
		})
	}

	for i := 0; i < *iterations; i++ {
		var g errgroup.Group
		var mu sync.Mutex
		filters1 := nodeidmap.Make[retainInfo]()
		filters2 := nodeidmap.Make[retainInfo]()
		pieces := make(map[storj.NodeID][]storj.PieceID)
		for j := 0; j < *nodes; j++ {
			j := j
			g.Go(func() error {
				mem, bf := generate(nodeIDs[j].seed, *count)

				mu.Lock()
				filters1.Store(nodeIDs[j].node, retainInfo{filter: bf, count: *count})
				pieces[nodeIDs[j].node] = mem
				mu.Unlock()

				return nil
			})
		}
		if err := g.Wait(); err != nil {
			panic(err)
		}
		for j := 0; j < *nodes; j++ {
			j := j
			g.Go(func() error {
				mem, bf := generate(nodeIDs[j].seed, *count)

				mu.Lock()
				filters2.Store(nodeIDs[j].node, retainInfo{filter: bf, count: *count})
				pieces[nodeIDs[j].node] = append(pieces[nodeIDs[j].node], mem...)
				mu.Unlock()

				return nil
			})
		}
		if err := g.Wait(); err != nil {
			panic(err)
		}

		filters1.Add(filters2, func(old, new retainInfo) retainInfo {
			old.count += new.count
			if err := old.filter.AddFilter(new.filter); err != nil {
				panic(err)
			}
			return old
		})

		for j := 0; j < *nodes; j++ {
			ri, ok := filters1.Load(nodeIDs[j].node)
			if !ok {
				log.Printf("missing node %s", nodeIDs[j].node)
			}
			if ri.count != 2*(*count) {
				log.Printf("node %s: count %d != %d (len(pieces[node])=%d)", nodeIDs[j].node, ri.count, 2*(*count), len(pieces[nodeIDs[j].node]))
			}
			checkContains(nodeIDs[j].node.String(), pieces[nodeIDs[j].node], ri.filter)
		}
		log.Printf("Iteration %d finished", i)
	}
}

func checkContains(id string, mem []storj.PieceID, filter *bloomfilter.Filter) {
	for j, p := range mem {
		if !filter.Contains(p) {
			seed, hc, size := filter.SeedAndParameters()
			log.Printf("%s: %d: (seed=%v, count=%v, size=%d, fill=%.2f) does not contain %s", id, j, seed, hc, size, filter.FillRate(), p)
		}
	}
}

func generate(seed byte, count int) ([]storj.PieceID, *bloomfilter.Filter) {
	hashCount, tableSize := bloomfilter.OptimalParameters(int64(count), 0.1, 35000000)
	filter := bloomfilter.NewExplicit(seed, hashCount, tableSize)

	var mem []storj.PieceID
	for j := 0; j < count; j++ {
		p := testrand.PieceID()
		filter.Add(p)
		mem = append(mem, p)
	}

	return mem, filter
}
