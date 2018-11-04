package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	a := make([]byte, 20)
	rand.Read(a)
	fmt.Println(hex.EncodeToString(a))
}

//
// #=> 363cad5276c1408031641de483dcd4e7e9e2445b
