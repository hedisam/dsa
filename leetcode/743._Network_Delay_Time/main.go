package main

import (
	"fmt"
)

const inf = 1<<31 - 1

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make([][]EdgeInfo, n+1)
	for _, edge := range times {
		adj[edge[0]] = append(adj[edge[0]], EdgeInfo{V: edge[1], W: edge[2]})
	}

	// using dijkstra
	pq := NewMinHeap(len(times))
	dist := make([]int, n+1)
	for i, _ := range dist {
		dist[i] = inf
	}
	dist[k] = 0
	dist[0] = 0 // vertices start from n=1

	pq.InsertKey(&EdgeInfo{V: k, W: 0})

	for pq.Size() > 0 {
		path, _ := pq.ExtractMin()
		if path.W < dist[path.V] {
			dist[path.V] = path.W
		}
		for _, e := range adj[path.V] {
			cost := e.W + path.W
			if cost < dist[e.V] {
				pq.InsertKey(&EdgeInfo{V: e.V, W: cost})
			}
		}
	}

	time := 0
	for _, d := range dist {
		if d >= inf {
			return -1
		}
		time = max(time, d)
	}

	return time
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := `
[[1,3,68],[1,4,20],[4,1,65],[3,2,74],[2,1,44],[3,4,61],[4,3,68],[3,1,26],[5,1,60],[5,3,3],[4,5,5],[2,5,36],[2,3,94],[1,2,0],[3,5,90],[2,4,28],[4,2,12],[5,4,52],[5,2,85],[1,5,42]]
5
4
`
	adj, n, src := parseInput(input)
	toString(adj)

	result := networkDelayTime(adj, n, src)
	fmt.Println(result)
}

