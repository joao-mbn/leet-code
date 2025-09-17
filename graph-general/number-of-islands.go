// https://leetcode.com/problems/number-of-islands/?envType=study-plan-v2&envId=top-interview-150

package graphgeneral

func NumIslands(grid [][]byte) int {
	return numIslands(grid)
}

func numIslands(grid [][]byte) int {
	count := 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == byte('1') {
				expand(i, j, grid)
				count++
			}
		}
	}

	return count
}

func expand(i, j int, grid [][]byte) [][]byte {
	grid[i][j] = 'x'

	neighbors := [][2]int{
		{i - 1, j}, // north
		{i, j + 1}, // east
		{i, j - 1}, // west
		{i + 1, j}, // south
	}

	for _, neighbor := range neighbors {
		i := neighbor[0]
		if i < 0 || i >= len(grid) {
			continue
		}

		j := neighbor[1]
		if j < 0 || j >= len(grid[0]) {
			continue
		}

		neighborValue := grid[i][j]
		if neighborValue == byte('1') {
			expand(i, j, grid)
		}
	}

	return grid
}
