package validanagram

// https://leetcode.com/problems/valid-anagram
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := make(map[rune]int)

	for _, c := range s {
		if _, ok := chars[c]; !ok {
			chars[c] = 1
			continue
		}
		chars[c]++
	}

	for _, c := range t {
		v, ok := chars[c]
		if !ok || v <= 0 {
			return false
		}

		chars[c]--
	}

	for _, v := range chars {
		if v > 0 {
			return false
		}
	}

	return true
}
