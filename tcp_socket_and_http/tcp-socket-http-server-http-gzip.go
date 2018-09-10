package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// クライアントはgzipを受け入れ可能か
func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

// 1セッションの処理をする
func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		request, err := http.ReadRequest(
			bufio.NewReader(conn))

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
			ProtoMinor: 1,
			Header:     make(http.Header),
		}

		var buffer bytes.Buffer

		// クライアントがgzipが受け入れ可能かどうかに応じて、中に入れるコンテンツを変える
		if isGZipAcceptable(request) {
			content := "Hello World (gzipped)\n"
			// コンテンツをgzip化して転送
			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()

			response.Body = ioutil.NopCloser(&buffer)
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "Hello World\n"
			response.Body = ioutil.NopCloser(strings.NewReader(content))
			response.ContentLength = int64(len(content))
		}

		response.Write(conn)
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
// Accept 127.0.0.1:53541
// GET / HTTP/1.1
// Host: localhost:8888
// Accept: */*
// User-Agent: curl/7.43.0

// clisnt
// curl http://localhost:8888
// Hello World
