package chapter1

import (
	"strconv"
	"strings"
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

// BetterCompressString is a better version of CompressString
func BetterCompressString(s string) string {
	var output strings.Builder
	var count int
	var curChar = s[0]

	for i := range s {
		if s[i] == curChar {
			count++
			continue
		}
		output.WriteByte(curChar)
		output.WriteString(strconv.Itoa(count))
		count = 1
		curChar = s[i]
	}

	output.WriteByte(curChar)
	output.WriteString(strconv.Itoa(count))

	if output.Len() >= len(s) {
		return s
	}
	return output.String()
}

// BetterBetterCompressString is a better version of BetterCompressString
func BetterBetterCompressString(s string) string {
	var output strings.Builder
	var count int
	var curChar = s[0]

	for _, v := range s {
		if byte(v) == curChar {
			count++
			continue
		}
		output.WriteByte(curChar)
		output.WriteString(strconv.Itoa(count))
		count = 1
		curChar = byte(v)
	}

	output.WriteByte(curChar)
	output.WriteString(strconv.Itoa(count))

	if output.Len() >= len(s) {
		return s
	}
	return output.String()
}
