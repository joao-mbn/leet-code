package binarysearch

func FindPeakElement(nums []int) int {
	return findPeakElement(nums)
}

func findPeakElement(nums []int) int {
	return divide(nums, 0)
}

func divide(nums []int, i int) int {
	size := len(nums)
	if size == 1 {
		return i
	}

	split := (size - 1) / 2
	midPoint := nums[split]

	if split == 0 {
		if midPoint > nums[split+1] {
			return i
		}

		return i + 1
	}

	if midPoint > nums[split+1] && midPoint > nums[split-1] {
		return i + split
	}

	if midPoint < nums[split-1] {
		return divide(nums[:split], i)
	}

	return divide(nums[split+1:], i+split+1)
}
