package main

import (
	"bufio"
	"io"
	"net"
	"time"
)

type TelnetClient struct {
	address string
	timeout time.Duration
	in      io.Reader
	out     io.Writer
	conn    net.Conn
}

type TelnetInterface interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

func (tl *TelnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", tl.address, tl.timeout)
	if err != nil {
		return err
	}

	tl.conn = conn

	return nil
}

func (tl *TelnetClient) Close() error {
	if tl.conn != nil {
		return tl.conn.Close()
	}

	return nil
}

func (tl *TelnetClient) Receive() error {
	return tl.handleMessage(tl.conn, tl.out)
}

func (tl *TelnetClient) Send() error {
	return tl.handleMessage(tl.in, tl.conn)
}

func (tl *TelnetClient) handleMessage(from io.Reader, to io.Writer) error {
	scanner := bufio.NewScanner(from)
	for scanner.Scan() {
		_, err := to.Write(append(scanner.Bytes(), '\n'))
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetInterface {
	return &TelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}
