// Package main is the entry point for the program'.'
package main

import (
	"fmt"
	math "leet-code/math"
)

func main() {

	for _, input := range []struct {
		num int
	}{
		{num: 4554},
		{num: 100000001},
		{num: 1000000001},
		{num: 45654},
		{num: 45},
		{num: 121},
		{num: 1234},
		{num: 4},
		{num: -9},
		{num: -121},
	} {
		result := math.IsPalindrome(input.num)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input, result)
	}
}
