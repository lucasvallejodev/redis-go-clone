package main

import (
	"fmt"
	"io"
	"net"
	"os"
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
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 6379")

	errCh := make(chan error)

	go handleErrors(errCh)

	for {
		fmt.Println("Waiting for client to connect")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			os.Exit(1)
		}

		go handleClient(conn, errCh)
	}
}

func handleClient(conn net.Conn, channel chan error) (err error) {
	response := []byte("+PONG\r\n")
	fmt.Println("Client connected")
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
				return nil
			}
			channel <- err
			return nil
		}

		fmt.Printf("Received %d bytes: %s\n", n, string(buf[:n]))
		conn.Write(response)
	}
}

func handleErrors(errCh <-chan error) {
	for err := range errCh {
		if err != nil {
			fmt.Println("Error handling client:", err)
			os.Exit(1)
		}
	}
}
