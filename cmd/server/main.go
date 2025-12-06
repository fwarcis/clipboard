package main

import (
	"bufio"
	"log"
	"net"

	"clipboard/internal/server"
)

func main() {
	listener, err := net.Listen("unix", "@clipboard")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}

		reader := bufio.NewReader(conn)
		writer := bufio.NewWriter(conn)

		go server.Handler(conn, reader, writer)
	}
}
