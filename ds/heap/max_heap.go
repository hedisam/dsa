package heap

import (
	"fmt"
	"math"
	"strings"
)

type iMaxHeap interface {
	ExtractMax() (int, bool)
	InsertKey(key int) bool
	Max() (int, bool)
	Size() int
	DeleteKey(i int) bool
	increaseKey(i, newValue int) bool
	right(i int) int
	left(i int) int
	parent(i int) int
	siftUp(i int)
	siftDown(i int)
}

type MaxHeap struct {
	array []int
	capacity int
	size int
}

func (h *MaxHeap) ExtractMax() (int, bool) {
	if h.size == 0 {
		return 0, false
	} else if h.size == 1 {
		h.size--
		return h.array[0], true
	}

	root := h.array[0]
	h.array[0] = h.array[h.size - 1]
	h.size--

	h.siftDown(0)

	return root, true
}

func (h *MaxHeap) InsertKey(key int) bool {
	if h.size == h.capacity {
		return false
	}

	i := h.size
	h.array[i] = key
	h.size++

	h.siftUp(i)

	return true
}

func (h *MaxHeap) Max() (int, bool) {
	return h.array[0], h.size != 0
}

func (h *MaxHeap) Size() int {
	return h.size
}

func (h *MaxHeap) DeleteKey(i int) bool {
	if i >= h.size {
		return false
	}

	h.increaseKey(i, math.MaxInt64)

	h.ExtractMax()

	return true
}

// newValue must be greater than current value
func (h *MaxHeap) increaseKey(i, newValue int) bool {
	if i >= h.size || newValue <= h.array[i] {
		return false
	}

	h.array[i] = newValue
	h.siftUp(i)

	return true
}

func (h *MaxHeap) right(i int) int {
	return 2 * i + 2
}

func (h *MaxHeap) left(i int) int {
	return 2 * i + 1
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) siftUp(i int) {
	for i > 0 && h.array[i] > h.array[h.parent(i)] {
		swap(h.array, i, h.parent(i))
		i = h.parent(i)
	}
}

func (h *MaxHeap) siftDown(i int) {
	left := h.left(i)
	right := h.right(i)
	largest := i

	if left < h.size && h.array[left] > h.array[i] {
		largest = left
	}
	if right < h.size && h.array[right] > h.array[largest] {
		largest = right
	}

	if i != largest {
		swap(h.array, i, largest)
		h.siftDown(largest)
	}
}

func (h *MaxHeap) String() string {
	if h.size == 0 {
		return "Empty Max-Heap"
	}

	var sb strings.Builder
	sb.WriteString("[")

	for i, e := range h.array {
		_, _ = fmt.Fprint(&sb, e)
		if i == h.size {
			break
		}
		sb.WriteString(", ")
	}
	sb.WriteString("]")
	return sb.String()
}

func NewMaxHeap(capacity int) *MaxHeap {
	return &MaxHeap{
		array:    make([]int, capacity),
		capacity: capacity,
		size:     0,
	}
}

