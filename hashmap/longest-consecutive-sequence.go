package hashmap

func LongestConsecutiveSequence(nums []int) int {
	return longestConsecutive(nums)
}

func longestConsecutive(nums []int) int {
	uniqueElements := map[int]bool{}
	biggestSequence := 0

	for _, num := range nums {
		uniqueElements[num] = true
	}

	for len(uniqueElements) > 0 {
		seed := 0
		for key := range uniqueElements {
			seed = key
			break
		}
		delete(uniqueElements, seed)

		currentSequence := 1
		sequenceTail := seed
		sequenceHead := seed
		okTail, okHead := true, true

		for okHead || okTail {
			_, okTail = uniqueElements[sequenceTail-1]
			if okTail {
				sequenceTail--
				currentSequence++
				delete(uniqueElements, sequenceTail)
			}

			_, okHead = uniqueElements[sequenceHead+1]
			if okHead {
				sequenceHead++
				currentSequence++
				delete(uniqueElements, sequenceHead)
			}
		}

		biggestSequence = max(currentSequence, biggestSequence)
	}

	return biggestSequence
}
