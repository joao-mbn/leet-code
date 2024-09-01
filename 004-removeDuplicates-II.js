/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
  let atMostDuplicateCount = Math.min(nums.length, 2);

  for (let i = 2; i < nums.length; i++) {
    if (nums[i] !== nums[atMostDuplicateCount - 2]) {
      atMostDuplicateCount++;
    }

    nums[atMostDuplicateCount - 1] = nums[i];
  }

  return atMostDuplicateCount;
};
