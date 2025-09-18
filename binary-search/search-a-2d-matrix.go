// https://leetcode.com/problems/search-a-2d-matrix/description/?envType=study-plan-v2&envId=top-interview-150

package binarysearch

func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}

// n = 2 | midrow = 0 | slices = :1, 1:
// n = 3 | midrow = 0 | slices = :1, 1:
// n = 4 | midrow = 1 | slices = :2, 2:
// n = 5 | midrow = 1 | slices = :2, 2:
// n = 6 | midrow = 2 | slices = :3, 3:

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 1 {
		return searchSlice(matrix[0], target)
	}

	midRowIndex := len(matrix)/2 - 1
	midRow := matrix[midRowIndex]

	if target > midRow[len(midRow)-1] {
		return searchMatrix(matrix[midRowIndex+1:], target)
	}
	return searchMatrix(matrix[:midRowIndex+1], target)
}

func searchSlice(slice []int, target int) bool {
	if len(slice) == 1 {
		return slice[0] == target
	}

	midIndex := len(slice)/2 - 1

	if target > slice[midIndex] {
		return searchSlice(slice[midIndex+1:], target)
	}
	return searchSlice(slice[:midIndex+1], target)
}
