package main

import (
	"fmt"
	"net"

	"github.com/matti/betterio"
)

func main() {
	dialer := net.Dialer{}
	upstream, err := dialer.Dial("tcp", "localhost:5900")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(betterio.ReadBytesUntilRune(upstream, '\n')))
}
