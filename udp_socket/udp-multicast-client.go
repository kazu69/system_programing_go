package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listen tick server at 224.0.0.1:9999")
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")

	if err != nil {
		panic(err)
	}

	// クライアント側で複数のネットワーク接続があるとき
	// 特定のLAN環境のマルチキャストを受信するには

	// eh := net.InterfaceByName("en0")
	// net.ListenMulticastUDP("udp", eh, address)
	listener, err := net.ListenMulticastUDP("udp", nil, address)

	defer listener.Close()

	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := listener.ReadFromUDP(buffer)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Server %v\n", remoteAddress)
		fmt.Printf("Now    %s\n", string(buffer[:length]))
	}
}

// cient
// Listen tick server at 224.0.0.1:9999
// Server 192.168.1.3:62124
// Now    2018-09-14 10:46:50.001619 +0900 JST m=+33.136421145
// Server 192.168.1.3:62124
// Now    2018-09-14 10:47:00.028996 +0900 JST m=+43.163419390
// Server 192.168.1.3:62124
// Now    2018-09-14 10:47:10.006642 +0900 JST m=+53.140741737
