// Package main is the entry point for the program'.'
package main

import (
	"fmt"
	divideandconquer "leet-code/divide-and-conquer"
	linkedlist "leet-code/linked-list"
)

func main() {

	for _, input := range []struct {
		slice []int
	}{
		{slice: []int{1, 2, 3, 4, 5}},
		{slice: []int{4, 1, 2, 3}},
		{slice: []int{-1, 5, 3, 4, 0}},
		{slice: []int{}},
		{slice: []int{-1}},
	} {
		result := divideandconquer.SortList(linkedlist.FromSlice(input.slice))
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input, linkedlist.ToSlice(result))
	}
}
