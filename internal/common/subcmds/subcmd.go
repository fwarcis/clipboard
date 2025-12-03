// Package subcmds
package subcmds

import "slices"

type SubCommand string

const (
	Copy  SubCommand = "copy"
	Paste SubCommand = "paste"
)

var List = []SubCommand{Copy, Paste}

func Exists(text string) bool {
	return slices.Contains(List, SubCommand(text))
}
