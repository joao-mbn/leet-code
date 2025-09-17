// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/?envType=study-plan-v2&envId=top-interview-150

package binarytreegeneral

import (
	"leet-code/tree"
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *tree.TreeNode {
	root := &tree.TreeNode{
		Val: preorder[0],
	}

	if len(preorder) == 1 {
		return root
	}

	rootIndex := slices.Index(inorder, preorder[0])
	if len(preorder[1:rootIndex+1]) > 0 {
		root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	}

	if len(preorder[rootIndex+1:]) > 0 {
		root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	}

	return root
}

/*
1. Find index of inorder equal to first element of preorder
2. For inorder: Get slice of inorder to the left of index as left, the other as right
3. For preorder: Get N elements past the first, where N = index - 1 as left, the other as right
4. Are there two elements or less ? Build and return the tree : return call to step 1 with slices

D = 0

Inorder
4, 2, 9, 5, 1, 6, 3, 15, 7

Preorder
1, 2, 4, 5, 9, 3, 6, 7, 15

--------------------------------
D = 1

Inorder
4, 2, 9, 5

Preorder
2, 4, 5, 9

--------------------------------
D = 2

Inorder
4, 2

Preorder
2, 4

--------------------------------

*/
