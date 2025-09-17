// https://leetcode.com/problems/candy/?envType=study-plan-v2&envId=top-interview-150

package dynamicprogramming

func candy(ratings []int) int {
	size := len(ratings)

	forwardStreak := make([]int, size)
	forwardStreak[0] = 1
	for i := 1; i < size; i++ {
		if ratings[i] > ratings[i-1] {
			forwardStreak[i] = forwardStreak[i-1] + 1
		} else {
			forwardStreak[i] = 1
		}
	}

	backwardStreak := make([]int, size)
	backwardStreak[size-1] = 1
	for i := size - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			backwardStreak[i] = backwardStreak[i+1] + 1
		} else {
			backwardStreak[i] = 1
		}
	}

	totalCandies := 0

	for i := 0; i < size; i++ {
		rating := ratings[i]
		candies := 1

		if i > 0 && rating > ratings[i-1] {
			candies = forwardStreak[i-1] + 1
		}

		if i < size-1 && rating > ratings[i+1] {
			candies = max(candies, backwardStreak[i+1]+1)
		}

		totalCandies += candies
	}

	return totalCandies
}
