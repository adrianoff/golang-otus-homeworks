package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

var (
	timeout       time.Duration
	ErrConnection = errors.New("connection error")
)

const defaultTimeout = 10

func init() {
	flag.DurationVar(&timeout, "timeout", time.Second*defaultTimeout, "connection timeout [default=10s]")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	host := args[0]
	port := args[1]

	address := net.JoinHostPort(host, port)
	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)

	if err := run(client); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(client TelnetInterface) error {
	if err := client.Connect(); err != nil {
		return ErrConnection
	}
	defer func(client TelnetInterface) {
		err := client.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Close connection error\n")
		}
	}(client)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	go func() {
		defer stop()
		if err := client.Send(); err != nil {
			log.Fatalf("Error sending data. Error: %s", err)
		}
		fmt.Fprintf(os.Stderr, "EOF\n")
	}()

	go func() {
		defer stop()
		if err := client.Receive(); err != nil {
			log.Fatalf("Error receiving data. Error: %s", err)
		}
		fmt.Fprintf(os.Stderr, "Connection was closed by peer\n")
	}()

	<-ctx.Done()
	return nil
}
