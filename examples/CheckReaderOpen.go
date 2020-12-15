package main

import (
	"fmt"
	"net"
	"time"

	"github.com/matti/betterio"
)

func main() {
	dialer := net.Dialer{}
	upstream, err := dialer.Dial("tcp", "google.com:80")
	if err != nil {
		panic(err)
	}
	if err := betterio.CheckReaderOpen(upstream); err == nil {
		fmt.Println("upstream conn ok")
	}

	upstream.Write([]byte("google don't be evil, okay?\n\n"))
	time.Sleep(1 * time.Second)

	if err := betterio.CheckReaderOpen(upstream); err != nil {
		fmt.Println("upstream conn not ok after 1s")
	}
}
