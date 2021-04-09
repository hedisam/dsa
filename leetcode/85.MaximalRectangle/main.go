package main

import "fmt"

var zero byte = 48
var one byte = 49

func main() {
	m := [][]byte{
		{'1','0','1','0','0'},
		{'1','0','1','1','1'},
		{'1','1','1','1','1'},
		{'1','0','0','1','0'},
	}
	//m := [][]byte{
	//	{one, zero, zero, zero, one},
	//	{one, one, zero, one, one},
	//	{one, one, one, one, one},
	//}
	fmt.Println(maximalRectangle(m))
}

func maximalRectangle(matrix [][]byte) int {
	return maxRect(matrix, 0, 0, 0, 0)
}

func maxRect(matrix [][]byte, i, j, iEnd, jEnd int) int {
	if iEnd >= len(matrix) || jEnd >= len(matrix[0]) {
		return 0
	}

	area := 0
	for k := i; k <= iEnd; k++ {
		for l := j; l <= jEnd; l++ {
			if matrix[k][l] == '0' {
				a := maxRect(matrix, k+1, l, iEnd, jEnd)
				b := maxRect(matrix, k, l+1, iEnd, jEnd)

				return max(a, b, area)
			}
			area++
		}
	}

	a := maxRect(matrix, i, j, iEnd+1, jEnd)
	b := maxRect(matrix, i, j, iEnd, jEnd+1)

	return max(a, b, area)
}

func max(a, b, c int) int {
	v := a
	if b > v {
		v = b
	}
	if c > v {
		v = c
	}
	return v
}