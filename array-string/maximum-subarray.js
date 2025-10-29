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

console.log(maxSubArray([-2,3,1,3]))
console.log(maxSubArray([-2,1,-3,4,-1,2,1,-5,4]))
console.log(maxSubArray([1]))
console.log(maxSubArray([5,4,-1,7,8]))
console.log(maxSubArray([-1,-2]))
console.log(maxSubArray([5000,1,9999,1999]))
console.log(maxSubArray([1,4,-9999,5,6]))