// Package main is the entry point for the program.
package main

import (
	"fmt"
)

func main() {
	var (
		input  = []int{2, 3, 1, 1, 4}
		result = jump(input)
	)
	fmt.Println(result)
}
