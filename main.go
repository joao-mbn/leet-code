// Package main is the entry point for the program'.'
package main

import (
	"fmt"
	heap "leet-code/Heap"
)

func main() {

	for _, input := range []struct {
		nums []int
		k    int
	}{
		{nums: []int{-1, 2, 0}, k: 1},
		{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4},
		{nums: []int{3, 3, 3, 3}, k: 4},
		{nums: []int{3, 3, 3, 3, 4}, k: 4},
		{nums: []int{3, 3, 3, 3, 2}, k: 4},
		{nums: []int{3}, k: 1},
		{nums: []int{3, 3, 3, 1, 2}, k: 4},
		{nums: []int{3, 3, 1, 2}, k: 3},
		{nums: []int{3, 2, 1, 5, 6, 4}, k: 2},
	} {
		result := heap.FindKthLargest(input.nums, input.k)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input, result)
	}
}
