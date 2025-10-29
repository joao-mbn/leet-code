/**
 * @param {number[]} nums
 * @return {number}
 */
var maxProduct = function(nums) {
  let currentProduct = 1
  let maxValue = -Infinity

  for (let i = 0; i < nums.length; i++) {
    const num = nums[i]
    currentProduct *= num
    maxValue = Math.max(maxValue, currentProduct)

    if (currentProduct === 0) {
      currentProduct = 1
    }
  }

  currentProduct = 1

  for (let i = nums.length - 1; i >= 0; i--) {
    const num = nums[i]
    currentProduct *= num
    maxValue = Math.max(maxValue, currentProduct)

    if (currentProduct === 0) {
      currentProduct = 1
    }
  }

  return maxValue
};

console.log(maxProduct([2,3,-2,4]))
console.log(maxProduct([-2,0,-1]))
console.log(maxProduct([-2,-1]))
console.log(maxProduct([0,2]))