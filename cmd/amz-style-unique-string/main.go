// usage: amz-style-unique-string
package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"storj.io/common/uuid"
)

func main() {
	u, err := uuid.New()
	if err != nil {
		panic(err)
	}
	var result [16]byte
	hex.Encode(result[0:16], u.Bytes()[0:8])
	fmt.Println(strings.ToUpper(string(result[:])))
}
