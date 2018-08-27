package main

import (
	"fmt"
	"math"
)

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 10000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func main() {
	pn := primeNumber()
	// 帰ってくるチャネルは for .. range の中で配列のように扱える
	// for loopはチャネルがオープンしている間回り続ける
	for n := range pn {
		fmt.Println(n)
	}
}
