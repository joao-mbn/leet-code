// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150

/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function (nums, val) {
  let eqValCount = 0;

  for (let i = 0; i < nums.length - eqValCount; i++) {
    const item = nums[i];

    if (val === item) {
      const swapPointer = nums.length - eqValCount - 1;
      const swapItem = nums[swapPointer];

      nums[i] = swapItem;

      i--;
      eqValCount++;
    }
  }

  return nums.length - eqValCount;
};
