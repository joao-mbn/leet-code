package matrix

func SetZeroes(matrix [][]int) {
	setZeroes(matrix)
}

func setZeroes(matrix [][]int) {
	replacement := int(^uint64(0) >> 1)

	for rowIndex, row := range matrix {
		for colIndex, cell := range row {
			if cell != 0 {
				continue
			}

			for i := range len(matrix) {
				if matrix[i][colIndex] != 0 {
					matrix[i][colIndex] = replacement
				}
			}

			for j := range len(row) {
				if matrix[rowIndex][j] != 0 {
					matrix[rowIndex][j] = replacement
				}
			}
		}
	}

	for i, row := range matrix {
		for j, cell := range row {
			if cell == replacement {
				matrix[i][j] = 0
			}
		}
	}
}
