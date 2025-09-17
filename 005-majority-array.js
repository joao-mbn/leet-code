// https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150

/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function (nums) {
  const counter = new Map();
  let highestCount = 0;
  let highestCountNumber;

  for (let i = 0; i < nums.length; i++) {
    const num = nums[i];
    const currCounter = (counter.get(num) ?? 0) + 1;
    counter.set(num, currCounter);

    if (currCounter > highestCount) {
      highestCount = currCounter;
      highestCountNumber = num;
    }
  }

  return highestCountNumber;
};

/**
 * @param {number[]} nums
 * @return {number}
var majorityElement = function (nums) {
  // Boyer-Moore algorithm
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
 */