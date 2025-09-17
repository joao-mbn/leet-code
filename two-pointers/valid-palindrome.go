// https://leetcode.com/problems/valid-palindrome/?envType=study-plan-v2&envId=top-interview-150

package arraystring

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {

	re := regexp.MustCompile(`[^\d\w]|_`)
	s = string(re.ReplaceAll([]byte(s), []byte("")))

	s = strings.ToLower(s)
	size := len(s)

	if size == 0 {
		return true
	}

	halfway := ((size + 1) / 2)
	fmt.Printf("%v \n", halfway)

	for i := range halfway {
		forward := s[i]
		backward := s[size-i-1]

		if forward != backward {
			return false
		}
	}

	return true
}
