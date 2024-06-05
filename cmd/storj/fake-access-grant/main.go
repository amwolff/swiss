// usage: fake-access-grant
package main

import (
	"fmt"

	"storj.io/common/grant"
	"storj.io/common/macaroon"
	"storj.io/common/testrand"
)

func main() {
	key := testrand.Key()

	apiKey, err := macaroon.NewAPIKey(key[:])
	if err != nil {
		panic(err)
	}

	inner := grant.Access{
		SatelliteAddress: "12m535MGGhhNNGNXQTf96fmbg1JxDXMHPbrBozwJbmwQhxb6m3w@example.com:7777",
		APIKey:           apiKey,
		EncAccess:        grant.NewEncryptionAccess(),
	}

	serialized, err := inner.Serialize()
	if err != nil {
		panic(err)
	}

	fmt.Println(serialized)
}
