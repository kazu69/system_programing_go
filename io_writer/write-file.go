package main

import (
	"os"
)

func main() {
	// ファイルへの書き込み
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	file.Write([]byte("os File example\n"))
	file.Close()
}
