package server

import (
	"bufio"
	"log"
	"net"

	"clipboard/internal/cli/provider"
	"clipboard/internal/common/subcmds"
	"clipboard/internal/socket/packet"
)

func Handler(conn net.Conn, reader *bufio.Reader, writer *bufio.Writer) {
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
	switch data.SubCommand {
	case subcmds.Copy:
		bufferText := data.Value
		if bufferText == "" {
			return
		}
		provider.TermuxClipboardSet(bufferText)
		packet.TryWriteBlock(writer, Success)
		packet.TrySendWriten(writer)
		logln(subcmds.Paste, bufferText)
	case subcmds.Paste:
		packet.TryWriteBlock(writer, Success)
		bufferText := provider.TermuxClipboardGet()
		packet.TryWrite(writer, bufferText)
		packet.TrySendWriten(writer)
		logln(subcmds.Paste, bufferText)
	}
}

func logln(subCommand subcmds.SubCommand, text string) {
	log.Println(string(subCommand) + ": " + text + "\x00\t\x00\n\x00\n")
}
