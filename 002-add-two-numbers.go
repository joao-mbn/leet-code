// https://leetcode.com/problems/add-two-numbers/?envType=study-plan-v2&envId=top-interview-150

package main

// ListNode is the Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultSlice := []int{}

	trailing := 0
	for {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + trailing
		trailing = sum / 10
		resultSlice = append(resultSlice, sum%10)

		if l1 == nil && l2 == nil {
			if trailing == 1 {
				resultSlice = append(resultSlice, 1)
			}
			break
		}
	}

	var resultNodes *ListNode
	for i := len(resultSlice) - 1; i >= 0; i-- {
		item := resultSlice[i]
		resultNodes = &ListNode{Val: item, Next: resultNodes}
	}
	return resultNodes
}
