package ispalindrome

import (
	"strings"
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return -1
	}, s)

	return check(s)
}

func check(s string) bool {
	if len(s) <= 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}
	return check(s[1 : len(s)-1])
}
