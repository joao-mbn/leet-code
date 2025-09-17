// https://leetcode.com/problems/length-of-last-word/?envType=study-plan-v2&envId=top-interview-150

package arraystring

import "unicode"

func lengthOfLastWord(s string) int {

	last := 0
	beginning := 0

	for i := range s {
		isSpace := unicode.IsSpace(rune(s[i]))
		wasSpace := i > 0 && unicode.IsSpace(rune(s[i-1]))

		if !isSpace {
			if wasSpace {
				beginning = i
			}

			last = i + 1 - beginning
		}
	}

	return last
}
