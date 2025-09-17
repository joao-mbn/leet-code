package math

import go_math "math"

func IsPalindrome(x int) bool {
	return isPalindrome(x)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	xfloat := float64(x)
	highestPower := int(go_math.Log10(xfloat))
	highMultiplier := int(go_math.Pow10(highestPower))
	highLastRemainder := x

	lowMultiplier := 10
	lowLastRemainder := 0

	for highMultiplier >= lowMultiplier {
		highDecimal := highLastRemainder / highMultiplier

		lowRemainder := x % lowMultiplier
		lowDecimal := (lowRemainder - lowLastRemainder) / (lowMultiplier / 10)

		if highDecimal != lowDecimal {
			return false
		}

		highLastRemainder %= highMultiplier
		highMultiplier /= 10

		lowMultiplier *= 10
		lowLastRemainder = lowRemainder
	}

	return true
}
