package main

func canCompleteCircuit(gas []int, cost []int) int {

	size := len(gas)
	deltas := make([]int, size)

	for i, v := range gas {
		deltas[i] = v - cost[i]
	}

	i := 0
	endPointer := getNextPointerInCycle(i, size)
	currentGas := deltas[i]

	if size == 1 && currentGas >= 0 {
		return 0
	}

	for i <= size-1 {

		if currentGas > 0 {

			currentGas += deltas[endPointer]

			if hasCycledSlice(i, endPointer, size) && currentGas >= 0 {
				return i
			}

			endPointer = getNextPointerInCycle(endPointer, size)

		} else {

			currentGas -= deltas[i]
			i++

			if i == endPointer {
				currentGas = deltas[i]
				endPointer = getNextPointerInCycle(endPointer, size)
			}

		}

	}

	return -1

}

func hasCycledSlice(i1, i2, size int) bool {
	if i2 > i1 {
		return i2 == size-1 && i1 == 0
	}

	return (i2 + 1 + size - i1) == size
}

func getNextPointerInCycle(i, size int) int {
	if i == size-1 {
		return 0
	}

	return i + 1
}

// for i := 0; i < size-1; i++ {
// 	currentGas := deltas[i]
// 	endPointer := getNextPointerInCycle(i, size)

// 	for currentGas > 0 {
// 		currentGas += deltas[endPointer]

// 		if hasCycledSlice(i, endPointer, size) {
// 			return i
// 		}

// 		endPointer = getNextPointerInCycle(endPointer, size)
// 	}
// }
