// https://leetcode.com/problems/invert-binary-tree/?envType=study-plan-v2&envId=top-interview-150

package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var prevLeft *TreeNode
	if root.Left != nil {
		prevLeftVal := *root.Left
		prevLeft = &prevLeftVal
	}
	root.Left = invertTree(root.Right)
	root.Right = invertTree(prevLeft)

	return root
}
