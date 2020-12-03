package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
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
		c.timeout,
	)

	var err error
	c.conn, err = c.dealer.DialContext(c.ctx, "tcp", c.address)
	if err != nil {
		return err
	}

	return nil
}

func (c *telnetClient) Close() error {
	fmt.Fprintln(c.out, "Client is Closed.")
	c.cancel()
	return nil
}

func (c *telnetClient) Send() error {
	scanner := bufio.NewScanner(c.in)
	defer c.Close()
OUTER:
	for {
		select {
		case <-c.ctx.Done():
			fmt.Fprintln(c.out, "Sender DONE")
			break OUTER
		default:
			if !scanner.Scan() {
				fmt.Fprintln(c.out, "Not Scan. EOF.")
				if err := scanner.Err(); err != nil {
					fmt.Fprintf(c.out, "Sender Scan Error: %v\n", err)
				}
				break OUTER
			}
			str := scanner.Text()
			fmt.Fprintf(c.out, "To server %v\n", str)

			c.conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}
	fmt.Fprintln(c.out, "Finished Sender")
	return nil
}

func (c *telnetClient) Receive() error {
	scanner := bufio.NewScanner(c.conn)
	defer c.Close()
OUTER:
	for {
		select {
		case <-c.ctx.Done():
			fmt.Fprintln(c.out, "Receiver DONE")
			break OUTER
		default:
			if !scanner.Scan() {
				fmt.Fprintln(c.out, "CANNOT SCAN")
				if err := scanner.Err(); err != nil {
					fmt.Fprintf(c.out, "Receiver Scan Error: %v\n", err)
				}
				break OUTER
			}
			text := scanner.Text()

			fmt.Fprintf(c.out, "From server: %s\n", text)
		}
	}
	fmt.Fprintln(c.out, "Finished Receiver")
	return nil
}
