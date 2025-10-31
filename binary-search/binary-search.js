var search = function(nums, target) {
  return binarySearch(0, nums, target)
};

/**
 * @param {number} globalIndex
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
function binarySearch(globalIndex, nums, target) {
  if (!nums.length) {
    return -1
  }

  if (nums[0] === globalIndex) {
    return globalIndex
  }

  const halfway = Math.floor(nums.length / 2)
  const halfwayValue = nums[halfway]

  if (halfwayValue === target) {
    return globalIndex + halfway
  } else if (halfwayValue < target) {
    return binarySearch(globalIndex + halfway + 1, nums.slice(halfway+1), target)
  }
  return binarySearch(globalIndex, nums.slice(0, halfway), target)
}

search([1, 2, 3], 3)