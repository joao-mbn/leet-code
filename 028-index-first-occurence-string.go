package main

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
