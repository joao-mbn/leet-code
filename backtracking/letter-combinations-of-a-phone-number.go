// https://leetcode.com/problems/letter-combinations-of-a-phone-number/?envType=study-plan-v2&envId=top-interview-150

package backtracking

var letters = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func LetterCombinations(digits string) []string {
	return letterCombinations(digits)
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	lettersOfFirstDigit := letters[string(digits[0])]
	if len(digits) == 1 {
		return lettersOfFirstDigit
	}

	combinations := []string{}
	innerResults := letterCombinations(digits[1:])

	for _, result := range innerResults {
		for _, letter := range lettersOfFirstDigit {
			combinations = append(combinations, letter+result)
		}
	}

	return combinations
}
