package stack

import "strconv"

func Calculate(s string) int {
	return calculate(s)
}

const (
	openParens = iota
	closeParens
	plus
	minus
	number
)

func calculate(s string) int {
	tokens := lexer(s)
	return evaluate(tokens)
}

func evaluate(tokens [][2]int) int {
	stack := [][2]int{}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		op := token[0]

		switch op {
		case openParens:
			closeIndex := i + 1 + findClosingParens(tokens[i+1:])
			parensResult := evaluate(tokens[i+1 : closeIndex])
			stack = append(stack, [2]int{number, parensResult})
			i = closeIndex
		case closeParens:
			continue
		case number, plus, minus:
			stack = append(stack, token)
		}

		if len(stack) < 2 {
			continue
		}

		if len(stack) > 3 {
			panic("stack is too big with 4+ elements")
		}

		if stack[0][0] == minus {
			if len(stack) > 2 {
				panic("stack has accumulated too many elements for a prefix negative")
			}

			if stack[1][0] != number {
				panic("negating non-number on prefix minus")
			}

			stack = [][2]int{{number, -1 * stack[1][1]}}
		}

		if len(stack) == 3 {
			if stack[0][0] != number {
				panic("stack has three elements and the first one is not a number")
			}
			if stack[2][0] != number {
				panic("stack has three elements and the third one is not a number")
			}
			if stack[1][0] != minus && stack[1][0] != plus {
				panic("stack has three elements and the second one is not a infix operator")
			}

			if stack[1][0] == plus {
				stack = [][2]int{{number, stack[0][1] + stack[2][1]}}
			} else {
				stack = [][2]int{{number, stack[0][1] - stack[2][1]}}
			}
		}
	}

	return stack[0][1]
}

func findClosingParens(tokens [][2]int) int {
	openCount := 1
	for i := range tokens {
		token := tokens[i]
		op := token[0]

		switch op {
		case openParens:
			openCount++
		case closeParens:
			openCount--
		}

		if openCount == 0 {
			return i
		}
	}

	panic("closing parens out of bounds")
}

func lexer(s string) [][2]int {
	tokens := [][2]int{}

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char < '0' || char > '9' {
			switch char {
			case '(':
				tokens = append(tokens, [2]int{openParens})
			case ')':
				tokens = append(tokens, [2]int{closeParens})
			case '+':
				tokens = append(tokens, [2]int{plus})
			case '-':
				tokens = append(tokens, [2]int{minus})
			}
			continue
		}

		token := ""
		for char >= '0' && char <= '9' {
			token += string(char)

			i++
			if i < len(s) {
				char = s[i]
			} else {
				break
			}
		}
		value, _ := strconv.Atoi(token)
		tokens = append(tokens, [2]int{number, value})
		i--
	}

	return tokens
}
