// Package main is the entry point for the program'.'
package main

import "fmt"

func main() {
	for _, input := range []struct {
		intervals [][]int
	}{
		{intervals: [][]int{{8, 10}, {1, 3}, {2, 6}, {15, 18}}},
		{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
		{intervals: [][]int{{15, 18}, {1, 4}, {4, 5}}},
		{intervals: [][]int{{1, 4}, {4, 5}}},
		{intervals: [][]int{{1, 4}}},
		{intervals: [][]int{{4, 4}}},
		{intervals: [][]int{{0, 0}}},
		{intervals: [][]int{{1, 1}, {1, 1}}},
		{intervals: [][]int{{1, 1}, {1, 4}}},
		{intervals: [][]int{{1, 1}, {2, 4}}},
	} {
		result := merge(input.intervals)
		fmt.Printf("Given the input: %v, the result is: %v\n", input.intervals, result)
	}
}
