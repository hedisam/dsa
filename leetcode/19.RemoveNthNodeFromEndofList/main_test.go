package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test(t *testing.T) {
	tt := []struct{
		arr []int
		n int 
		newArr []int 
	} {
		{
			arr: []int{1,2,3,4,5},
			n: 2,
			newArr: []int{1,2,3,5},
		},
		{
			arr: []int{1,2,3,4,5},
			n: 1,
			newArr: []int{1,2,3,4},
		},
		{
			arr: []int{1,2,3,4,5},
			n: 4,
			newArr: []int{1,3,4,5},
		},
		{
			arr: []int{1,2,3,4,5},
			n: 5,
			newArr: []int{2,3,4,5},
		},
	}

	for _, item := range tt {
		t.Run("nth item", func(t *testing.T) {

			head := list2LinkedList(item.arr)

			newHead := removeNthFromEnd(head, item.n)

			afterArr := ll2list(newHead)

			if !assert.Equal(t, item.newArr, afterArr) {
				t.Errorf("n: %d, arr: %v\n", item.n, item.arr)
			}
		})
	}
}