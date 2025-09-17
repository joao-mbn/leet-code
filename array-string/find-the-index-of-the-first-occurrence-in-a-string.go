// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/?envType=study-plan-v2&envId=top-interview-150

package arraystring

func StrStr(haystack string, needle string) int {
	return strStr(haystack, needle)
}

func strStr(haystack string, needle string) int {
	sizeHaystack := len(haystack)
	sizeNeedle := len(needle)

	if sizeNeedle > sizeHaystack {
		return -1
	}

	if sizeHaystack == sizeNeedle && haystack != needle {
		return -1
	}

	windowLeft := 0
	windowRight := sizeNeedle
	for windowRight <= sizeHaystack {
		window := haystack[windowLeft:windowRight]
		if window == needle {
			return windowLeft
		}

		windowLeft++
		windowRight++
	}

	return -1
}
