package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	// POSTするメッセージ
	sendMessages := []string{
		"TEST",
		"GOLANG",
		"SYSTEM_PROGRAMING",
	}

	current := 0
	var conn net.Conn = nil

	// リトライ用にループで全体を囲う
	for {
		// まだコネクションを張ってない / エラーでリトライ時はDialから行う
		if conn == nil {
			conn, err := net.Dial("tcp", "localhost:8888")

			if err != nil {
				panic(err)
			}

			fmt.Printf("Access: %d\n", current)

			// POSTで文字列を送るリクエストを作成
			request, err := http.NewRequest(
				"POST",
				"http://localhost:8888",
				strings.NewReader(sendMessages[current]))

			if err != nil {
				panic(err)
			}

			err = request.Write(conn)

			if err != nil {
				panic(err)
			}

			// サーバから読み込む。タイムアウトはここでエラーになるのでリトライ
			response, err := http.ReadResponse(bufio.NewReader(conn), request)

			if err != nil {
				fmt.Println("Retry")
				conn = nil
				continue
			}

			dump, err := httputil.DumpResponse(response, true)

			if err != nil {
				panic(err)
			}

			fmt.Println(string(dump))

			// 全部送信完了していれば終了
			current++

			if current == len(sendMessages) {
				break
			}
		}
	}

}

// server
// Accept 127.0.0.1:63046
// POST / HTTP/1.1
// Host: localhost:8888
// Content-Length: 4
// User-Agent: Go-http-client/1.1
//
// TEST
// Accept 127.0.0.1:63048
// POST / HTTP/1.1
// Host: localhost:8888
// Content-Length: 6
// User-Agent: Go-http-client/1.1
//
// GOLANG
// Accept 127.0.0.1:63050
// POST / HTTP/1.1
// Host: localhost:8888
// Content-Length: 17
// User-Agent: Go-http-client/1.1
//
// SYSTEM_PROGRAMING

// client
// Access: 0
// HTTP/1.1 200 OK
// Content-Length: 12
//
// Hello World
//
// Access: 1
// HTTP/1.1 200 OK
// Content-Length: 12
//
// Hello World
//
// Access: 2
// HTTP/1.1 200 OK
// Content-Length: 12
//
// Hello World
