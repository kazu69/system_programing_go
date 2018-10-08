package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `1 1行目
2 2行目
3 3行目`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Println("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}
