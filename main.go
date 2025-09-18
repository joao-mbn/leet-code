// Package main is the entry point for the program'.'
package main

import (
	"fmt"
	binarysearch "leet-code/binary-search"
)

func main() {

	for _, input := range []struct {
		matrix [][]int
		target int
	}{
		{target: 3, matrix: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}},
		{target: 15, matrix: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}},
		{target: 61, matrix: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}},
		{target: 0, matrix: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}},
		{target: 5, matrix: [][]int{
			{1, 3, 5, 7},
		}},
		{target: 5, matrix: [][]int{
			{1},
			{3},
			{5},
			{7},
		}},
		{target: 5, matrix: [][]int{
			{1},
		}},
	} {
		result := binarysearch.SearchMatrix(input.matrix, input.target)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input, result)
	}
}
