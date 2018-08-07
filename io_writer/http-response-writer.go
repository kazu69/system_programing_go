package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample")
}

func main() {
	// ブラウザにメッセージを書き込む
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
