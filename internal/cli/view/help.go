package view

import "fmt"

const ProgramName = "clipboard"

func Help() string {
	return fmt.Sprintf("%s:", ProgramName)
}
