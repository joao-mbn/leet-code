// https://leetcode.com/problems/minimum-size-subarray-sum/?envType=study-plan-v2&envId=top-interview-150

package main

func minSubArrayLen(target int, nums []int) int {
	minLength := 0
	track := 0
	sum := 0

	for i, curr := range nums {
		if curr >= target {
			return 1
		}

		sum += curr
		for sum >= target {
			if minLength == 0 {
				minLength = i - track + 1
			} else {
				minLength = min(i-track+1, minLength)
			}

			if minLength == 1 {
				return 1
			}

			sum -= nums[track]
			track++
		}
	}

	return minLength
}
