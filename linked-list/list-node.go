package linkedlist

// ListNode is the Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func ToSlice(listNode *ListNode) []int {
	result := []int{}
	for listNode != nil {
		result = append(result, listNode.Val)
		listNode = listNode.Next
	}

	return result
}

func FromSlice(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}

	next := &ListNode{Val: slice[0]}
	listNode := next
	for _, value := range slice[1:] {
		newNode := &ListNode{Val: value}
		next.Next = newNode
		next = next.Next
	}

	return listNode
}
