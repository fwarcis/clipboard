package main

import (
	"log"

	handlers "clipboard/internal/cli"
	"clipboard/internal/cli/checks"
	"clipboard/internal/cli/view"
	"clipboard/internal/common/subcmds"

	"clipboard/pkg/cli"
)

func main() {
	log.SetFlags(0)

	if len(cli.Args) == 0 {
		log.Fatalln(view.Help())
	} else if !subcmds.Exists(cli.Args[0]) {
		log.Fatalln(view.UndefinedSubCommandText(cli.Args[0]))
	}

	subCommand := subcmds.SubCommand(cli.Args[0])

	checks.RequireLengthOfArguments(subCommand)

	bufferPath := "@clipboard"
	hasFlag := len(cli.Args) == 3 || len(cli.Args) == 4
	if hasFlag {
		bufferPath = checks.RequireFileFlag()
		handlers.OnRegularFile(subCommand, bufferPath)
	}
	handlers.OnSocket(subCommand, bufferPath)
}
