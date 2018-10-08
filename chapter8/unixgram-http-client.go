package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "unixdomainsocket-sample")

	// ソケットファイルを削除
	os.Remove(path)

	fmt.Println("Server is running at " + path)

	// 	conn, err := net.Dial("unixgram", path)
	// 上記の場合 開いたsocketファイルは送信専用のため、レスポンスを受け取れない
	// そのため net.ListenPacket で WriteTo, ReadFromを使って送受信できるようにする必要がある

	conn, err := net.ListenPacket("unixgram", path)

	if err != nil {
		panic(err)
	}

	// 送信先のアドレス
	//
	unixServerAddr, err := net.ResolveUnixAddr("unixgram", path)
	var serverAddr net.Addr = unixServerAddr

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	log.Println("Sending to server")

	_, err = conn.WriteTo([]byte("Hello from Client"), serverAddr)

	if err != nil {
		panic(err)
	}

	log.Println("Receiving from server")

	buffer := make([]byte, 1500)

	length, _, err := conn.ReadFrom(buffer)

	if err != nil {
		panic(err)
	}

	log.Printf("Received: %s\n", string(buffer[:length]))
}
