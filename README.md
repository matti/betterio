# betterio

[examples/CopyBidirUntilCloseAndReturnBytesWritten.go](examples/CopyBidirUntilCloseAndReturnBytesWritten.go)

    bytesDown, bytesUp := betterio.CopyBidirUntilCloseAndReturnBytesWritten(conn, upstream)

[examples/ReadBytesUntilRune.go](examples/ReadBytesUntilRune.go)

    fmt.Println(string(betterio.ReadBytesUntilRune(upstream, '\n')))
