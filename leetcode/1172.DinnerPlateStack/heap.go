package platestack

import (
	"math"
)

type MinHeap struct {
	capacity int
	size int
	array []int
}

// Size returns the count of elements currently saved in the heap
func (h *MinHeap) Size() int {
	return h.size
}

// InsertKey inserts the key value into the heap.
func (h *MinHeap) InsertKey(key int) bool {
	if h.size >= len(h.array) {
		// by appending a random value like 0, we double the size of our array
		// value zero will be saved at index h.size but then it will be replaced by the value of key
		h.array = append(h.array, 0)
	}

	// insert the new key at the end
	i := h.size
	h.array[i] = key
	h.size++

	// fix the min-heap property if it's violated
	h.siftUp(i)

	return true
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 && h.array[h.parent(i)] > h.array[i] {
		// swap key at index i with its parent
		swap(h.array, i, h.parent(i))
		i = h.parent(i)
	}
}

// DeleteKey deletes key at index i.
// Returns false if index is out of range.
func (h *MinHeap) DeleteKey(i int) bool {
	if i >= h.size {
		return false
	}

	// decrease the key value to minus infinity so it goes and sits on the root.
	_ = h.decreaseKey(i, math.MinInt64)
	// now extracting the min key value will make it deleted from the heap.
	_, _ = h.ExtractMin()

	return true
}

// Min returns the min value of the heap which is the root key value.
// Returns (0, false) if the heap is empty.
func (h *MinHeap) Min() (int, bool) {
	if h.size == 0 {
		return 0, false
	}

	return h.array[0], true
}

// decreaseKey decreases key value at index i to newValue.
// newValue must be smaller than current key value.
// Returns false if index i is out of range, or if newValue >= value at index i.
func (h *MinHeap) decreaseKey(i, newValue int) bool {
	if newValue >= h.array[i] {
		return false
	}

	h.array[i] = newValue

	// fix any possible min-heap property violation
	h.siftUp(i)

	return true
}

// ExtractMin removes and returns the min key value of the min-heap, which is its root key.
// it returns (0, false) if the heap is empty.
func (h *MinHeap) ExtractMin() (int, bool) {
	if h.size == 0 {
		// heap is empty
		return 0, false
	} else if h.size == 1 {
		h.size--
		return h.array[0], true
	}

	// the first element or the root value is the min value
	root := h.array[0]
	// replace the root value by the last element and then fix any violation caused by that (heapify)
	h.array[0] = h.array[h.size - 1]
	h.size--
	h.siftDown(0)

	return root, true
}

// right returns index of right element of the key value stored at index i
func (h *MinHeap) right(i int) int {
	return (i * 2) + 2
}

// left returns index of left element of the key value stored at index i
func (h *MinHeap) left(i int) int {
	return (i * 2) + 1
}

// parent returns index of parent element of the key value stored at index i
func (h *MinHeap) parent(i int) int {
	return (i - 1) / 2
}

// swap swaps key value at index i with key value at index j
func swap(array []int, i, j int) {
	jKey := array[j]
	array[j] = array[i]
	array[i] = jKey
}

// siftDown recursively heapify a subtree with with the root at given index.
// It assumes that the subtrees are already heapified.
func (h *MinHeap) siftDown(i int) {
	left := h.left(i)
	right := h.right(i)
	smallest := i

	if left < h.size && h.array[left] < h.array[i] {
		smallest = left
	}
	if right < h.size && h.array[right] < h.array[smallest] {
		smallest = right
	}
	if smallest != i {
		swap(h.array, i, smallest)
		h.siftDown(smallest)
	}
}

func NewMinHeap(capacity int) *MinHeap {
	return &MinHeap{
		capacity: capacity,
		size:     0,
		array:    make([]int, capacity),
	}
}
