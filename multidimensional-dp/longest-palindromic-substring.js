/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
  if (s.length === 1) {
    return s[0]
  }

  let longestPalindromeSubstring = ""
  for (let i = 0; i < s.length; i++) {
    const oddPalindrome = walkPalindrome(s, i, i)
    const evenPalindrome = walkPalindrome(s, i, i+1)

    if (oddPalindrome.length > longestPalindromeSubstring.length) {
      longestPalindromeSubstring = oddPalindrome
    }

    if (evenPalindrome.length > longestPalindromeSubstring.length) {
      longestPalindromeSubstring = evenPalindrome
    }
  }

  return longestPalindromeSubstring
};

/**
 * @param {string} s
 * @param {number} left
 * @param {number} right
 * @return {string}
 */
function walkPalindrome(s, left, right) {
  while (left >= 0 && right < s.length && s[left] === s[right]) {
    left--
    right++
  }

  return s.slice(left+1,right)
}

console.log(longestPalindrome("db"))
console.log(longestPalindrome("ddcb"))
console.log(longestPalindrome("babad"))
console.log(longestPalindrome("cbbd"))
console.log(longestPalindrome("cbdd"))
console.log(longestPalindrome("babccba"))