package main

import (
	"bufio"
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
		doneCh:  make(chan struct{}),
	}
}

type telnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	doneCh  chan struct{}
	conn    net.Conn
}

func (c *telnetClient) Connect() error {
	var err error
	c.conn, err = net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return err
	}
	return nil
}

func (c *telnetClient) Close() error {
	fmt.Fprintln(c.out, "Client is Closed.")
	close(c.doneCh)
	return c.conn.Close()
}

func (c *telnetClient) Send() error {
	scanner := bufio.NewScanner(c.in)
OUTER:
	for {
		select {
		case <-c.doneCh:
			fmt.Fprintln(c.out, "Sender DONE")
			break OUTER
		default:
			fmt.Fprintln(c.out, "try scan (Send)")
			if !scanner.Scan() {
				fmt.Fprintln(c.out, "Not Scan. EOF.")
				if err := scanner.Err(); err != nil {
					fmt.Fprintf(c.out, "Sender Scan Error: %v\n", err)
				}
				break OUTER
			}
			fmt.Fprintln(c.out, "pre scan text")
			str := scanner.Text()
			fmt.Fprintf(c.out, "To server %v\n", str)

			_, err := c.conn.Write([]byte(fmt.Sprintf("%s\n", str)))
			if err != nil {
				fmt.Fprintf(c.out, "Sender Write Error: %v\n", err)
			}
		}
	}
	fmt.Fprintln(c.out, "Finished Sender")
	return nil
}

func (c *telnetClient) Receive() error {
	scanner := bufio.NewScanner(c.conn)
OUTER:
	for {
		select {
		case <-c.doneCh:
			fmt.Fprintln(c.out, "Receiver DONE")
			break OUTER
		default:
			fmt.Fprintln(c.out, "try scan (Receive)")
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
