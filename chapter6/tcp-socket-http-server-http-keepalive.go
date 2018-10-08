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
			defer conn.Close()

			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			// Accept後のソケットで何度も応答を返すためにループ
			// コネクションが張られた後に何度もリクエストを受けられるようにしている
			for {
				// タイムアウトを設定
				// 通信がしばらくないとタイムアウトのエラーでRead()の呼び出しを終了します。
				// 設定しなければ相手からレスポンスがあるまでずっとブロックし続けます。
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))

				// リクエストを読み込む
				// HTTPリクエストのヘッダー、メソッド、パスなどの情報を切り出す
				request, err := http.ReadRequest(
					bufio.NewReader(conn))

				// タイムアウトもしくはソケットクローズ時は終了
				// それ以外はエラーにする
				if err != nil {
					// ダウンキャスト
					neterr, ok := err.(net.Error)
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}

				dump, err := httputil.DumpRequest(request, true)

				if err != nil {
					panic(err)
				}

				fmt.Println(string(dump))
				content := "Hello World\n"
				// http.Response構造体はWrite()メソッドを持っているので、
				// 作成したレスポンスのコンテンツをio.Writerに直接書き込むことができる。
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body: ioutil.NopCloser(
						strings.NewReader(content)),
				}
				response.Write(conn)
			}
		}()
	}
}

// server
// Server is running at localhost:8888
// Accept 127.0.0.1:50228
// GET / HTTP/1.1
// Host: localhost:8888
// Accept: */*
// User-Agent: curl/7.43.0

// clisnt
// curl http://localhost:8888
// Hello World
