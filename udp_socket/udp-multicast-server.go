package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Start tick server at 224.0.0.1:9999")
	conn, err := net.Dial("udp", "224.0.0.1:9999")

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	start := time.Now()

	wait := 10*time.Second - time.Nanosecond*time.Duration(start.UnixNano()%(10*1000*1000*1000))
	time.Sleep(wait)

	ticker := time.Tick(10 * time.Second)
	for now := range ticker {
		conn.Write([]byte(now.String()))
		fmt.Println("Tick: ", now.String())
	}
}

// server
// Start tick server at 224.0.0.1:9999
// Tick:  2018-09-14 10:46:31.099404 +0900 JST m=+14.234908317
// Tick:  2018-09-14 10:46:40.003227 +0900 JST m=+23.138408200
// Tick:  2018-09-14 10:46:50.001619 +0900 JST m=+33.136421145
