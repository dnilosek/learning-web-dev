package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func handle(c net.Conn) {
	scanner := bufio.NewScanner(c)
	defer c.Close()

	for scanner.Scan() {
		ln := scanner.Text()
		log.Println(ln)
		if ln == "" {
			break
		}
	}

	log.Println("You got here")
	io.WriteString(c, "You are connected\n")
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
