//go:build !solution

package spacecollapse

import (
	"unicode"
	"unicode/utf8"
)

func CollapseSpaces(input string) string {
	res := make([]rune, 0, len(input))
	previousCharIsSpace := false

	for _, char := range input {
		if !utf8.ValidRune(char) {
			res = append(res, '\uFFFD')
		} else if unicode.IsSpace(char) && !previousCharIsSpace {
			previousCharIsSpace = true
			res = append(res, ' ')
		} else if !unicode.IsSpace(char) {
			previousCharIsSpace = false
			res = append(res, char)
		}
	}

	return string(res)
}
