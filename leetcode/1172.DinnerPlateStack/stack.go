package platestack

type Stack struct {
	arr []int
	pos int  
}

func NewStack(capacity int) *Stack {
	return &Stack{arr: make([]int, capacity), pos: -1}
}

// Push pushes val into the stack. it will panic if the stack is already full.
func (s *Stack) Push(val int) {
	s.pos++
	s.arr[s.pos] = val
}

// Pop returns and deletes the top element in the stack. it will panic if the stack is empty.
func (s *Stack) Pop() int {
	val := s.arr[s.pos]
	s.pos-- 
	return val
}

func (s *Stack) IsEmpty() bool {
	return s.pos < 0
}

func (s *Stack) IsFull() bool {
	return s.pos + 1 >= len(s.arr)
}

func (s *Stack) Size() int {
	return s.pos + 1
}