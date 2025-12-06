package cli

import (
	"fmt"
)

func ExtraArgumentsText(extras ...string) string {
	return ErrorText(fmt.Sprintf("extra arguments: %s", extras))
}

func MissingArgumentsText(missings ...string) string {
	return ErrorText(fmt.Sprintf("missing arguments: %s", missings))
}

func UndefinedFlagText(text string) string {
	return ErrorText(fmt.Sprintf("undefined flag: %s", text))
}

func ErrorText(text string) string {
	return fmt.Sprintf("clipboard: error: %s", text)
}
