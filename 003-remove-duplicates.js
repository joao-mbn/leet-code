// https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150

/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
  let uniqueCount = 0;

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] !== nums[i - 1]) {
      nums[uniqueCount] = nums[i];
      uniqueCount++;
    }
  }

  return uniqueCount;
};

const nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
removeDuplicates(nums);
console.log(nums);

/**
 * @param {number[]} nums
 * @return {number}
var removeDuplicates = function (nums) {
  let uniqueCount = nums.length;
  let lastUnique;

  for (let i = 0; i < uniqueCount; i++) {
    const num = nums[i];

    if (num === lastUnique) {
      for (let j = i; j < uniqueCount - 1; j++) {
        const current = nums[j];
        const next = nums[j + 1];

        nums[j] = next;
        nums[j + 1] = current;
      }

      i--;
      uniqueCount--;
    }

    lastUnique = num;
  }

  return uniqueCount;
};

const nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
removeDuplicates(nums);
console.log(nums);
 */