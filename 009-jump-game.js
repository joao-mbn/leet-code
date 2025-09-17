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

/**
 * @param {number[]} nums
 * @return {boolean}
var canJump = function (nums) {
  const zeroesIndexes = [];

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] === 0) {
      zeroesIndexes.push(i);
    }
  }

  if (zeroesIndexes.length === 0) return true;

  const failedJumpRecord = new Set();

  function canJumpFromPosition(i) {
    let jump = nums[i];

    if (i + jump >= nums.length - 1) {
      return true;
    }

    let win = false;

    while (!win && jump > 0) {
      if (failedJumpRecord.has(i + jump)) {
        break;
      }

      win = canJumpFromPosition(i + jump);
      jump--;
    }

    if (!win) {
      failedJumpRecord.add(i);
    }

    return win;
  }

  return canJumpFromPosition(0);
};
*/