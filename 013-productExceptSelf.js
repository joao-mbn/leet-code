/**
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function (nums) {
  const answer = [];

  let forwardPointer = 1;
  for (let i = 0; i < nums.length; i++) {
    answer[i] = forwardPointer;
    forwardPointer *= nums[i];
  }

  let backwardsPointer = 1;
  for (let j = nums.length - 1; j >= 0; j--) {
    answer[j] *= backwardsPointer;
    backwardsPointer *= nums[j];
  }

  return answer;
};
