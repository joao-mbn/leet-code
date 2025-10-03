package linkedlist

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEnd(head, n)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead, _ := removeOnePass(head, 0, n)

	return newHead
}

func removeOnePass(head *ListNode, prevSize, n int) (*ListNode, int) {
	currentSize := prevSize + 1

	var totalSize int
	if head.Next == nil {
		totalSize = currentSize
	} else {
		next, newTotalSize := removeOnePass(head.Next, currentSize, n)
		head.Next = next
		totalSize = newTotalSize
	}

	if totalSize-currentSize == n-1 {
		head = head.Next
	}

	return head, totalSize
}
