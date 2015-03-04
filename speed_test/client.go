package main

import (
    "fmt"
    "os"
    "net"
    "strconv"
    "time"
)

var total_sent = map[net.Conn]int{}

func handle_conn(conn net.Conn) {
    defer conn.Close()

    const buf_size = 1024
    buf := make([]byte, buf_size)

    total_sent[conn] = 0
    for {
        write_len, err := conn.Write(buf)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Write error: %s\n", err.Error())
            break;
        }

        total_sent[conn] += write_len
    }
}

func main() {
    if len(os.Args) != 4 {
        fmt.Fprintf(os.Stderr, "Usage: %s host port nr_threads\n", os.Args[0])
        return
    }

    nr_threads,_ := strconv.Atoi(os.Args[3])
    for i := 0; i < nr_threads; i++ {
        conn, err := net.Dial("tcp", os.Args[1] + ":" + os.Args[2])
        if err != nil {
            fmt.Fprintf(os.Stderr, "Dial error: %s\n", err.Error())
            return
        }

        go handle_conn(conn)
    }

    var last_sent int
    last_ts := time.Now()
    for {
        time.Sleep(time.Second)
        var sent int
        for _, value := range total_sent {
            sent += value
        }

        ts := time.Now()
        fmt.Printf("%f\n", float64(sent-last_sent) / ts.Sub(last_ts).Seconds())
        last_sent = sent
        last_ts = ts
    }
}
