/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function(head) {
  if (!head?.next) {
    return head
  }

  const nodeStack = []
  while(head) {
    nodeStack.push(head)
    head = head.next
  }

  const firstNode = nodeStack.pop()
  let tail = firstNode
  while(nodeStack.length) {
    tail.next = nodeStack.pop()
    tail = tail.next
  }

  tail.next = null

  return firstNode
};

function ListNode(val, next) {
  this.val = (val===undefined ? 0 : val)
  this.next = (next===undefined ? null : next)
}

reverseList(
  new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5)))))
)
