package hashmap

func CanConstruct(ransomNote string, magazine string) bool {
	return canConstruct(ransomNote, magazine)
}

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := map[rune]int{}

	for _, letter := range magazine {
		if count, ok := magazineMap[letter]; ok {
			magazineMap[letter] = count + 1
		} else {
			magazineMap[letter] = 1
		}
	}

	for _, letter := range ransomNote {
		if count, ok := magazineMap[letter]; ok {
			count--
			magazineMap[letter] = count
			if count == 0 {
				delete(magazineMap, letter)
			}
		} else {
			return false
		}
	}

	return true
}
