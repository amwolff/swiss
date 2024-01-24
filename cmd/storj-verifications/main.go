package main

import (
	"context"

	"storj.io/common/pb"
	"storj.io/uplink/private/piecestore"
)

func main() {
	withPieceHashAlgo()
}

func withPieceHashAlgo() {
	ctx := context.TODO()
	ctx = piecestore.WithPieceHashAlgo(ctx, pb.PieceHashAlgorithm_BLAKE3)
	algo := piecestore.GetPieceHashAlgo(ctx)
	if algo != pb.PieceHashAlgorithm_BLAKE3 {
		panic(algo.String())
	}
}
