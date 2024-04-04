package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/pkg/errors"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	listener, err := net.Listen("tcp", "0.0.0.0:6379")

	if err != nil {
		return errors.Wrap(err, "Failed to bind to port 6379")
	}

	defer closeConn(listener, &err, "Close listener")

	fmt.Println("Server is listening on port 6379")

	conn, err := listener.Accept()
	if err != nil {
		return errors.Wrap(err, "Error accepting connection")
	}

	defer closeConn(conn, &err, "Close connection")

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return errors.Wrap(err, "Error reading from connection")
	}

	fmt.Printf("Received: %s\n", buf[:n])

	_, err = conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		return errors.Wrap(err, "Error writing to connection")
	}

	return nil
}

func closeConn(conn io.Closer, errp *error, msg string) {
	err := conn.Close()
	if errp != nil {
		*errp = errors.Wrap(err, msg)
	}
}
