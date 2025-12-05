package handlers

import (
	"bufio"
	"log"
	"net"

	"clipboard/internal/common/subcmds"
	"clipboard/internal/socket/packet"
	"clipboard/pkg/cli"
)

func OnSocket(subCommand subcmds.SubCommand, path string) {
	conn, err := net.Dial("unix", path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	switch subCommand {
	case subcmds.Copy:
		text := cli.LastArg
		request := string(subcmds.Copy) + " " + text
		packet.TryAddTo(writer, request)
		packet.TrySend(writer)
	case subcmds.Paste:
		reader := bufio.NewReader(conn)

		request := string(subcmds.Paste)
		packet.TryAddTo(writer, request)
		packet.TrySend(writer)
		text := packet.TryGetContent(reader)

		print(text)
	}
}
