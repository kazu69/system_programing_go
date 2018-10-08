package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "unixdomainsocket-sample")

	// ソケットファイルを削除
	os.Remove(path)

	fmt.Println("Server is running at " + path)

	// unixgram
	// udp相当の使い方のできるデータグラム型のUnixドメインソケット
	conn, err := net.ListenPacket("unixgram", path)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	buffer := make([]byte, 1500)

	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)

		if err != nil {
			panic(err)
		}

		fmt.Println("Received from %v: %v\n", remoteAddress, string(buffer[:length]))

		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)

		if err != nil {
			panic(err)
		}
	}
}
