// Package checks
package checks

import (
	"log"

	"clipboard/internal/cli/flags"
	"clipboard/internal/cli/values"
	"clipboard/internal/cli/view"
	"clipboard/internal/common/subcmds"
	"clipboard/pkg/cli"
)

func TryHandleArgsNumber(subCommand subcmds.SubCommand) {
	switch subCommand {
	case subcmds.Copy:
		if len(cli.Args) == 1 {
			log.Fatalln(view.MissingArgumentsText(values.Requireds[subcmds.Copy]...))
		} else if len(cli.Args) > 4 {
			log.Fatalln(view.ExtraArgumentsText(cli.Args[1 : len(cli.Args)-1]...))
		}
	case subcmds.Paste:
		if len(cli.Args) > 3 {
			log.Fatalln(view.ExtraArgumentsText(cli.Args[1:]...))
		}
	}
}

func TryHandleFlag() string {
	if cli.Args[1] != string(flags.File) {
		log.Fatalln(view.UndefinedFlagText(cli.Args[1]))
	}
	bufferPath := cli.Args[2]
	return bufferPath
}
