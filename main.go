// Package main is the entry point for the program'.'
package main

import "fmt"

func main() {
	for _, input := range []struct {
		strs []string
	}{
		{strs: []string{""}},
		{strs: []string{"a"}},
		{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
	} {
		result := groupAnagrams(input.strs)
		fmt.Printf("Given the slice %v, the grouped anagrams are: %v\n", input.strs, result)
	}
}
