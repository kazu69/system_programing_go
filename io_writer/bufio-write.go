package main

import (
	"bufio"
	"os"
)

func main() {
	// 出力結果を保持しておきFlushで出力
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer\n")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
