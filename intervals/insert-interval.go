package intervals

func Insert(intervals [][]int, newInterval []int) [][]int {
	return insert(intervals, newInterval)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	overlaps := [][]int{}
	overlapIndex := -1

	for i, interval := range intervals {
		if newInterval[1] < interval[0] && i == 0 {
			return append([][]int{newInterval}, intervals...)
		} else if newInterval[0] > interval[1] && i == len(intervals)-1 {
			return append(intervals, newInterval)
		} else if interval[0] <= newInterval[0] && interval[1] >= newInterval[1] {
			return intervals
		} else if i > 0 && newInterval[0] > intervals[i-1][1] && newInterval[1] < interval[0] {
			newIntervals := make([][]int, len(intervals[:i]))
			copy(newIntervals, intervals[:i])
			newIntervals = append(newIntervals, newInterval)
			newIntervals = append(newIntervals, intervals[i:]...)
			return newIntervals
		} else if newInterval[1] < interval[0] || newInterval[0] > interval[1] {
			continue
		}

		if overlapIndex == -1 {
			overlapIndex = i
		}
		overlaps = append(overlaps, interval)
	}

	if overlapIndex == -1 {
		panic("negative start index")
	}

	newIntervals := intervals[:overlapIndex]
	newIntervals = append(newIntervals, []int{
		min(overlaps[0][0], newInterval[0]),
		max(overlaps[len(overlaps)-1][1], newInterval[1]),
	})
	newIntervals = append(newIntervals, intervals[overlapIndex+len(overlaps):]...)

	return newIntervals
}
