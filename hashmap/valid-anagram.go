package hashmap

func IsAnagram(s string, t string) bool {
	return isAnagram(s, t)
}

func isAnagram(s string, t string) bool {
	codex := map[rune]int{}

	for _, char := range s {
		if count, has := codex[char]; has {
			codex[char] = count + 1
		} else {
			codex[char] = 1
		}
	}

	for _, char := range t {
		if count := codex[char]; count > 1 {
			codex[char] = count - 1
		} else if count == 1 {
			delete(codex, char)
		} else {
			return false
		}
	}

	for range codex {
		return false
	}

	return true
}
