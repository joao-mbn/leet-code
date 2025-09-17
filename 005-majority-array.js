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
