/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function (nums1, m, nums2, n) {
  if (n === 0) return;

  let pointer = 0;
  let size = m;

  for (let i = 0; i < n; i++) {
    const curr2 = nums2[i];

    while (curr2 > nums1[pointer] && pointer < size) {
      pointer++;
    }

    nums1.splice(pointer, 0, curr2);
    nums1.pop();
    size++;
  }
};

const nums1 = [0];
const m = nums1.filter((n) => n > 0).length;
const nums2 = [1];
const n = nums2.length;

merge(nums1, m, nums2, n);
console.log(nums1);
