package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test(t *testing.T) {
	graph := strToArray(t, "[[3,1],[4,6,7,2,5],[4,6,3],[6,4],[7,6,5],[6],[7],[]]")
	expected := strToArray(t, "[[0,3,6,7],[0,3,4,7],[0,3,4,6,7],[0,3,4,5,6,7],[0,1,4,7],[0,1,4,6,7],[0,1,4,5,6,7],[0,1,6,7],[0,1,7],[0,1,2,4,7],[0,1,2,4,6,7],[0,1,2,4,5,6,7],[0,1,2,6,7],[0,1,2,3,6,7],[0,1,2,3,4,7],[0,1,2,3,4,6,7],[0,1,2,3,4,5,6,7],[0,1,5,6,7]]")
	
	paths := allPathsSourceTarget(graph)
	
	assert.Equal(t, expected, paths)
}

func strToArray(t *testing.T, s string) [][]int {
	var arr2d [][]int 
	split := strings.Split(s, "],[")

	var arrStr string 
	
	if len(split) == 1 {
		arrStr = strings.TrimPrefix(split[0], "[[")
		arrStr = strings.TrimSuffix(arrStr, "]]")
		arr2d = append(arr2d, splitComma(arrStr))
		return arr2d
	}

	arrStr = strings.TrimPrefix(split[0], "[[")
	arr2d = append(arr2d, splitComma(arrStr))

	for i := 1; i < len(split) - 1; i++ {
		arr2d = append(arr2d, splitComma(split[i]))
	}

	arrStr = strings.TrimSuffix(split[len(split)-1], "]]")
	arr2d = append(arr2d, splitComma(arrStr))

	return arr2d
}

func splitComma(s string) []int {
	split := strings.Split(s, ",")
	arr := make([]int, len(split))
	
	for i, str := range split {
		digit, err := strconv.Atoi(str)
		if err != nil {
			return []int{}
		}
		arr[i] = digit
	}

	return arr
}