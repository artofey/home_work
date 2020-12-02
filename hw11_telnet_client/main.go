package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	host, port, address string
	timeout             int64
)

func init() {
	flag.Int64Var(&timeout, "timeout", 10, "connection timeout")
	flag.Parse()

	host = flag.Arg(0)
	if flag.Arg(1) == "" {
		port = "23"
	} else {
		port = flag.Arg(1)
	}
	address = net.JoinHostPort(host, port)
}

func main() {
	client := NewTelnetClient(address, time.Duration(timeout), os.Stdin, os.Stdout)
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(1 * time.Minute)
	fmt.Println("Exita app.")
}
