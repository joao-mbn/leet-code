// Package main is the entry point for the program'.'
package main

import (
	"fmt"
	"leet-code/hashmap"
)

func main() {

	for _, input := range []struct {
		s, t string
	}{
		{s: "nagaram", t: "anagram"},
	} {
		result := hashmap.IsAnagram(input.s, input.t)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input, result)
	}
}
