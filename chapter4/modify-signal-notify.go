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

	exit_chan := make(chan int)

	// SIGINT (ctrl + C)を受け取る
	signal.Notify(signals, syscall.SIGINT)

	// シグナルが来るまで待つ
	fmt.Println("waiting SIGINT (CTRL+C)")
	go func() {
		for {
			s := <-signals
			switch s {
			case syscall.SIGINT:
				fmt.Println("signal arrive syscall.SIGINT")
				fmt.Println("signal arrive ", s)
				exit_chan <- 0
			default:
				fmt.Println("signal arrive ", s)
				exit_chan <- 0
			}
		}
	}()
	fmt.Println("SIGINT arrived")
	exit_code := <-exit_chan
	os.Exit(exit_code)
}
