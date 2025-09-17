// https://leetcode.com/problems/reverse-words-in-a-string/?envType=study-plan-v2&envId=top-interview-150

package arraystring

import "unicode"

func reverseWords(s string) string {

	beginWordPointer := -1
	reversed := ""

	for i := len(s) - 1; i >= 0; i-- {
		curr := s[i]

		var prev byte
		if i < len(s)-1 {
			prev = s[i+1]
		}

		var next byte
		if i > 0 {
			next = s[i-1]
		}

		currIsSpace := unicode.IsSpace(rune(curr))
		prevIsSpace := unicode.IsSpace(rune(prev))
		nextIsSpace := unicode.IsSpace(rune(next))

		if !currIsSpace && (prevIsSpace || prev == 0) {
			beginWordPointer = i
		}

		if !currIsSpace && (nextIsSpace || next == 0) {
			if reversed != "" {
				reversed += string(" ")
			}

			word := s[i : beginWordPointer+1]
			reversed += word

			beginWordPointer = -1
		}
	}

	return reversed
}
