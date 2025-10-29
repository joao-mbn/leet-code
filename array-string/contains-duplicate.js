/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
  const uniqueElements = new Set()

  for (const num of nums) {
    if (uniqueElements.has(num)) {
      return true
    }

    uniqueElements.add(num)
  }

  return false
};