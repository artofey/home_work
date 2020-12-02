package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// TelnetClient ...
type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

// NewTelnetClient create new instance telnet client.
func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
		dealer:  net.Dialer{},
	}
}

type telnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	dealer  net.Dialer
	ctx     context.Context
	conn    net.Conn
	cancel  context.CancelFunc
}

func (c *telnetClient) Connect() error {
	c.ctx, c.cancel = context.WithTimeout(
		context.Background(),
		time.Duration(timeout)*time.Second,
	)

	var err error
	c.conn, err = c.dealer.DialContext(c.ctx, "tcp", c.address)
	if err != nil {
		return err
	}

	return nil
}

func (c *telnetClient) Close() error {
	c.cancel()
	return nil
}

func (c *telnetClient) Send() error {
	return nil
}

func (c *telnetClient) Receive() error {
	return nil
}

func (c *telnetClient) readRoutine() {
	scanner := bufio.NewScanner(c.conn)
OUTER:
	for {
		select {
		case <-c.ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN")
				break OUTER
			}
			text := scanner.Text()
			log.Printf("From server: %s", text)
		}
	}
	log.Printf("Finished readRoutine")
}

func (c *telnetClient) writeRoutine() {
	scanner := bufio.NewScanner(c.in)
OUTER:
	for {
		select {
		case <-c.ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				break OUTER
			}
			str := scanner.Text()
			log.Printf("To server %v\n", str)

			c.conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}
	log.Printf("Finished writeRoutine")
}
