package chapter1

import (
	"fmt"
	"unicode/utf8"
)

// CompressString compresses the string and returns the compressed string OR returns the original if it cannot be compressed
func CompressString(s string) string {
	var output string
	var curChar rune
	var count int
	curChar, _ = utf8.DecodeRuneInString(s)

	for _, v := range s {
		if v == curChar {
			count++
			continue
		} else {
			output += fmt.Sprintf("%s%d", string(curChar), count)
			count = 1
			curChar = v
		}
	}
	output += fmt.Sprintf("%s%d", string(curChar), count)

	if len(output) == len(s) {
		return s
	}

	return output
}
