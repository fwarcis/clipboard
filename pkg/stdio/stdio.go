package stdio

import (
	"io"
	"log"
	"os"
)

func RequireInput() string {
	input, err := io.ReadAll(os.Stdin)
	if err != io.EOF && err != nil {
		log.Fatalln(err.Error())
	}
	return string(input)
}
