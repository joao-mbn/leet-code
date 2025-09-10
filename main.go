// Package main is the entry point for the program'.'
package main

import "fmt"

func main() {
	for _, input := range []struct {
		grid [][]byte
	}{
		{grid: [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}},
		{grid: [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
		}},
		{grid: [][]byte{
			{'X', 'X'},
			{'X', 'X'},
		}},
		{grid: [][]byte{
			{'O', 'O'},
			{'O', 'O'},
		}},
		{grid: [][]byte{
			{'O'},
		}},
		{grid: [][]byte{
			{'X', 'X', 'X'},
			{'X', 'O', 'X'},
			{'X', 'X', 'X'},
		}},
	} {
		solve(input.grid)
		fmt.Printf("Given the inputs: %v, the result is: %v\n", input)
	}
}
