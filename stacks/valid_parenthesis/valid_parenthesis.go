package validparenthesis

// Solution for https://leetcode.com/problems/valid-parentheses
func IsValid(s string) bool {
	stack := []rune{}

	for _, char := range s {
		switch char {
		case '(':
			stack = append(stack, '(')

		case '{':
			stack = append(stack, '{')

		case '[':
			stack = append(stack, '[')

		case ')':
			stack = checkOrPop(stack, '(')
			if stack == nil {
				return false
			}

		case '}':
			stack = checkOrPop(stack, '{')
			if stack == nil {
				return false
			}

		case ']':
			stack = checkOrPop(stack, '[')
			if stack == nil {
				return false
			}

		}
	}

	return len(stack) == 0
}

func checkOrPop(stack []rune, char rune) []rune {
	if len(stack) == 0 {
		return nil
	}

	last := stack[len(stack)-1]
	if last != char {
		return nil
	}

	return stack[:len(stack)-1]
}
