// usage: node-ids-to-hex path/to/nodes.csv
//
// nodes.csv is
//
// id1,…
// id2,…
// …

package main

import (
	"encoding/csv"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"storj.io/common/storj"
)

func main() {
	flag.Parse()

	if n := flag.NArg(); n != 1 {
		log.Fatalf("not enough/too many args (has %d, needs 1)", n)
	}

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Panicf("Open: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	for i := 0; ; i++ {
		records, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Panicf("Read: %v", err)
		}
		id, err := storj.NodeIDFromString(records[0])
		if err != nil {
			log.Printf("couldn't decode %q: %v", records[0], err)
			continue
		}
		fmt.Println(hex.EncodeToString(id.Bytes()))
	}
}
