package stack

import "fmt"

type FloatArrayStack struct {
	arr []float64
	pos int
}

func NewFloatArrayStack(cap int) *FloatArrayStack {
	return &FloatArrayStack{
		arr: make([]float64, cap),
		pos: -1,
	}
}

func (s *FloatArrayStack) Full() bool {
	return s.pos == len(s.arr) - 1
}

func (s *FloatArrayStack) Empty() bool {
	return s.pos < 0
}

func (s *FloatArrayStack) Push(x float64) {
	s.pos++
	s.arr[s.pos] = x
}

func (s *FloatArrayStack) Pop() float64 {
	x := s.arr[s.pos]
	s.pos--
	return x
}

func (s *FloatArrayStack) Top() float64 {
	return s.arr[s.pos]
}

func (s *FloatArrayStack) String() string {
	return fmt.Sprint("Stack:", s.arr)
}

func (s *FloatArrayStack) Array() []float64 {
	arr := make([]float64, s.pos+1)
	j := 0
	for i := s.pos; i >= 0; i-- {
		arr[j] = s.arr[i]
		j++
	}
	return arr
}

