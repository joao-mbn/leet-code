package hashmap

import (
	"math"
	"strconv"
)

func IsHappy(n int) bool {
	return isHappy(n)
}

func isHappy(n int) bool {
	prevResults := map[int]bool{}

	for n != 1 {
		sum := 0
		for _, char := range strconv.Itoa(n) {
			sum += int(math.Pow(float64(char-'0'), 2))
		}

		if sum == 1 {
			return true
		}

		if _, ok := prevResults[sum]; ok {
			return false
		}
		prevResults[sum] = true
		n = sum
	}

	return true
}
