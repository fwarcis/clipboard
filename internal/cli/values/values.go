// Package values
package values

import "clipboard/internal/common/subcmds"

var Requireds = map[subcmds.SubCommand][]string{
	subcmds.Copy:  {"<TEXT>"},
	subcmds.Paste: {},
}
