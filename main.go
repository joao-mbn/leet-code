// Package main is the entry point for the program.
package main

import "fmt"

func main() {
	for _, input := range [][]string{
		{"flower", "flow", "flight"},
		{"flower", "flow", ""},
		{"", "flow", "flower"},
		{"flow", "flow", "flow"},
		{"flow", "flow", "flower"},
		{},
		{"flower"},
		{"dog", "racecar", "car"},
	} {
		result := longestCommonPrefix(input)
		fmt.Printf("The largest common prefix for input %v is : %v\n", input, result)
	}
}
