package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 書かれた内容をbufferで記憶
	var buffer bytes.Buffer
	buffer.Write([]byte("byte.Buffer example\n"))
	fmt.Println(buffer.String())
}}
