package packet

import (
	"bufio"
	"log"
	"strings"
)


const (
	EndOfBlock byte = 29
	EndOfPacket byte = '\x00'
)

func TryWrite(writer *bufio.Writer, part string) {
	_, err := writer.WriteString(part)
	if err != nil {
		log.Println(err.Error())
	}
}

func TryWriteBlock(writer *bufio.Writer, content string) {
	TryWrite(writer, content+string(EndOfBlock))
}

func TrySendWriten(writer *bufio.Writer) {
	TryWrite(writer, string(EndOfPacket))
	err := writer.Flush()
	if err != nil {
		log.Println(err.Error())
	}
}

func TrySendHeaderAndBody(writer *bufio.Writer, header, body string) {
	TryWriteBlock(writer, header)
	TryWrite(writer, body)
	TrySendWriten(writer)
}

func NextPacket(reader *bufio.Reader) ([]string, error) {
	packet, err := reader.ReadString(EndOfPacket)
	if err != nil {
		return nil, err
	}

	packetContent := packet[:len(packet)-1]

	if !strings.ContainsRune(packetContent, rune(EndOfBlock)) {
		return nil, &NoEndOfBlockError{Content: packetContent}
	} else if packetContent == string(EndOfBlock) {
		return nil, &OnlyEndOfBlockError{Content: packetContent}
	}
	
	packetParts := strings.SplitN(packetContent, string(EndOfBlock), 2)
	if packetParts[0] == "" {
		return packetParts, &NoHeaderError{Content: packetContent}
	}
	return packetParts, nil
}

