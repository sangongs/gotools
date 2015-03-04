package main

import "os"
import "fmt"
import "net"

func handle_conn(conn net.Conn) {
    defer conn.Close()

    const buf_size = 1280

    buf := make([]byte, buf_size)

    for {
        read_len, err := conn.Read(buf)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Read error: %s\n", err.Error())
            break
        }

        if read_len == 0 {
            break
        }
    }
}

func main() {
    ln, err := net.Listen("tcp", ":"+os.Args[1])
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
        return
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Accept error: %s\n", err.Error())
            continue
        }

        go handle_conn(conn)
    }
}
