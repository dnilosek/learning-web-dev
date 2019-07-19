package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handle(c net.Conn) {
	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		log.Println(scanner.Text())
	}

	fmt.Fprintln(c, "You are connected")
	defer c.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
