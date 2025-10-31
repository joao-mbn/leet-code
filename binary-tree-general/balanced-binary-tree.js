/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isBalanced = function(root) {
  function findBalance(root) {
    if (!root) {
      return 0
    }

    const leftDepth = findBalance(root.left)
    const rightDepth = findBalance(root.right)

    if (leftDepth === false || rightDepth === false) {
      return false
    }

    if (Math.abs(leftDepth - rightDepth) > 1) {
      return false
    }

    return 1 + Math.max(leftDepth, rightDepth)
  }

  return findBalance(root) !== false
};



function TreeNode(val, left, right) {
  this.val = (val===undefined ? 0 : val)
  this.left = (left===undefined ? null : left)
  this.right = (right===undefined ? null : right)
}