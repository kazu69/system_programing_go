package main

import "os"

func main() {
	// 標準出力への書き込み
	os.Stdout.Write([]byte("os.Stdout example\n"))
}
