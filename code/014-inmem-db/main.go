package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	data := make(map[string]string)
	for scanner.Scan() {
		fs := strings.Fields(scanner.Text())

		switch fs[0] {
		case "GET":
			key := fs[1]
			value := data[key]
			fmt.Fprintf(conn, "%s\n", value)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE")
			}
			key := fs[1]
			value := fs[2]
			data[key] = value

		case "DEL":
			key := fs[1]
			delete(data, key)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND", fs[0])
		}
	}
}
