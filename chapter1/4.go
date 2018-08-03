package chapter1

import "strings"

// Write a method to replace all spaces in a string with '%20'. You may assume that
// the string has sufficient space at the end of the string to hold the additional
// characters and you are given the "true" length of the string.
// EXAMPLE
// Input: 	"Mr John Smith    ", 13
// Output:  "Mr%20John%20Smith"

// EscapeSpace replaces all whitespace chracters with with '%20'
func EscapeSpace(a string) string {
	res := strings.TrimSpace(a)
	res = strings.Replace(a, " ", "%20", -1)
	return res
}
