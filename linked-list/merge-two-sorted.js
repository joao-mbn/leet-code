function ListNode(val, next) {
  this.val = (val===undefined ? 0 : val)
  this.next = (next===undefined ? null : next)
}

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
var mergeTwoLists = function(list1, list2) {
  if (!(list1 || list2)) {
    return null
  }

  let head = new ListNode()
  let tail = head

  while (list1 || list2) {
    if (list1 && (!list2 || list1.val < list2.val)) {
      tail.val = list1.val
      list1 = list1.next
    } else {
      tail.val = list2.val
      list2 = list2.next
    }

    if (list1 || list2) {
      tail.next = new ListNode()
      tail = tail.next
    }
  }

  return head
};

mergeTwoLists(
  new ListNode(),
  new ListNode()
)

mergeTwoLists(
  new ListNode(1, new ListNode(2, new ListNode(4))),
  new ListNode(1, new ListNode(3, new ListNode(4)))
)