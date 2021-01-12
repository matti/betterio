package main

import (
	"fmt"
	"os"

	"github.com/matti/betterio"
)

func main() {
	file, _ := os.Open("/etc/resolv.conf")
	fmt.Println(betterio.ReadAllString(file))
}
