package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Listening on port 8080
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		// Accept incoming connections
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// Handle connection
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("here")
}
