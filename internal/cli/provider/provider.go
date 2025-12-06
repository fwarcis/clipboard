package provider

import (
	"log"
	"os/exec"
)

func TermuxClipboardGet() string {
	output, err := exec.Command("termux-clipboard-get").Output()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return string(output)
}

func TermuxClipboardSet(text string) string {
	output, err := exec.Command("termux-clipboard-set", text).Output()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return string(output)
}
