package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 乱数のシードを生成
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		// 浮動小数点の乱数を生成
		fmt.Println(rand.Float64())
	}
}
