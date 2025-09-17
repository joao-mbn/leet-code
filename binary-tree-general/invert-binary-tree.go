// https://leetcode.com/problems/invert-binary-tree/?envType=study-plan-v2&envId=top-interview-150

package binarytreegeneral

import "leet-code/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func InvertTree(root *tree.TreeNode) *tree.TreeNode {
	return invertTree(root)
}

func invertTree(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return root
	}

	var prevLeft *tree.TreeNode
	if root.Left != nil {
		prevLeftVal := *root.Left
		prevLeft = &prevLeftVal
	}
	root.Left = invertTree(root.Right)
	root.Right = invertTree(prevLeft)

	return root
}
