package queue

type ArrayQueue struct {
	arr []int
}

func NewArrayQueue(cap int) *ArrayQueue {
	return &ArrayQueue{
		arr: make([]int, 0, cap),
	}
}

func (s *ArrayQueue) Enqueue(x int) {
	s.arr = append(s.arr, x)
}

func (s *ArrayQueue) Dequeue() int {
	x := s.arr[0]
	s.arr = s.arr[1:]
	return x
}

func (s *ArrayQueue) Top() int {
	return s.arr[0]
}

func (s *ArrayQueue) Empty() bool {
	return len(s.arr) == 0
}

func (s *ArrayQueue) Array() []int {
	return s.arr
}
