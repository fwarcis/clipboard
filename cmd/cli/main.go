package main

import (
	"log"

	handlers "clipboard/internal/cli"
	"clipboard/internal/cli/checks"
	"clipboard/internal/cli/view"
	"clipboard/internal/common/subcmds"
	cliutils "clipboard/pkg/cli-utils"
)

func main() {
	log.SetFlags(0)

	if len(cliutils.Args) == 0 {
		log.Fatalln(view.Help())
	} else if !subcmds.Exists(cliutils.Args[0]) {
		err := subcmds.UndefinedSubCommandError{SubCommandText: cliutils.Args[0]}
		log.Fatalln("clipboard: " + err.Error())
	}

	subCommand := subcmds.SubCommand(cliutils.Args[0])

	checks.RequireLengthOfArguments(subCommand)

	bufferPath := "@clipboard"
	hasFlag := len(cliutils.Args) == 3 || len(cliutils.Args) == 4
	if hasFlag {
		bufferPath = checks.RequireFileFlag()
		handlers.OnRegularFile(subCommand, bufferPath)
	}
	handlers.OnSocket(subCommand, bufferPath)
}
