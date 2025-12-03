// Package view
package view

import "fmt"

func UndefinedHeader(header string) string {
	return fmt.Sprintf("undefined header '%s'", header)
}
