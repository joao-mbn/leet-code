// Package main is the entry point for the program'.'
package main

import "fmt"

func main() {
	for _, input := range []struct {
		preorder []int
		ineorder []int
	}{
		{
			preorder: []int{1, 2},
			ineorder: []int{1, 2},
		},
		{
			preorder: []int{1, 2, 4, 5, 9, 3, 6, 7, 15},
			ineorder: []int{4, 2, 9, 5, 1, 6, 3, 15, 7},
		},
		{
			preorder: []int{3, 9, 20, 15, 7},
			ineorder: []int{9, 3, 15, 20, 7},
		},
		{
			preorder: []int{-1},
			ineorder: []int{-1},
		},
	} {
		result := buildTree(input.preorder, input.ineorder)
		fmt.Printf("Given the inputs: %v and %v, the result is: %v\n", input.preorder, input.ineorder, result)
	}
}
