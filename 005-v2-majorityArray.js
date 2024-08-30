/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function (nums) {
  /* Boyer-Moore algorithm */
  let candidate = nums[0];
  let counter = 1;

  for (let i = 1; i < nums.length; i++) {
    const num = nums[i];

    if (num === candidate) {
      counter++;
    } else {
      counter--;
    }

    if (counter === 0) {
      candidate = num;
      counter = 1;
    }
  }

  return candidate;
};
