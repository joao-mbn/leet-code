// Package main is the entry point for the program'.'
package main

import (
	"fmt"
	binarytreegeneral "leet-code/binary-tree-general"
	treenode "leet-code/tree"
)

func main() {

	for _, input := range []struct {
		root *treenode.TreeNode
	}{
		{root: &treenode.TreeNode{
			Val: 1,
			Left: &treenode.TreeNode{
				Val:   2,
				Right: &treenode.TreeNode{Val: 4},
			},
			Right: &treenode.TreeNode{Val: 5},
		}},
	} {
		binarytreegeneral.Flatten(input.root)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input)
	}
}
