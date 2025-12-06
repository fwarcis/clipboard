package cli

import (
	"io"
	"log"
	"os"

	"clipboard/internal/common/subcmds"
	cliutils "clipboard/pkg/cli-utils"
	"clipboard/pkg/fsys"
)

func OnRegularFile(subCommand subcmds.SubCommand, path string) {
	switch subCommand {
	case subcmds.Copy:
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, fsys.URDWR)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer file.Close()

		text := cliutils.LastArg
		file.WriteString(text)
	case subcmds.Paste:
		file, err := os.OpenFile(path, os.O_RDONLY, fsys.URDWR)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer file.Close()

		text, err := io.ReadAll(file)
		if err != nil {
			log.Fatalln(err.Error())
		}
		print(text)
	}
}
