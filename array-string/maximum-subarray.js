/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
  if (nums.length === 1) {
    return nums[0]
  }

  let leftPointer = 0
  let currentSum = nums[leftPointer]
  let maxSum = currentSum

  for (let i = 1; i < nums.length; i++) {
    while (currentSum <= 0 && leftPointer < i) {
      currentSum -= nums[leftPointer]
      leftPointer++
    }

    const right = nums[i]
    currentSum += right
    maxSum = Math.max(maxSum, currentSum)
  }

  return maxSum
};

