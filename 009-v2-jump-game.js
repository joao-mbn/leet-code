/**
 * @param {number[]} nums
 * @return {boolean}
 */
var canJump = function (nums) {
  if (nums.length === 1) return true;
  if (nums[0] === 0) return false;

  let maxJump = nums[0];

  for (let i = 1; i < nums.length; i++) {
    if (i > maxJump) {
      return false; // exceeded the jump capacity without reaching the last index without any chance to update to a greater position
    }

    if (i + nums[i] > maxJump) {
      maxJump = i + nums[i]; // encompasses all jump possibilities of the previous index plus more
    }

    if (maxJump >= nums.length - 1) {
      return true;
    }
  }

  return false;
};
