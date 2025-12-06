package server

import (
	"bufio"
	"log"
	"net"

	"clipboard/internal/common/subcmds"
	"clipboard/internal/socket/packet"
)

func Handler(conn net.Conn, reader *bufio.Reader, writer *bufio.Writer, bufferStorage *string) {
	defer conn.Close()

	request, err := packet.NextPacket(reader)
	if err != nil {
		packet.TrySendHeaderAndBody(writer, Error, err.Error())
		log.Println(err.Error())
		return
	}

	data, err := TryParse(request)
	if err != nil {
		packet.TrySendHeaderAndBody(writer, Error, err.Error())
		log.Println(err.Error())
		return
	}
	subCommand := subcmds.SubCommand(data.SubCommand)
	switch subCommand {
	case subcmds.Copy:
		if data.Value == "" {
			return
		}
		*bufferStorage = data.Value
		packet.TryWriteBlock(writer, Success)
		packet.TrySendWriten(writer)
	case subcmds.Paste:
		packet.TryWriteBlock(writer, Success)
		packet.TryWrite(writer, *bufferStorage)
		packet.TrySendWriten(writer)
	}
}
