package utils

import (
	"strings"
)

func SplitFile(s string, fileName string) []string {
	if fileName == "thinkertoy" {
		return strings.Split(s, "\r\n")
	}
	return strings.Split(s, "\n")
}
