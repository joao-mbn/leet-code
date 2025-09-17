// https://leetcode.com/problems/valid-sudoku/?envType=study-plan-v2&envId=top-interview-150

package main

func isValidSudoku(board [][]byte) bool {

	// check rows, while building the columns
	for i, row := range board {
		rowSet := map[byte]bool{}
		colSet := map[byte]bool{}
		for j, cell := range row {
			if _, has := rowSet[cell]; has && cell != '.' {
				return false
			}
			rowSet[cell] = true

			transposedCoordCell := board[j][i]
			if _, has := colSet[transposedCoordCell]; has && transposedCoordCell != '.' {
				return false
			}
			colSet[transposedCoordCell] = true
		}
	}

	// check sub-boxes
	for i := range 3 {
		for j := range 3 {
			box := map[byte]bool{}

			for bi := range 3 {
				for bj := range 3 {
					bij := board[i*3+bi][j*3+bj]
					if bij == '.' {
						continue
					}

					if _, has := box[bij]; has {
						return false
					}

					box[bij] = true
				}
			}
		}
	}

	return true
}
