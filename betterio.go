package betterio

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"syscall"
)

// ReadBytesUntilRune ...
func ReadBytesUntilRune(reader io.Reader, r rune) []byte {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)

	buf := new(bytes.Buffer)

	for scanner.Scan() {
		buf.Write(scanner.Bytes())

		if scanner.Bytes()[0] == byte(r) {
			break
		}
	}

	return buf.Bytes()
}

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

// CheckReaderOpen ...
func CheckReaderOpen(conn io.Reader) error {
	var errUnexpectedRead = errors.New("unexpected read")
	var sysErr error

	sysConn, ok := conn.(syscall.Conn)
	if !ok {
		return nil
	}
	rawConn, err := sysConn.SyscallConn()
	if err != nil {
		return err
	}

	err = rawConn.Read(func(fd uintptr) bool {
		var buf [1]byte
		n, err := syscall.Read(int(fd), buf[:])
		switch {
		case n == 0 && err == nil:
			sysErr = io.EOF
		case n > 0:
			sysErr = errUnexpectedRead
		case err == syscall.EAGAIN || err == syscall.EWOULDBLOCK:
			sysErr = nil
		default:
			sysErr = err
		}
		return true
	})
	if err != nil {
		return err
	}

	return sysErr
}
