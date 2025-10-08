package twopointers

func MaxArea(height []int) int {
	return maxArea(height)
}

func maxArea(height []int) int {
	biggestArea := 0
	left := 0
	right := len(height) - 1

	for left != right {
		area := (right - left) * min(height[left], height[right])
		biggestArea = max(biggestArea, area)

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return biggestArea
}
