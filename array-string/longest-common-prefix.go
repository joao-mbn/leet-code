// https://leetcode.com/problems/longest-common-prefix/?envType=study-plan-v2&envId=top-interview-150

package arraystring

import "strings"

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := ""
	for i, word := range strs {
		if len(word) == 0 {
			return ""
		}

		if i == 0 {
			prefix = word
			continue
		}

		largestPossibleSize := min(len(word), len(prefix))
		var sb strings.Builder
		for j := range largestPossibleSize {
			if word[j] != prefix[j] {
				break
			}

			sb.WriteByte(word[j])
		}

		if sb.String() == "" {
			return ""
		}

		prefix = sb.String()
	}

	return prefix
}
