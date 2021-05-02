package main

import "fmt"

func main() {
	s := Constructor()
	// push 2
	s.Push(2)
	// push 0
	s.Push(0)
	// push 3
	s.Push(3)
	// push 0
	s.Push(0)
	// get min
	fmt.Println(s.GetMin())
	// pop
	s.Pop()
	// get min
	fmt.Println(s.GetMin())
	// pop
	s.Pop()
	// get min
	fmt.Println(s.GetMin())
	// pop
	s.Pop()
	// get min
	fmt.Println(s.GetMin())
}

type MinStack struct {
	data *stack
	mins *stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: newStack(),
		mins: newStack(),
	}
}

func (this *MinStack) Push(val int) {
	this.data.push(val)

	if this.mins.size() == 0 || val < this.data.arr[this.mins.top()] {
		this.mins.push(this.data.size() - 1)
	}
}

func (this *MinStack) Pop() {
	pos := this.data.pos
	this.data.pop()

	if pos == this.mins.top() {
		this.mins.pop()
	}
}

func (this *MinStack) Top() int {
	return this.data.top()
}

func (this *MinStack) GetMin() int {
	pos := this.mins.top()
	return this.data.arr[pos]
}

type stack struct {
	arr []int
	pos int
}

func newStack() *stack {
	return &stack{pos: -1}
}

func (s *stack) push(val int) {
	s.pos++
	if s.pos >= len(s.arr) {
		s.arr = append(s.arr, val)
	} else {
		s.arr[s.pos] = val
	}
}

func (s *stack) pop() {
	s.pos--
}

func (s *stack) top() int {
	return s.arr[s.pos]
}

func (s *stack) size() int {
	return s.pos + 1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */