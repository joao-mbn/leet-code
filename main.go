// Package main is the entry point for the program'.'
package main

import "fmt"

func main() {
	for _, input := range []struct {
		l1 *ListNode
		l2 *ListNode
	}{
		{
			l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
		},
		{
			l1: &ListNode{Val: 0},
			l2: &ListNode{Val: 0},
		},
		{
			l1: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}},
			l2: &ListNode{Val: 9, Next: &ListNode{Val: 9}},
		},
	} {
		result := addTwoNumbers(input.l1, input.l2)
		fmt.Printf("Given the inputs: %v and %v, the result is: %v\n", input.l1, input.l2, result)
	}
}
