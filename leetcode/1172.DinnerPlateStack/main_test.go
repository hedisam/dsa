package platestack

import "testing"

func Test(t *testing.T) {
	d := Constructor(1)
	d.Push(1)
	d.Push(2)
	d.Push(3)
	d.PopAtStack(1)
	d.Pop()
	d.Pop()
	d.Pop()
}