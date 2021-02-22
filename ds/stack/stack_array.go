package stack

import "fmt"

type ArrayStack struct {
	arr []int
	pos int
}

func NewArrayStack(cap int) *ArrayStack {
	return &ArrayStack{
		arr: make([]int, cap),
		pos: -1,
	}
}

func (s *ArrayStack) Full() bool {
	return s.pos == len(s.arr) - 1
}

func (s *ArrayStack) Empty() bool {
	return s.pos < 0
}

func (s *ArrayStack) Push(x int) {
	s.pos++
	s.arr[s.pos] = x
}

func (s *ArrayStack) Pop() int {
	x := s.arr[s.pos]
	s.pos--
	return x
}

func (s *ArrayStack) Top() int {
	return s.arr[s.pos]
}

func (s *ArrayStack) String() string {
	return fmt.Sprint("TopSort:", s.arr)
}

func (s *ArrayStack) Array() []int {
	arr := make([]int, s.pos+1)
	j := 0
	for i := s.pos; i >= 0; i-- {
		arr[j] = s.arr[i]
		j++
	}
	return arr
}
