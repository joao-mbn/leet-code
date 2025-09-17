// https://leetcode.com/problems/merge-intervals/?envType=study-plan-v2&envId=top-interview-150

package intervals

func Merge(intervals [][]int) [][]int {
	return merge(intervals)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	intervals = sort(intervals)

	for i := 1; i < len(intervals); i++ {
		prevInterval := intervals[i-1]
		currInterval := intervals[i]

		if currInterval[0] <= prevInterval[1] {
			mergedInterval := []int{prevInterval[0], max(currInterval[1], prevInterval[1])}
			intervals[i-1] = mergedInterval
			intervals = append(intervals[0:i], intervals[i+1:]...)
			i--
		}
	}

	return intervals
}

func sort(intervals [][]int) [][]int {
	size := len(intervals)

	if size > 2 {
		sortedLeft := sort(intervals[0:(size / 2)])
		sortedRight := sort(intervals[size/2:])

		leftPointer := 0
		rightPointer := 0
		pointerSum := leftPointer + rightPointer
		sortedIntervals := make([][]int, size)

		for pointerSum < size {
			if rightPointer >= len(sortedRight) ||
				(leftPointer < len(sortedLeft) && sortedLeft[leftPointer][0] < sortedRight[rightPointer][0]) {
				sortedIntervals[pointerSum] = sortedLeft[leftPointer]
				leftPointer++
				pointerSum++
			} else {
				sortedIntervals[pointerSum] = sortedRight[rightPointer]
				rightPointer++
				pointerSum++
			}
		}

		return sortedIntervals
	} else if size == 2 {
		if intervals[0][0] > intervals[1][0] {
			return [][]int{intervals[1], intervals[0]}
		}
	}

	return intervals
}
