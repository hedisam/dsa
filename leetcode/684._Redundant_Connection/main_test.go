package main

import (
	"reflect"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	edges := [][]int{{1,2}, {2,3}, {3,4}, {1,4}, {1,5}}
	want := []int{1,4}

	result := FindRedundantConnection(edges)

	if !reflect.DeepEqual(want, result) {
		t.Errorf("wrong result, got: %v, want: %v\n", result, want)
	}
}
