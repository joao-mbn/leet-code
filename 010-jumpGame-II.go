package main

func jump(nums []int) int {

	if (len(nums) == 1) {
		return 0
	}

	var jumpInternally func(index int) int

	knownPaths := make(map[int]int)

	jumpInternally = func (i int) int {
		value, exists := knownPaths[i]
		if exists {
			return value
		}

		distance := nums[i]

		if (distance + i >= len(nums) - 1) {
			return 1
		}

		if distance == 0 {
			return -1
		}

		pathJump := -1
		for j := distance + i; j > i; j-- {
			shortestPathCandidate := jumpInternally(j)
			if shortestPathCandidate != -1 {
				knownPaths[j] = shortestPathCandidate
				shortestPathCandidate++
			}

			if pathJump == -1 || (shortestPathCandidate != -1 && shortestPathCandidate < pathJump) {
				pathJump = shortestPathCandidate
			}
		}

		return pathJump
	}

	return jumpInternally(0)
}