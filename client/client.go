package main

import (
    "fmt"
    "log"
    "net/rpc"
    "os"
)

func main() {
	if len(os.Args) != 2 {
        fmt.Println("Usage: server:port")
        os.Exit(1)
    }
    service := os.Args[1]

    _, err := rpc.Dial("tcp", service)
    if err != nil {
        log.Fatal("dialing:", err)
    }

}
