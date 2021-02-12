package queue

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	s := NewArrayQueue(10)

	for i := 1; i <= 5; i++ {
		s.Enqueue(i)
		fmt.Println("Enqueue:", s.arr)
	}

	for i := 0; i < 5; i++ {
		s.Dequeue()
		fmt.Println("Dequeue:", s.arr)
	}
}
