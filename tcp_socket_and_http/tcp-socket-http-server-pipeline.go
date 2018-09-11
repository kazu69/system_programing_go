package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// 順番に従ってconnに書き出しをする
func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	// 順番に取り出す
	for sessionResponse := range sessionResponses {
		// 選択された仕事が終わるまで待つ
		response := <-sessionResponse
		response.Write(conn)
		close(sessionResponse)
	}
}

// セッション内のリクエストを処理する
func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
	dump, err := httputil.DumpRequest(request, true)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(dump))

	content := "Hello World\n"
	// レスポンスを書き込む
	// セッションを維持するためにKeep-Aliveでないといけない
	response := &http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(len(content)),
		Body: ioutil.NopCloser(
			strings.NewReader(content)),
	}
	// 処理が終わったらチャネルに書き込み、
	// ブロックされていたwriteToConnの処理を再始動する
	resultReceiver <- response
}

// セッション1つを処理
func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())

	// セッション内のリクエストを順に処理するためのチャネル
	// バッファ指定する(50)
	// バッファありの場合は、指定した個数までは自由に投入できる
	// 指定した個数のデータが入っているときにさらに追加でデータを投入しようとすると
	// 投入しようとしたスレッド（ゴルーチン）がブロックされる
	//  他のゴルーチンで値を取り出してきて個数が減ると、ブロックが解除されてデータが投入できる
	// バッファなしの場合は、投入しようとするスレッドは即座にブロック
	// channel over channel
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)
	// レスポンスを直列化してソケットに書き出す専用のゴルーチン
	go writeToConn(sessionResponses, conn)
	reader := bufio.NewReader(conn)
	for {
		// レスポンスを受け取ってセッションのキューに
		// 入れる
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		// リクエストを読み込む
		request, err := http.ReadRequest(reader)
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		// 順序を守ってレスポンスが書けるようにunbufferのチャネル
		sessionResponse := make(chan *http.Response)
		// リクエスト処理が終わるまで待つため、送信データをためるバッファなしのチャネル
		// バッファなしのチャネルがバッファありのチャネルに包括されている
		sessionResponses <- sessionResponse
		// 非同期でレスポンスを実行
		go handleRequest(request, sessionResponse)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running at localhost:8888")

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go processSession(conn)
	}
}

// server
// Server is running at localhost:8888
// Accept 127.0.0.1:60143
// GET / HTTP/1.1
// Host: localhost:8888
// Accept: */*
// User-Agent: curl/7.43.0

// clisnt
// curl http://localhost:8888
// Hello World
