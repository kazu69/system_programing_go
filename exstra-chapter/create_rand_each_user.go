package main

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

func main() {
	// seedからソース生成
	seed, _ := crand.Int(crand.Read, big.NewInt(math.MaxInt64))
	src := rand.NewSource(seed.Int64())
	// ソースから乱数生成器を生成
	rng := rand.New(src)
}
