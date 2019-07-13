package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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

	// Set timeout
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	// Scan for incoming text
	// Exmaple: $ nc localhost 8080
	//          $ something
	// response:$ I heard you say something
	//
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say %s\n", ln)
	}
	defer conn.Close()

	// Wont get here until connection drops
	fmt.Println("Connection disconnected")
}
