package packet

import (
	"bufio"
	"log"
	"strings"

	netio "clipboard/pkg/net-io"
)

func TryAddTo(writer *bufio.Writer, part string) {
	_, err := writer.WriteString(part)
	if err != nil {
		log.Println(err.Error())
	}
}

func TrySend(writer *bufio.Writer) {
	TryAddTo(writer, string(netio.EndOfPacket))
	err := writer.Flush()
	if err != nil {
		log.Println(err.Error())
	}
}

func TryGetContent(reader *bufio.Reader) string {
	packet, err := reader.ReadString(netio.EndOfPacket)
	if err != nil {
		log.Println(err.Error())
	}
	return packet[:len(packet)-1]
}

func Header(content string) (string, error) {
	splited := splitHeaderAndBody(content)

	if splited == nil {
		return "", &NullContentError{}
	}
	return splited[0], nil
}

func Body(content string) (string, error) {
	splited := splitHeaderAndBody(content)

	if splited == nil {
		return "", &NullContentError{}
	}

	if len(splited) == 1 {
		return "", nil
	}
	return splited[1], nil
}

func splitHeaderAndBody(content string) []string {
	return strings.SplitN(content, " ", 2)
}
