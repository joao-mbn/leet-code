// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/?envType=study-plan-v2&envId=top-interview-150

package tree

func SortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *TreeNode {
	size := len(nums)
	switch size {
	case 0:
		return nil
	default:
		left := sortedArrayToBST(nums[:size/2])
		right := sortedArrayToBST(nums[min(size, size/2+1):])

		return &TreeNode{
			Val:   nums[size/2],
			Left:  left,
			Right: right,
		}
	}
}
