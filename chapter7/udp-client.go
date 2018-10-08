package main

import (
	"fmt"
	"net"
)

func main() {
	// クライアントでは相手がわかった上でDial()するので、
	// TCPの場合と同じようにio.Reader、io.Writerインタフェースのまま使うこともできる
	conn, err := net.Dial("udp", "127.0.0.1:8888")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Sending to server")

	_, err = conn.Write([]byte("Hello from Client"))

	if err != nil {
		panic(err)
	}

	fmt.Println("Receiving from server")

	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Received: %s\n", string(buffer[:length]))
}
