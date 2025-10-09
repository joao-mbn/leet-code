package binarytreegeneral

import "leet-code/tree"

type TreeNode = tree.TreeNode

func MaxDepth(root *TreeNode) int {
	return maxDepth(root)
}

func maxDepth(root *TreeNode) int {
	return getDepth(root, 0)
}

func getDepth(root *TreeNode, prevDepth int) int {
	if root == nil {
		return prevDepth
	}

	currDepth := prevDepth + 1

	return max(getDepth(root.Left, currDepth), getDepth(root.Right, currDepth))
}
