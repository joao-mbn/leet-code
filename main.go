// Package main is the entry point for the program'.'
package main

import (
	"fmt"
	linkedlist "leet-code/linked-list"
)

func main() {

	for _, input := range []struct {
		head []int
		n    int
	}{
		{head: []int{1}, n: 1},
		{head: []int{1, 2}, n: 1},
		{head: []int{1, 2, 3}, n: 3},
		{head: []int{1, 2, 3, 4, 5}, n: 2},
		{head: []int{1, 2, 3, 4, 5}, n: 4},
	} {
		result := linkedlist.RemoveNthFromEnd(linkedlist.FromSlice(input.head), input.n)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input, linkedlist.ToSlice(result))
	}
}
