package oneddp

func WordBreak(s string, wordDict []string) bool {
	return wordBreak(s, wordDict)
}

func wordBreak(s string, wordDict []string) bool {
	size := len(s)
	dp := make([]bool, size+1)
	dp[0] = true

	for i := range size + 1 {
		for _, word := range wordDict {
			subStart := i - len(word)
			if subStart < 0 {
				continue
			}

			if dp[subStart] && word == s[subStart:i] {
				dp[i] = true
				break
			}
		}
	}

	return dp[size]
}
