package twopointers

func IsSubsequence(s string, t string) bool {
	return isSubsequence(s, t)
}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	if t == "" {
		return false
	}

	sPos := 0
	for _, tRune := range t {
		if byte(tRune) == s[sPos] {
			sPos++
		}

		if sPos == len(s) {
			return true
		}
	}

	return false
}
