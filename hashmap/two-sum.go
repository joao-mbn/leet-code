package hashmap

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {

	pastValues := map[int]int{}

	for i, num := range nums {
		diff := target - num
		secondNumIndex, has := pastValues[diff]

		if has {
			return []int{i, secondNumIndex}
		}

		pastValues[num] = i
	}

	panic("unreachable")
}
