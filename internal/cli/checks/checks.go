// Package checks
package checks

import (
	"log"

	"clipboard/internal/cli"
	"clipboard/internal/cli/flags"
	"clipboard/internal/cli/values"
	"clipboard/internal/common/subcmds"
	cliutils "clipboard/pkg/cli-utils"
)

func RequireLengthOfArguments(subCommand subcmds.SubCommand) {
	switch subCommand {
	case subcmds.Copy:
		if len(cliutils.Args) == 1 {
			log.Fatalln(cli.MissingArgumentsText(values.Requireds[subcmds.Copy]...))
		} else if len(cliutils.Args) > 4 {
			log.Fatalln(cli.ExtraArgumentsText(cliutils.Args[1 : len(cliutils.Args)-1]...))
		}
	case subcmds.Paste:
		if len(cliutils.Args) > 3 {
			log.Fatalln(cli.ExtraArgumentsText(cliutils.Args[1:]...))
		}
	}
}

func RequireFileFlag() string {
	if cliutils.Args[1] != string(flags.File) {
		log.Fatalln(cli.UndefinedFlagText(cliutils.Args[1]))
	}
	bufferPath := cliutils.Args[2]
	return bufferPath
}
