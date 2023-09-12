package palindromepermutation

import (
	"regexp"
	"strings"
)

/*
	Given a string, write a function to check if it is a permutation of a palindrome.
	A palindrome is a word or phrase that is the same forwards and backwards.
	A permutation is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
*/
func PalindomePermutation(s string) bool {
	wc := make(map[rune]int)
	s = strings.ToLower(s)
	for _, v := range s {
		if match, _ := regexp.MatchString("[a-z]", string(v)); !match {
			continue
		}
		wc[v]++
	}

	sum := 0

	for _, occurrences := range wc {
		sum += occurrences
	}

	return sum == 1 || sum == 2 || sum%2 != 0
}
