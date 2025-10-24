package hashmap

func IsIsomorphic(s string, t string) bool {
	return isIsomorphic(s, t)
}

func isIsomorphic(s string, t string) bool {
	isomorphicMap := map[rune]rune{}
	tMapped := map[rune]bool{}

	for i, char := range s {
		tMap, ok := isomorphicMap[char]
		tChar := rune(t[i])

		if ok && tMap != tChar {
			return false
		}

		if _, tCharMapped := tMapped[tChar]; !ok && tCharMapped {
			return false
		}

		if !ok {
			tMapped[tChar] = true
			isomorphicMap[char] = tChar
		}
	}

	return true
}
