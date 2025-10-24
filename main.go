// Package main is the entry point for the program'.'
package main

import (
	"fmt"
	"leet-code/intervals"
)

func main() {

	for _, input := range []struct {
		intervals   [][]int
		newInterval []int
	}{
		{intervals: [][]int{{3, 5}, {12, 15}}, newInterval: []int{6, 6}},
	} {
		result := intervals.Insert(input.intervals, input.newInterval)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input, result)
	}
}
