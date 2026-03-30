func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	pair := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, char := range s {
		if opening, ok := pair[char]; ok {
			if len(stack) == 0 || opening != stack[len(stack) - 1] {
				return false
			}
			stack = stack[:len(stack) -1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
