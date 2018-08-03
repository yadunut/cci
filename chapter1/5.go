package chapter1

import (
	"strconv"
)

// CompressString compresses the string and returns the compressed string OR returns the original if it cannot be compressed
func CompressString(s string) string {
	var output string
	var curChar rune
	var count int
	for _, v := range s {
		curChar = v
		break
	}

	for _, v := range s {
		if v == curChar {
			count++
			continue
		}
		output += string(curChar)
		output += strconv.Itoa(count)
		count = 1
		curChar = v
	}
	output += string(curChar)
	output += strconv.Itoa(count)

	if len(output) >= len(s) {
		return s
	}
	return output
}
