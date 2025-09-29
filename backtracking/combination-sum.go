package backtracking

func CombinationSum(candidates []int, target int) [][]int {
	return combinationSum(candidates, target)
}

func combinationSum(candidates []int, target int) [][]int {
	return findCombinations(candidates, []int{}, target, 0)
}

func findCombinations(candidates, partialCombination []int, target, currentSum int) [][]int {
	combinations := [][]int{}

	for i, candidate := range candidates {
		updatedCurrentSum := currentSum + candidate
		updatedPartialCombination := append(partialCombination, candidate)

		if updatedCurrentSum < target {
			newCombinations := findCombinations(candidates[i:], updatedPartialCombination, target, updatedCurrentSum)
			combinations = append(combinations, newCombinations...)
		} else if updatedCurrentSum == target {
			combinations = append(combinations, append([]int{}, updatedPartialCombination...))
		}
	}

	return combinations
}
