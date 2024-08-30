/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function (nums1, m, nums2, n) {
  if (n === 0) return;

  if (m === 0) {
    for (let i = 0; i < n; i++) {
      nums1[i] = nums2[i];
    }
    return;
  }

  let p1 = m - 1;
  let insertionP = m + n - 1;

  for (let p2 = n - 1; p2 >= 0; p2--) {
    let curr1 = nums1[p1];
    const curr2 = nums2[p2];

    while (curr2 < curr1) {
      nums1[insertionP] = curr1;
      insertionP--;

      p1--;
      curr1 = nums1[p1];
    }

    nums1[insertionP] = curr2;
    insertionP--;
  }
};

const nums1 = [1, 2, 3, 0, 0, 0];
const m = nums1.filter((n) => n > 0).length;
const nums2 = [2, 5, 6];
const n = nums2.length;

merge(nums1, m, nums2, n);
console.log(nums1);
