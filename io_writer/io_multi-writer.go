package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("mutiwriter.txt")
	if err != nil {
		panic(err)
	}

	// file stdout にwrite
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}
