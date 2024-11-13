// usage: fuzz-bloom-filters [-c=(piece)count] [-d=duration]

package main

import (
	"flag"
	"log"
	"time"

	"storj.io/common/storj"
	"storj.io/common/testrand"
	"storj.io/storj/shared/bloomfilter"
)

func main() {
	count, duration := flag.Int("c", 13535192, "piece count"), flag.Duration("d", 24*time.Hour, "duration")
	flag.Parse()

	t := time.Now()
	for i := 0; time.Now().Before(t.Add(*duration)); i++ {
		hashCount, tableSize := bloomfilter.OptimalParameters(400000, 0.1, 35000000)
		filter := bloomfilter.NewExplicit(bloomfilter.GenerateSeed(), hashCount, tableSize)

		var mem []storj.PieceID
		for j := 0; j < *count; j++ {
			p := testrand.PieceID()
			filter.Add(p)
			mem = append(mem, p)
		}

		for j, p := range mem {
			if !filter.Contains(p) {
				seed, hc, size := filter.SeedAndParameters()
				log.Printf("%d: (seed=%v, count=%v, size=%d, fill=%.2f) does not contain %s", j, seed, hc, size, filter.FillRate(), p)
			}
		}

		log.Printf("Iteration %d (size=%d, fill=%.2f)", i, filter.Size(), filter.FillRate())
	}
}
