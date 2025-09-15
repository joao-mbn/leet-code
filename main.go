// Package main is the entry point for the program'.'
package main

import (
	"fmt"
	tree "leet-code/Tree"
)

func main() {

	for _, input := range []struct {
		nums []int
	}{
		{nums: []int{-10}},
		{nums: []int{-10, -3}},
		{nums: []int{-10, -3, 0, 5, 9}},
	} {
		result := tree.SortedArrayToBST(input.nums)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input, result)
	}
}
