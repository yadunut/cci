package chapter1

// Given 2 strings, write a method to decide if 1 is a permutation of another

// Assumptions taken: Characterset is utf-8. Whitespaces matter in the data set

// createMap creates a map from string
func createMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	return m
}

// PermutateStr returns true if a is a permutation of b
func PermutateStr(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	aMap := createMap(a)
	bMap := createMap(b)

	for i := range aMap {
		if aMap[i] != bMap[i] {
			return false
		}
	}

	return true
}
