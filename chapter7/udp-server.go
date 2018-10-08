package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running at localhost:8888")

	// TCPのように「クライアントを待つ」インタフェースではなく、
	// データ送受信のためのnet.PacketConnというインタフェースが即座に返される
	conn, err := net.ListenPacket("udp", "localhost:8888")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	buffer := make([]byte, 1500)

	for {
		// ReadFrom()メソッドを使うと、通信内容を読み込むと同時に、
		// 接続してきた相手のアドレス情報が受け取れる
		// net.PacketConnは、サーバ側でクライアントを知らない状態で開かれるソケットなので、
		// このインタフェースを使ってサーバから先にメッセージを送ることはできません

		// ReadFrom()では、TCPのときに紹介した「データの終了を探りながら受信」といった高度な読み込みはできない
		// そのため、データサイズが決まらないデータに対しては、フレームサイズ分のバッファや、
		// 期待されるデータの最大サイズ分のバッファを作り、そこにデータをまとめて読み込むことになります。
		// あるいは、バイナリ形式のデータにしてヘッダにデータ長などを格納しておき、
		// そこまで先読みしてから必要なバッファを確保して読み込む
		length, remoteAddress, err := conn.ReadFrom(buffer)

		if err != nil {
			panic(err)
		}

		fmt.Println("Received from %v: %v\n", remoteAddress, string(buffer[:length]))

		// ReadFrom()で取得したアドレスに対しては、net.PacketConnインタフェースのWriteTo()メソッドを使ってデータを返送することができる
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)

		if err != nil {
			panic(err)
		}
	}
}
