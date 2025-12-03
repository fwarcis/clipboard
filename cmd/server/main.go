package main

import (
	"bufio"
	"log"
	"net"

	"clipboard/internal/common/subcmds"
	"clipboard/internal/server/view"
	"clipboard/internal/socket/packet"
)

func main() {
	listener, err := net.Listen("unix", "@clipboard")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	bufferText := ""

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}

		reader := bufio.NewReader(conn)
		writer := bufio.NewWriter(conn)

		go func() {
			defer conn.Close()

			request := packet.TryGetContent(reader)
			header, err := packet.Header(request)
			if err != nil {
				log.Println(err.Error())
			} else if !subcmds.Exists(header) {
				log.Println(view.UndefinedHeader(header))
			}

			subCommand := subcmds.SubCommand(header)
			switch subCommand {
			case subcmds.Copy:
				body, err := packet.Body(request)
				if err != nil {
					log.Println(err.Error())
				}

				if body == "" {
					return
				}
				bufferText = body
			case subcmds.Paste:
				packet.TryAddTo(writer, bufferText)
				packet.TrySend(writer)
			}
		}()
	}
}
