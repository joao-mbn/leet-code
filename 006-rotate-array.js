/**
 * @param {number[]} nums
 * @param {number} k
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var rotate = function (nums, k) {
  if (k === 0) return;

  const size = nums.length;

  if (size === 1) return;
  if (k >= size && k % size === 0) return;

  let rotationCount = 1;
  let index = 0;
  let replacement = nums[index];

  while (rotationCount <= size) {
    let newIndex = index + k;
    if (newIndex > size - 1) {
      newIndex %= size;
    }

    const willLoop = k > 1 && rotationCount > 1 && (size % (k * rotationCount) === 0 || (k * rotationCount) % size === 0);

    if (willLoop) {
      nums[newIndex] = replacement;
      index = newIndex + 1;
      replacement = nums[index];
    } else {
      const newReplacement = nums[newIndex];
      nums[newIndex] = replacement;
      index = newIndex;
      replacement = newReplacement;
    }

    rotationCount++;
  }
};
