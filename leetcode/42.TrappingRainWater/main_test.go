package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestTrap(t *testing.T) {
	h := []int{0,1,2,0,3,0,1,2,0,0,4,2,1,2,5,0,1,2,0,2}
	got := Trap(h)

	want := 26

	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}

func getTests() ([][]int, []int) {
	s := rand.NewSource(time.Now().Unix())
	rnd := rand.New(s)

	tests := make([][]int, 100)
	results := make([]int, 100)

	for i, _ := range tests {
		tests[i] = make([]int, rnd.Int31n(100))
		for j := 0; j < len(tests[i]); j++ {
			tests[i][j] = rnd.Intn(100)
		}
		results[i] = Trap(tests[i])
	}

	return tests, results
}