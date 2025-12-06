package main

import (
	"bufio"
	"log"
	"net"

	"clipboard/internal/cli"
	handlers "clipboard/internal/cli"
	"clipboard/internal/common/subcmds"
	cliutils "clipboard/pkg/cli-utils"
	"clipboard/pkg/stdio"
)

func main() {
	log.SetFlags(0)

	if len(cliutils.Args) == 0 {
		log.Fatalln(cli.HelpText())
	} else if !subcmds.Exists(cliutils.Args[0]) {
		err := subcmds.UndefinedSubCommandError{SubCommandText: cliutils.Args[0]}
		log.Fatalln("clipboard: " + err.Error())
	}

	subCommand := subcmds.SubCommand(cliutils.Args[0])
	switch subCommand {
	case subcmds.Copy:
		if len(cliutils.Args)-1 > subcmds.OnCopyMaxArgs {
			log.Fatalln(cli.ExtraArgumentsText(cliutils.Args[2:len(cliutils.Args)]...))
		}
	case subcmds.Paste:
		if len(cliutils.Args)-1 > subcmds.OnPasteMaxArgs {
			log.Fatalln(cli.ExtraArgumentsText(cliutils.Args[1:]...))
		}
	}

	bufferPath := "@clipboard"
	conn, err := net.Dial("unix", bufferPath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	switch subCommand {
	case subcmds.Copy:
		hasNoTextInArgs := len(cliutils.Args)-1 == 0
		if hasNoTextInArgs {
			text := stdio.RequireInput()
			handlers.Copy(reader, writer, text)
			return
		}
		handlers.Copy(reader, writer, cliutils.Args[1])
	case subcmds.Paste:
		text := handlers.Paste(reader, writer)
		print(text)
	}
}
