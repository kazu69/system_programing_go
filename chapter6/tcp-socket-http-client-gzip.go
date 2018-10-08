package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
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

	for {
		if conn == nil {
			conn, err := net.Dial("tcp", "localhost:8888")

			if err != nil {
				panic(err)
			}

			fmt.Printf("Access: %d\n", current)

			request, err := http.NewRequest(
				"POST",
				"http://localhost:8888",
				strings.NewReader(sendMessages[current]))

			if err != nil {
				panic(err)
			}

			// リクエストヘッダの"Accept-Encoding"にgzip圧縮追加
			request.Header.Set("Accept-Encoding", "gzip")

			err = request.Write(conn)

			if err != nil {
				panic(err)
			}

			response, err := http.ReadResponse(bufio.NewReader(conn), request)

			if err != nil {
				fmt.Println("Retry")
				conn = nil
				continue
			}

			dump, err := httputil.DumpResponse(response, false)

			if err != nil {
				panic(err)
			}

			defer response.Body.Close()
			fmt.Println(string(dump))

			// Accept-Encodingで表明した圧縮メソッドにサーバが対応していたかどうかは、
			// Content-Encodingヘッダーで確認
			// 表明したアルゴリズムに対応していれば、そのアルゴリズム名がそのまま返ってくる
			var reader io.ReadCloser
			switch response.Header.Get("Content-Encoding") {
			case "gzip":
				reader, err = gzip.NewReader(response.Body)

				if err != nil {
					panic(err)
				}

				defer reader.Close()
			default:
				reader = response.Body
			}

			io.Copy(os.Stdout, reader)

			current++

			if current == len(sendMessages) {
				break
			}
		}
	}

}

// server
// Accept 127.0.0.1:52954
// POST / HTTP/1.1
// Host: localhost:8888
// Accept-Encoding: gzip
// Content-Length: 4
// User-Agent: Go-http-client/1.1

// TEST
// Accept 127.0.0.1:52956
// POST / HTTP/1.1
// Host: localhost:8888
// Accept-Encoding: gzip
// Content-Length: 6
// User-Agent: Go-http-client/1.1

// GOLANG
// Accept 127.0.0.1:52958
// POST / HTTP/1.1
// Host: localhost:8888
// Accept-Encoding: gzip
// Content-Length: 17
// User-Agent: Go-http-client/1.1

// SYSTEM_PROGRAMING

// client
// Access: 0
// HTTP/1.1 200 OK
// Content-Length: 46
// Content-Encoding: gzip

// Hello World (gzipped)
// Access: 1
// HTTP/1.1 200 OK
// Content-Length: 46
// Content-Encoding: gzip

// Hello World (gzipped)
// Access: 2
// HTTP/1.1 200 OK
// Content-Length: 46
// Content-Encoding: gzip

// Hello World (gzipped)
