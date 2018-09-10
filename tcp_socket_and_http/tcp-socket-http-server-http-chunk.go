package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

// 青空文庫: ごんぎつねより
// http://www.aozora.gr.jp/cards/000121/card628.html
var contents = []string{
	"これは、私わたしが小さいときに、村の茂平もへいというおじいさんからきいたお話です。",
	"むかしは、私たちの村のちかくの、中山なかやまというところに小さなお城があって、",
	"中山さまというおとのさまが、おられたそうです。",
	"その中山から、少しはなれた山の中に、「ごん狐ぎつね」という狐がいました。",
	"ごんは、一人ひとりぼっちの小狐で、しだの一ぱいしげった森の中に穴をほって住んでいました。",
	"そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。",
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()

	for {
		// リクエストを読み込む
		request, err := http.ReadRequest(bufio.NewReader(conn))

		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		dump, err := httputil.DumpRequest(request, true)

		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))

		// レスポンスを書き込む
		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"",
			"",
		}, "\r\n"))

		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}

		fmt.Fprintf(conn, "0\r\n\r\n")
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
// Accept 127.0.0.1:58436
// GET / HTTP/1.1
// Host: localhost:8888
// Accept: */*
// User-Agent: curl/7.43.0

// 123 これは、私わたしが小さいときに、村の茂平もへいというおじいさんからきいたお話です。
// 117 むかしは、私たちの村のちかくの、中山なかやまというところに小さなお城があって、
// 69 中山さまというおとのさまが、おられたそうです。
// 108 その中山から、少しはなれた山の中に、「ごん狐ぎつね」という狐がいました。
// 132 ごんは、一人ひとりぼっちの小狐で、しだの一ぱいしげった森の中に穴をほって住んでいました。
// 102 そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。
// 0

// clisnt
// curl http://localhost:8888
// これは、私わたしが小さいときに、村の茂平もへいというおじいさんからきいたお話です。むかしは、私たちの村のちかくの、中山なかやまというところに小さなお城があって、中山さまというおとのさまが、おられたそうです。その中山から、少しはなれた山の中に、「ごん狐ぎつね」という狐がいました。ごんは、一人ひとりぼっちの小狐で、しだの一ぱいしげった森の中に穴をほって住んでいました。そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。%
