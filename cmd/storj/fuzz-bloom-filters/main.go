// usage: fuzz-bloom-filters [-c=(piece)count] [-d=duration] [-p=parallelism] [-i=reportEvery]

package main

import (
	"flag"
	"log"
	"time"

	"golang.org/x/sync/errgroup"

	"storj.io/common/storj"
	"storj.io/common/testrand"
	"storj.io/storj/shared/bloomfilter"
)

func main() {
	count := flag.Int("c", 13535192, "max piece count")
	duration := flag.Duration("d", 24*time.Hour, "duration")
	parallelism := flag.Int("p", 8, "parallelism")
	reportEvery := flag.Int("i", 1, "print a report every ith iteration")
	flag.Parse()

	t := time.Now()
	f := func(id int) {
		for i := 0; time.Now().Before(t.Add(*duration)); i++ {
			randCount := testrand.Intn(*count)

			if h := *count / 2; randCount < h {
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

			for j, p := range mem {
				if !filter.Contains(p) {
					seed, hc, size := filter.SeedAndParameters()
					log.Printf("%d: %d: (seed=%v, count=%v, size=%d, fill=%.2f) does not contain %s", id, j, seed, hc, size, filter.FillRate(), p)
				}
			}

			if i > 0 && i%*reportEvery == 0 {
				log.Printf("%d: Iteration %d ((fill=%.2f, size=%d), (pieces=%d))", id, i, filter.FillRate(), filter.Size(), len(mem))
			}
		}
	}

	var g errgroup.Group

	for i := 0; i < *parallelism; i++ {
		i := i
		g.Go(func() error {
			f(i)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		panic(err)
	}
}
