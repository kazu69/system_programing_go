package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 32bit big endian 1000
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	// byte swapping
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data %d\n", i)
}
