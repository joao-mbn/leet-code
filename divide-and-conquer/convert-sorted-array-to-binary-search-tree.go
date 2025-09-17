// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/?envType=study-plan-v2&envId=top-interview-150

package divideandconquer

import "leet-code/tree"

func SortedArrayToBST(nums []int) *tree.TreeNode {
	return sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *tree.TreeNode {
	size := len(nums)
	switch size {
	case 0:
		return nil
	default:
		left := sortedArrayToBST(nums[:size/2])
		right := sortedArrayToBST(nums[min(size, size/2+1):])

		return &tree.TreeNode{
			Val:   nums[size/2],
			Left:  left,
			Right: right,
		}
	}
}
