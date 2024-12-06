package utils

import (
	"regexp"
	"strings"
)

func Dedent(s string) string {
	leadingNewlineRegex, _ := regexp.Compile(`^\s+`)
	leadingNewlineTrimmed := leadingNewlineRegex.ReplaceAllString(s, "")

	trailingNewlineRegex, _ := regexp.Compile(`\s+$`)
	trailingNewlineTrimmed := trailingNewlineRegex.ReplaceAllString(leadingNewlineTrimmed, "")

	lines := []string{}
	for _, line := range strings.Split(trailingNewlineTrimmed, "\n") {
		lines = append(lines, strings.TrimLeft(line, " \t"))
	}

	return strings.Join(lines, "\n")
}
