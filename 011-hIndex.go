package main

import (
	"slices"
)

func hIndex(citations []int) int {
	hIndex := 0

	slices.Sort(citations)

	size := len(citations)

	for i := 0; i < size; i++ {
		citation := citations[i]

		if citation < size - i {
			hIndex = citation
		} else if citation >= size - i {
			return size - i
		}
	}

	return hIndex
}