// Package main is the entry point for the program'.'
package main

import (
	"fmt"
	"leet-code/hashmap"
)

func main() {

	for _, input := range []struct {
		ransomNote string
		magazine   string
	}{
		{ransomNote: "fihjjjjei", magazine: "hjibagacbhadfaefdjaeaebgi"},
	} {
		result := hashmap.CanConstruct(input.ransomNote, input.magazine)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input, result)
	}
}
