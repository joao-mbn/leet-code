// Package main is the entry point for the program'.'
package main

import "fmt"

func main() {
	for _, input := range []struct {
		haystack string
		needle   string
	}{
		{
			haystack: "mississippi",
			needle:   "issip",
		},
		{
			haystack: "sabubut",
			needle:   "but",
		},
		{
			haystack: "sadbutsad",
			needle:   "sad",
		},
		{
			haystack: "sadbutsad",
			needle:   "but",
		},
		{
			haystack: "sadbutbut",
			needle:   "but",
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
		},
		{
			haystack: "sad",
			needle:   "sadbutsad",
		},
	} {
		result := strStr(input.haystack, input.needle)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input, result)
	}
}
