package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// TODO: write server program to handle concurrent client connections.
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("No Problem.Continue")
			continue

		}

		go handleConn(conn)
	}

}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server-Shailesh\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
