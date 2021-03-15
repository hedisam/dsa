package main

import "fmt"

func main() {
	grid := [][]int{
		{1,0,0,0},
		{0,0,0,0},
		{0,0,2,-1},
	}

	fmt.Println(uniquePathsIII(grid))
}

type path struct {
	visited [][]bool
	spaces int
	visitedNum int
}

var m, n int
// starting point
var si, sj int
func uniquePathsIII(grid [][]int) int {
	m = len(grid) - 1
	n = len(grid[0]) - 1

	p := &path{visited: make([][]bool, m+1), spaces: (m+1)*(n+1)}

	// counting the ending point
	p.visitedNum++

	for i := 0; i <= m; i++ {
		p.visited[i] = make([]bool, n+1)
		for j := 0; j <= n; j++ {
			if grid[i][j] == 1 {
				si, sj = i, j
			} else if grid[i][j] == -1 {
				p.visited[i][j] = true
				p.visitedNum++
			}
		}
	}
	return up3(grid, si, sj, p)
}

func up3(grid [][]int, i, j int, p *path) int {
	if p.visitedNum > p.spaces {return 0}
	if i > m || j > n || i < 0 || j < 0 {return 0}
	if p.visited[i][j] {return 0}

	// we've reched the end point
	if grid[i][j] == 2 {
		// have we walked through all free squares exactly once?
		if p.spaces == p.visitedNum {
			return 1
		}
		return 0
	}

	p.visitedNum++
	p.visited[i][j] = true

	down := up3(grid, i+1, j, p)
	up := up3(grid, i-1, j, p)
	right := up3(grid, i, j+1, p)
	left := up3(grid, i, j-1, p)
	p.visitedNum--
	p.visited[i][j] = false
	return down + up + right + left
}