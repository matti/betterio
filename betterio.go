package betterio

import (
	"io"
)

// CopyBidirUntilCloseAndReturnBytesWritten ...
func CopyBidirUntilCloseAndReturnBytesWritten(c1 io.ReadWriteCloser, c2 io.ReadWriteCloser) (int64, int64) {
	var c1N int64
	var c2N int64

	c1Close := make(chan int64)
	c2Close := make(chan int64)

	go func(src io.ReadWriteCloser, dst io.ReadWriteCloser) {
		c2Close <- copy(dst, src)
	}(c2, c1)

	go func(src io.ReadWriteCloser, dst io.ReadWriteCloser) {
		c1Close <- copy(dst, src)
	}(c1, c2)

	select {
	case c1N = <-c1Close:
		c2N = <-c2Close
	case c2N = <-c2Close:
		c1N = <-c1Close
	}

	return c1N, c2N
}

func copy(src io.ReadWriteCloser, dst io.ReadWriteCloser) int64 {
	n, err := io.Copy(src, dst)
	if err != nil {
		//		log.Printf("Copy error: %s", err)
	}
	if err := src.Close(); err != nil {
		//		log.Printf("Close error: %s", err)
	}

	return n
}
