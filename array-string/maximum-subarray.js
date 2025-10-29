/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
  if (nums.length === 1) {
    return nums[0]
  }

  let currentSum = nums[0]
  let maxSum = currentSum

  for (let i = 1; i < nums.length; i++) {
    const right = nums[i]
    currentSum = Math.max(currentSum, 0) + right
    maxSum = Math.max(maxSum, currentSum)
  }

  return maxSum
};

