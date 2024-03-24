// usage: decrypt-secret-key base64_encrypted_secret_key encryption_key
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"

	"storj.io/common/encryption"
	"storj.io/common/storj"
	"storj.io/edge/pkg/auth/authdb"
)

func main() {
	flag.Parse()

	if n := flag.NArg(); n != 2 {
		log.Fatalf("not enough/too many args (has %d, needs 2)", n)
	}

	sk, err := base64.RawStdEncoding.DecodeString(flag.Arg(0))
	if err != nil {
		log.Panicf("DecodeString: %v", err)
	}

	var ek authdb.EncryptionKey
	if err := ek.FromBase32(flag.Arg(1)); err != nil {
		log.Panicf("FromBase32: %v", err)
	}

	storjKey := ek.ToStorjKey()

	data, err := encryption.Decrypt(sk, storj.EncAESGCM, &storjKey, &storj.Nonce{})
	if err != nil {
		log.Panicf("Decrypt: %v", err)
	}

	fmt.Println(authdb.SecretKey(data).ToBase32())
}
