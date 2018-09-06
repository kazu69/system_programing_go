package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")

	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("GET", "http://localhost:8888", nil)

	if err != nil {
		panic(err)
	}

	request.Write(conn)

	response, err := http.ReadResponse(bufio.NewReader(conn), request)

	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, true)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(dump))
}

// server
// Server is running at localhost:8888
// Accept 127.0.0.1:50436
// GET / HTTP/1.1
// Host: localhost:8888
// User-Agent: Go-http-client/1.1

// client
// HTTP/1.0 200 OK
// Connection: close
//
// Hello World
