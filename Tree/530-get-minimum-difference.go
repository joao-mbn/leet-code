package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	minDiff := -1
	sortedBST := inorderTraversal(root, []int{})

	for i, curr := range sortedBST[1:] {
		currDiff := abs(curr - sortedBST[i])
		if currDiff == 1 {
			return 1
		} else if minDiff == -1 {
			minDiff = currDiff
		} else {
			minDiff = min(minDiff, currDiff)
		}
	}

	return minDiff
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func inorderTraversal(root *TreeNode, traversal []int) []int {
	if root == nil {
		return traversal
	}

	traversal = inorderTraversal(root.Left, traversal)
	traversal = append(traversal, root.Val)
	traversal = inorderTraversal(root.Right, traversal)

	return traversal
}
