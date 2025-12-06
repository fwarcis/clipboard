package server

import (
	"clipboard/internal/common/subcmds"
)

func TryParse(request []string) (*ClientRequestData, error) {
	if !subcmds.Exists(request[0]) {
		return nil, &subcmds.UndefinedSubCommandError{SubCommandText: request[0]}
	}

	return &ClientRequestData{
		SubCommand: subcmds.SubCommand(request[0]),
		Value:      request[1],
	}, nil
}

type ClientRequestData struct {
	SubCommand subcmds.SubCommand
	Value      string
}

const (
	Success = "0"
	Error   = "1"
)
