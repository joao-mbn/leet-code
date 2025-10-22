// Package main is the entry point for the program'.'
package main

import (
	"fmt"
	"leet-code/stack"
)

func main() {

	for _, input := range []struct {
		s string
	}{
		{s: "0"},
		{s: "1"},
		{s: "1 + 1"},
		{s: " 2-1 + 2 "},
		{s: "(1+(4+5+2)-3)+(6+8)"},
	} {
		result := stack.Calculate(input.s)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input, result)
	}
}
