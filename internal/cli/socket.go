package cli

import (
	"bufio"
	"log"
	"strconv"

	"clipboard/internal/common/subcmds"
	"clipboard/internal/socket/packet"
)

func Copy(reader *bufio.Reader, writer *bufio.Writer, text string) {
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
}

func Paste(reader *bufio.Reader, writer *bufio.Writer) string {
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

	return respData.Message
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
	Code    int
	Message string
}
