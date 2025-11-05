/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortArray = function(nums) {
  return mergeSort(nums)
};

var mergeSort = function(nums) {
  if (nums.length <= 1) {
    return nums
  }

  const halfway = Math.floor(nums.length/2)
  const sortedLeft = mergeSort(nums.slice(0, halfway))
  const sortedRight = mergeSort(nums.slice(halfway))
  const sortedNums = []
  let i = j = 0

  while (i < sortedLeft.length && j < sortedRight.length) {
    if (sortedLeft[i] < sortedRight[j]) {
      sortedNums.push(sortedLeft[i])
      i++
    } else {
      sortedNums.push(sortedRight[j])
      j++
    }
  }

  sortedNums.push(...sortedLeft.slice(i), ...sortedRight.slice(j))

  return sortedNums
}
