package main

import (
	"context"
	"fmt"
	"net"

	"github.com/matti/betterio"
)

func main() {
	ctx := context.Background()
	dialer := net.Dialer{}
	upstream, err := dialer.DialContext(ctx, "tcp", "google.com:80")
	if err != nil {
		panic(err)
	}

	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("listening 0.0.0.0:8080")
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}

	bytesDown, bytesUp := betterio.CopyBidirUntilCloseAndReturnBytesWritten(conn, upstream)
	fmt.Println(bytesDown, bytesUp)
}
