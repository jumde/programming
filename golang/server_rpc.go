package main

import (
    "fmt"
    "log"
    "net"
    "net/rpc"
)

type Listener int
type Reply struct {
    Data string
}

func (L *Listener) GetLine(line []byte, reply *Reply) error {
    rv := string(line)
    fmt.Printf("Receive: %v\n", rv)
    *reply = Reply{rv}
    return nil
}

func main() {
    address, err := net.ResolveTCPAddr("tcp", "0.0.0.0:12345")
    if err != nil {
        log.Fatal(err)
    }
    inbound, err := net.ListenTCP("tcp", address)
    if err != nil {
        log.Fatal(err)
    }
    listener := new(Listener)
    rpc.Register(listener)
    rpc.Accept(inbound)
}
