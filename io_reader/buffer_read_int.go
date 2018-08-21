package main

import (
	"fmt"
	"strings"
)

var source = "123 1.23 1.0e4 test"
var source2 = "123, 1.23, 1.0e4, test"

func main() {
	reader := strings.NewReader(source)
	var i int
	var f, g float64
	var s string
	// Fscan space-separated value
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Println("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)

	reader2 := strings.NewReader(source2)
	// comma-space--separated value
	fmt.Fscan(reader2, "%v, %v, %v, %v", &i, &f, &g, &s)
	fmt.Println("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}
