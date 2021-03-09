package main

import (
	"fmt"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	adj := make([][]EdgeInfo, n)
	for _, edge := range flights {
		adj[edge[0]] = append(adj[edge[0]], EdgeInfo{V: edge[1], W: edge[2]})
	}

	// using dijkstra
	pq := NewMinHeap(len(flights))
	pq.InsertKey(&EdgeInfo{V: src, W: 0, K: 0})

	for pq.Size() > 0 {
		path, _ := pq.ExtractMin()

		if path.K > K+1 {continue}

		if path.V == dst {
			return path.W
		}

		for _, e := range adj[path.V] {
			cost := path.W + e.W
			ek := path.K + 1
			if ek <= K + 1 {
				pq.InsertKey(&EdgeInfo{V: e.V, W: cost, K: path.K + 1})
			}
		}
	}

	return -1
}

func main() {
	input := `
5
[[0,1,5],[1,2,5],[0,3,2],[3,1,2],[1,4,1],[4,2,1]]
0
2
2
`
	// 7 wanted
	_, adj, _, _, _ := parseInput(input)
	toString(adj)

	result := findCheapestPrice(parseInput(input))
	fmt.Println(result)
}

