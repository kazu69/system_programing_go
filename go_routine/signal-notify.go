package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// サイズより大きいチャネル作成
	signals := make(chan os.Signal, 1)
	// SIGINT (ctrl + C)を受け取る
	signal.Notify(signals, syscall.SIGINT)

	// シグナルが来るまで待つ
	fmt.Println("waiting SIGINT (CTRL+C)")
	<-signals
	fmt.Println("SIGINT arrived")
}
