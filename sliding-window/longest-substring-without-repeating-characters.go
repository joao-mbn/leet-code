package slidingwindow

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}

func lengthOfLongestSubstring(s string) int {
	charStore := map[rune]int{}
	subStart := 0
	maxSub := 0

	for i, char := range s {
		if charIndex, ok := charStore[char]; ok {
			subStart = max(subStart, charIndex+1)
		}
		currentSub := i - subStart + 1
		maxSub = max(maxSub, currentSub)
		charStore[char] = i
	}

	return maxSub
}
