package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test(t *testing.T) {
	tt := []struct{
		Arr []int
		Palindrome bool
	}{
		{
			[]int{1,2,2,1},
			true,
		}, 
		{
			[]int{1,2},
			false,
		},
		{
			[]int{1},
			true,
		},
		{
			[]int{2,3,4,5,4,3,2},
			true,
		},
	}

	for _, item := range tt {
		head := array2LinkedList(item.Arr)
		ok := isPalindrome(head)
		assert.Equal(t, item.Palindrome, ok)
	}
}