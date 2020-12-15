# betterio

[examples/CopyBidirUntilCloseAndReturnBytesWritten.go](examples/CopyBidirUntilCloseAndReturnBytesWritten.go)

    bytesDown, bytesUp := betterio.CopyBidirUntilCloseAndReturnBytesWritten(conn, upstream)

[examples/ReadBytesUntilRune.go](examples/ReadBytesUntilRune.go)

    fmt.Println(string(betterio.ReadBytesUntilRune(upstream, '\n')))


[examples/CheckReaderOpen.go](examples/CheckReaderOpen.go)

    if err := betterio.CheckReaderOpen(upstream); err != nil {
        fmt.Println("not open")
    }
