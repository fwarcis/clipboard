// Package subcmds
package subcmds

import (
	"slices"
)

type SubCommand string

const (
	Copy  SubCommand = "copy"
	Paste SubCommand = "paste"
)

var List = []SubCommand{Copy, Paste}

func Exists(text string) bool {
	return slices.Contains(List, SubCommand(text))
}

type UndefinedSubCommandError struct {
	SubCommandText string
}

func (err *UndefinedSubCommandError) Error() string {
	return "error: undefined subcommand: " + err.SubCommandText
}
