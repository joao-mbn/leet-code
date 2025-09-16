package heap

type Heap struct {
	slice []int
}

func Heapify(slice []int) *Heap {
	heap := &Heap{
		slice: []int{},
	}

	for _, val := range slice {
		heap.Insert(val)
	}

	return heap
}

func (h *Heap) Insert(value int) {
	h.slice = append(h.slice, value)
	if len(h.slice) == 1 {
		return
	}

	childIndex := len(h.slice) - 1
	for {
		parentIndex := h.parentIndex(childIndex)

		if h.slice[parentIndex] <= h.slice[childIndex] {
			break
		}

		h.slice[parentIndex], h.slice[childIndex] = h.slice[childIndex], h.slice[parentIndex]

		if parentIndex <= 0 {
			break
		}
		childIndex = parentIndex
	}
}

func (h *Heap) Pop() {
	size := len(h.slice)
	switch size {
	case 1:
		h.slice = []int{}
		return
	case 2:
		h.slice = []int{h.slice[1]}
		return
	}

	h.slice = append([]int{h.slice[size-1]}, h.slice[1:size-1]...)

	parentIndex := 0
	for {
		left, right := h.childrenIndexes(parentIndex)

		if left >= size-1 {
			break
		}
		childIndex := left

		if right < size-1 && h.slice[left] > h.slice[right] {
			childIndex = right
		}

		if h.slice[parentIndex] <= h.slice[childIndex] {
			break
		}

		h.slice[parentIndex], h.slice[childIndex] = h.slice[childIndex], h.slice[parentIndex]

		if childIndex == size-1 {
			break
		}
		parentIndex = childIndex
	}
}

func (h *Heap) Replace(value int) {
	size := len(h.slice)
	h.slice[0] = value
	if len(h.slice) == 1 {
		return
	}

	parentIndex := 0
	for {
		left, right := h.childrenIndexes(parentIndex)

		if left > size-1 {
			break
		}
		childIndex := left

		if right <= size-1 && h.slice[left] > h.slice[right] {
			childIndex = right
		}

		if h.slice[parentIndex] <= h.slice[childIndex] {
			break
		}

		h.slice[parentIndex], h.slice[childIndex] = h.slice[childIndex], h.slice[parentIndex]
		parentIndex = childIndex
	}
}

func (h *Heap) First() int {
	return h.slice[0]
}

func (h *Heap) childrenIndexes(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}

func (h *Heap) parentIndex(i int) int {
	return (i - 1) / 2
}
