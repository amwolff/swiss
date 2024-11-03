// usage: fake-pieceids [-c=count]
package main

import (
	"flag"
	"fmt"

	"storj.io/common/testrand"
)

func main() {
	count := flag.Int("c", 1, "count")
	flag.Parse()

	for i := 0; i < *count; i++ {
		fmt.Println(testrand.PieceID().String())
	}
}
