// https://leetcode.com/problems/group-anagrams/?envType=study-plan-v2&envId=top-interview-150

package hashmap

import "fmt"

func GroupAnagrams(strs []string) [][]string {
	return groupAnagrams(strs)
}

func groupAnagrams(strs []string) [][]string {
	anagrams := [][]string{}
	keys := map[string][]string{}

	for _, char := range strs {
		size := len(char)
		sum := 0
		prod := 1

		for _, r := range char {
			sum += int(r)
			prod *= int(r)
		}

		// make the key size, sum and productory to avoid collisions
		key := fmt.Sprintf("%d.%d.%d", size, sum, prod)
		_, ok := keys[key]
		if ok {
			keys[key] = append(keys[key], char)
		} else {
			keys[key] = []string{char}
		}
	}

	for _, value := range keys {
		anagrams = append(anagrams, value)
	}

	return anagrams
}
