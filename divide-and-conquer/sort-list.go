package divideandconquer

import linkedlist "leet-code/linked-list"

type ListNode = linkedlist.ListNode

func SortList(head *ListNode) *ListNode {
	return sortList(head)
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast := head
	slow := head
	prev := head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	firstHalfSorted := sortList(head)
	secondHalfSorted := sortList(slow)

	return merge(firstHalfSorted, secondHalfSorted)
}

func merge(l1, l2 *ListNode) *ListNode {
	var next *ListNode
	if l1.Val < l2.Val {
		next = l1
		l1 = l1.Next
	} else {
		next = l2
		l2 = l2.Next
	}
	mergedList := next

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			next.Next = l1
			l1 = l1.Next
		} else {
			next.Next = l2
			l2 = l2.Next
		}

		next = next.Next
	}

	if l1 != nil {
		next.Next = l1
	}

	if l2 != nil {
		next.Next = l2
	}

	return mergedList
}
