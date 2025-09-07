// Package main is the entry point for the program'.'
package main

import "fmt"

func main() {
	for _, input := range []struct {
		root *TreeNode
	}{
		{root: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		},
	} {
		result := getMinimumDifference(input.root)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input, result)
	}
}
