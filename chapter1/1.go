package chapter1

// Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?

// UniqueStr returns 1 if there are only unique characters in the string, 0 if there are duplicate characters, and -1 if there are errors
func UniqueStr(s string) int {
	if s == "" {
		return -1
	}
	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}
	for _, v := range m {
		if v > 1 {
			return 0
		}
	}
	return 1
}

// BetterUniqueStr returns false if there are duplicate characters and returns true if there are non
func BetterUniqueStr(s string) bool {
	if s == "" {
		return false
	}

	m := make(map[rune]bool)
	for _, v := range s {
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
}
