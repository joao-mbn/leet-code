// https://leetcode.com/problems/house-robber/description/?envType=study-plan-v2&envId=top-interview-150
package oneddp

func Rob(nums []int) int {
	return rob(nums)
}

func rob(nums []int) int {
	prevSum := nums[0]
	if len(nums) == 1 {
		return prevSum
	}

	currSum := max(nums[1])
	if len(nums) == 2 {
		return max(prevSum, currSum)
	}

	for _, num := range nums[2:] {
		newCurrSum := num + prevSum
		prevSum = max(prevSum, currSum)
		currSum = newCurrSum
	}

	return max(prevSum, currSum)
}
