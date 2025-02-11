// usage: fuzz-bloom-filters-merge [-c=(piece)count] [-d=duration] [-p=parallelism]

package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"

	"storj.io/common/storj"
	"storj.io/common/testrand"
	"storj.io/storj/shared/bloomfilter"
)

func main() {
	count := flag.Int("c", 13535192, "max piece count per filter")
	duration := flag.Duration("d", 24*time.Hour, "duration")
	parallelism := flag.Int("p", 8, "parallelism")
	flag.Parse()

	var g errgroup.Group

	for i := 0; i < *parallelism; i++ {
		i := i
		t := time.Now()
		g.Go(func() error {
			for j := 0; time.Now().Before(t.Add(*duration)); j++ {
				id := strconv.Itoa(i) + "-" + strconv.Itoa(j)

				mem1, bf1 := generate(*count)
				mem2, bf2 := generate(*count)

				if err := bf1.AddFilter(bf2); err != nil {
					panic(err)
				}

				checkContains(id, mem1, bf1)
				checkContains(id, mem2, bf1)
			}
			return nil
		})
	}

	log.Println("testing…")

	if err := g.Wait(); err != nil {
		panic(err)
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

func generate(count int) ([]storj.PieceID, *bloomfilter.Filter) {
	randCount := testrand.Intn(count)

	if h := count / 2; randCount < h {
		randCount = h
	}

	hashCount, tableSize := bloomfilter.OptimalParameters(int64(randCount), 0.1, 35000000)
	filter := bloomfilter.NewExplicit(bloomfilter.GenerateSeed(), hashCount, tableSize)

	var mem []storj.PieceID
	for j := 0; j < randCount; j++ {
		p := testrand.PieceID()
		filter.Add(p)
		mem = append(mem, p)
	}

	return mem, filter
}
