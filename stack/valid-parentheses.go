package stack

func IsValid(s string) bool {
	return isValid(s)
}

func isValid(s string) bool {
	stack := []rune{}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		stackTop := stack[len(stack)-1]
		if (char == ')' && stackTop != '(') ||
			(char == ']' && stackTop != '[') ||
			(char == '}' && stackTop != '{') {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
