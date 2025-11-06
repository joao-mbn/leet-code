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
 * @return {number[][]}
 */
var levelOrder = function(root) {
  const output = []

  /**
   * @param {TreeNode} branch
   * @param {number} level
   */
  function traverse(branch, level) {
    if (!branch) {
      return
    }

    output[level] ??= []
    output[level].push(branch.val)

    traverse(branch.left, level + 1)
    traverse(branch.right, level + 1)
  }

  traverse(root, 0)

  return output
};

function TreeNode(val, left, right) {
  this.val = (val===undefined ? 0 : val)
  this.left = (left===undefined ? null : left)
  this.right = (right===undefined ? null : right)
}