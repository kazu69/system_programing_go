package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

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

		// 非同期実行される
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			// リクエストを読み込む
			// HTTPリクエストのヘッダー、メソッド、パスなどの情報を切り出す
			request, err := http.ReadRequest(
				bufio.NewReader(conn))

			if err != nil {
				panic(err)
			}

			dump, err := httputil.DumpRequest(request, true)

			if err != nil {
				panic(err)
			}

			fmt.Println(string(dump))
			// http.Response構造体はWrite()メソッドを持っているので、
			// 作成したレスポンスのコンテンツをio.Writerに直接書き込むことができる。
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body: ioutil.NopCloser(
					strings.NewReader("Hello World\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}

// server
// Server is running at localhost:8888
// Accept 127.0.0.1:65121
// GET / HTTP/1.1
// Host: localhost:8888
// Accept: */*
// User-Agent: curl/7.43.0

// clisnt
// curl http://localhost:8888
// Hello World
