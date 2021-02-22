package heap

// SortNaively uses a max-heap to sort the input array. it builds the heap naively (iteratively).
func SortNaively(array []int) {
	h := NewMaxHeap(len(array))
	for _, key := range array {
		h.InsertKey(key)
	}

	for i := 0; i < len(array); i++ {
		key, _ := h.ExtractMax()
		array[i] = key
	}
}

// Sort uses a max-heap to sort the input array. It would heapify the input array in-place without using extra space.
func Sort(array []int) {
	heapify(array)

	for i := len(array) - 1; i > 0; i-- {
		swap(array, 0, i)
		siftDown(array, i, 0)
	}
	//size := len(array)
	//for i := 0; i < len(array) - 1; i++ {
	//	swap(array, 0, size - 1)
	//	size--
	//	siftDown(array, size,0)
	//}
}

func heapify(array []int) {
	size := len(array)
	for i := size/2 - 1; i >= 0; i-- {
		siftDown(array, size, i)
	}
}

func siftDown(array []int, size, i int) {
	l := left(i)
	r := right(i)
	largest := i

	if l < size && array[l] > array[i] {
		largest = l
	}
	if r < size && array[r] > array[largest] {
		largest = r
	}

	if i != largest {
		swap(array, i, largest)
		siftDown(array, size, largest)
	}
}