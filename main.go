// Package main is the entry point for the program'.'
package main

import "fmt"

func main() {
	for _, input := range []struct {
		root *TreeNode
	}{
		{
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
		},
		{
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 20}},
		},
		{
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
		},
		{
			root: &TreeNode{Val: 3},
		},
	} {
		result := averageOfLevels(input.root)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input.root, result)
	}
}
