package chapter1

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
		} else {
			output += string(curChar)
			output += string(count)
			count = 1
			curChar = v
		}
	}
	output += string(curChar)
	output += string(count)

	if len(output) == len(s) {
		return s
	}
	return output
}
