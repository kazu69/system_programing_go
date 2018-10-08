package main

import (
	"fmt"
)

func main() {
	fmt.Println("start sub()")
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is Finish")
		done <- true
	}()
	// 終了待ち
	<-done
	fmt.Println("all tasks are finished")
}
