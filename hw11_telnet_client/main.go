package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
	if host == "" {
		log.Fatal("укажите адрес подключения")
	}
	if flag.Arg(1) == "" {
		port = "23"
	} else {
		port = flag.Arg(1)
	}
	address = net.JoinHostPort(host, port)
}

func main() {
	client := NewTelnetClient(address, time.Duration(timeout)*time.Second, os.Stdin, os.Stdout)
	err := client.Connect()
	if err != nil {
		log.Fatal(fmt.Errorf("onnection error %w", err))
	}
	// defer func() {
	// 	if err := client.Close(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		if err := client.Send(); err != nil {
			log.Println(err)
		}
		wg.Done()
	}()
	go func() {
		if err := client.Receive(); err != nil {
			log.Println(err)
		}
		wg.Done()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		fmt.Println(" Exit app.")
		client.Close()
	}()
	wg.Wait()
}
