package cli

import (
	"fmt"
)

func HelpText() string {
	return ""
}

func ExtraArgumentsText(extras ...string) string {
	return ErrorText(fmt.Sprintf("extra arguments: %s", extras))
}

func MissingArgumentsText(missings ...string) string {
	return ErrorText(fmt.Sprintf("missing arguments: %s", missings))
}

func ErrorText(text string) string {
	return fmt.Sprintf("clipboard: error: %s", text)
}
