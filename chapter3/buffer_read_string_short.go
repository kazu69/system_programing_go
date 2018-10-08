package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1 1行目
2 2行目
3 3行目`

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))

	// scannerはdefaultでdelimiterが改行
	// 以下の設定でspace区切りになる

	// scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println("%#v\n", scanner.Text())
	}
}
