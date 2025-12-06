package cli

import (
	"bufio"
	"log"
	"net"
	"strconv"

	"clipboard/internal/common/subcmds"
	"clipboard/internal/socket/packet"
	cliutils "clipboard/pkg/cli-utils"
)

func OnSocket(subCommand subcmds.SubCommand, path string) {
	conn, err := net.Dial("unix", path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	switch subCommand {
	case subcmds.Copy:
		text := cliutils.LastArg
		packet.TrySendHeaderAndBody(writer,
			 string(subcmds.Copy), text)

		response, err := packet.NextPacket(reader)
		if err != nil {
			log.Fatalln(err.Error())
		}
		respData := tryParse(response)
		if respData.Code != 0 {
			log.Fatalln(respData.Message)
		}
	case subcmds.Paste:
		request := string(subcmds.Paste)
		packet.TryWriteBlock(writer, request)
		packet.TrySendWriten(writer)

		response, err := packet.NextPacket(reader)
		if err != nil {
			log.Fatalln(err.Error())
		}
		respData := tryParse(response)
		if respData.Code != 0 {
			log.Fatalln(respData.Message)
		}

		print(respData.Message)
	}
}

func tryParse(response []string) ServerResponseData {
	code, err := strconv.ParseInt(response[0], 10, 32)
	if err != nil {
		log.Println(err.Error())
	}
	return ServerResponseData{
		Code:    int(code),
		Message: response[1],
	}
}

type ServerResponseData struct {
	Code int
	Message string
}

