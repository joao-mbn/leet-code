/**
 * @param {number[]} nums
 * @return {boolean}
 */
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
