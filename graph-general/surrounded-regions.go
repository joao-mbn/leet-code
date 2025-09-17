// https://leetcode.com/problems/surrounded-regions/?envType=study-plan-v2&envId=top-interview-150

package graphgeneral

func solve(board [][]byte) {
	visited := map[int]map[int]bool{}
	surroundedCells := [][2]int{}

	for i, row := range board {
		for j := range row {
			region, touchesBorder := chart(i, j, board, visited, [][2]int{}, false)

			if !touchesBorder {
				surroundedCells = append(surroundedCells, region...)
			}
		}
	}

	for _, surroundedCell := range surroundedCells {
		i, j := surroundedCell[0], surroundedCell[1]
		board[i][j] = 'X'
	}
}

func chart(i, j int, board [][]byte, visited map[int]map[int]bool, region [][2]int, touchesBorder bool) ([][2]int, bool) {
	if isOutOfBoard(i, j, board) {
		return region, touchesBorder
	}

	shouldVisit := isUnchartedRegion(i, j, board, visited)

	if _, ok := visited[i]; !ok {
		visited[i] = map[int]bool{}
	}
	visited[i][j] = true
	if !shouldVisit {
		return region, touchesBorder
	}

	touchesBorder = touchesBorder || cellTouchesBorder([2]int{i, j}, [2]int{len(board), len(board[0])})
	region = append(region, [2]int{i, j})

	neighbors := [][2]int{
		{i - 1, j}, // north
		{i, j + 1}, // east
		{i, j - 1}, // west
		{i + 1, j}, // south
	}
	for _, neighbor := range neighbors {
		region, touchesBorder = chart(neighbor[0], neighbor[1], board, visited, region, touchesBorder)
	}

	return region, touchesBorder
}

func cellTouchesBorder(cell [2]int, boardSize [2]int) bool {
	ic := cell[0]
	jc := cell[1]

	if ic == 0 || jc == 0 {
		return true
	}

	if ic == boardSize[0]-1 || jc == boardSize[1]-1 {
		return true
	}

	return false
}

func isOutOfBoard(i, j int, board [][]byte) bool {
	if i < 0 || i >= len(board) {
		return true
	}

	if j < 0 || j >= len(board[0]) {
		return true
	}

	return false
}

func isUnchartedRegion(i, j int, board [][]byte, visited map[int]map[int]bool) bool {
	if board[i][j] != byte('O') {
		return false
	}

	cellsVisited, visitedRow := visited[i]
	visitedCell := false
	if visitedRow {
		visitedCell = cellsVisited[j]
	}

	return !(visitedRow && visitedCell)
}
