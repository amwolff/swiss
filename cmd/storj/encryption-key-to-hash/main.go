// usage: encryption-key-to-hash encryption_key
//
// fixture: select * from records where encryption_key_hash = from_hex('…');
package main

import (
	"flag"
	"fmt"
	"log"

	"storj.io/edge/pkg/auth/authdb"
)

func main() {
	flag.Parse()

	if n := flag.NArg(); n != 1 {
		log.Fatalf("not enough/too many args (has %d, needs 1)", n)
	}

	var ek authdb.EncryptionKey
	if err := ek.FromBase32(flag.Arg(0)); err != nil {
		log.Panicf("FromBase32: %v", err)
	}

	fmt.Println(ek.Hash().ToHex())
}
