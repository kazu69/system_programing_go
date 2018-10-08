package main

import "os"
import "net/http"

func main() {
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-Test", "Custom Header")
	request.Write(os.Stdout)
}
