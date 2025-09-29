package matrix

func Rotate(matrix [][]int) {
	rotate(matrix)
}

/*
This solution is naive, albeit performant.
A solution is top reverse order of rows then transposing the matrix
*/
func rotate(matrix [][]int) {
	for i, row := range matrix {
		n := len(row) - 2*i
		if n <= 1 {
			break
		}

		for j := range row[i : len(matrix)-1-i] {
			nextValue := matrix[i][i+j]
			nextI, nextJ := 0, j

			for range 4 {
				nextI, nextJ = walkClockWise(nextI, nextJ, n)
				matrix[nextI+i][nextJ+i], nextValue = nextValue, matrix[nextI+i][nextJ+i]
			}
		}
	}
}

func walkClockWise(i, j, n int) (int, int) {
	stepsLeft := n - 1

	// find out how much could you possibly walk on a certain direction, given the current coordinates
	right := n - 1 - j
	down := n - 1 - i
	left := j
	up := i

	// try walk as much as possible to the left, then down, then right, then up
	newI := i
	newJ := j

	if right > 0 && up == 0 {
		stepsOnDirection := min(stepsLeft, right)
		stepsLeft -= stepsOnDirection
		newJ += stepsOnDirection
		newI += stepsLeft
	} else if down > 0 && right == 0 {
		stepsOnDirection := min(stepsLeft, down)
		stepsLeft -= stepsOnDirection
		newI += stepsOnDirection
		newJ -= stepsLeft
	} else if left > 0 && down == 0 {
		stepsOnDirection := min(stepsLeft, left)
		stepsLeft -= stepsOnDirection
		newJ -= stepsOnDirection
		newI -= stepsLeft
	} else {
		stepsOnDirection := min(stepsLeft, up)
		stepsLeft -= stepsOnDirection
		newI -= stepsOnDirection
		newJ += stepsLeft
	}

	return newI, newJ
}
