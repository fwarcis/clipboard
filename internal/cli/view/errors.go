package view

import (
	"fmt"
)

func ExtraArgumentsText(extras ...string) string {
	return ErrorText(fmt.Sprintf("extra arguments: %s", extras))
}

func MissingArgumentsText(missings ...string) string {
	return ErrorText(fmt.Sprintf("missing arguments: %s", missings))
}

func UndefinedSubCommandText(text string) string {
	return ErrorText(fmt.Sprintf("undefined subcommand: %s", text))
}

func UndefinedFlagText(text string) string {
	return ErrorText(fmt.Sprintf("undefined flag: %s", text))
}

func ErrorText(text string) string {
	return fmt.Sprintf("%s: error: %s", ProgramName, text)
}
