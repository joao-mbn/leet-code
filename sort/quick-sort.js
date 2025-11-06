/**
 * @param {number[]} nums
 * @return {number[]}
 */
var quickSort = function(nums, leftWall = 0, pivotIndex = nums.length - 1) {
  if (leftWall >= pivotIndex) {
    return
  }

  let swapPointer = leftWall - 1
  let pivot = nums.at(pivotIndex)

  for (let i = leftWall; i <= pivotIndex; i++) {
    const element = nums[i];

    if (element > pivot) {
      continue
    }

    swapPointer++
    if (swapPointer === i) {
      continue
    }

    nums[i] = nums[swapPointer]
    nums[swapPointer] = element
  }

  if (swapPointer > 0) {
    quickSort(nums, leftWall, swapPointer - 1)
  }

  quickSort(nums, swapPointer + 1, pivotIndex)
}

const nums = [3, 2, 5, 0, 1, 8, 7, 6, 9, 4]
quickSort(nums)
console.log(nums)