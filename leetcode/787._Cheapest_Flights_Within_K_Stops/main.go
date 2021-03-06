package main

import (
	"fmt"
	"strconv"
	"strings"
)
var inf = 11000
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	adj := make([][]EdgeInfo, n)
	for _, edge := range flights {
		adj[edge[0]] = append(adj[edge[0]], EdgeInfo{V: edge[1], W: edge[2]})
	}

	cost := f(adj, src, dst, 0, K)
	if cost < inf {
		return cost
	}
	return -1
}

func f(adj [][]EdgeInfo, src, dst, depth, K int) int {
	if depth > K + 1 {
		return inf
	} else if src == dst {
		return 0
	}

	treecost := inf
	for _, edge := range adj[src] {
		cost := f(adj, edge.V, dst, depth+1, K) + edge.W
		if cost < treecost {
			treecost = cost
		}
	}
	return treecost
}

func main() {
	input := `
18
[[16,1,81],[15,13,47],[1,0,24],[5,10,21],[7,1,72],[0,4,88],[16,4,39],[9,3,25],[10,11,28],[13,8,93],[10,3,62],[14,0,38],[3,10,58],[3,12,46],[3,8,2],[10,16,27],[6,9,90],[14,8,6],[0,13,31],[6,4,65],[14,17,29],[13,17,64],[12,5,26],[12,1,9],[12,15,79],[16,11,79],[16,15,17],[4,0,21],[15,10,75],[3,17,23],[8,5,55],[9,4,19],[0,10,83],[3,7,17],[0,12,31],[11,5,34],[17,14,98],[11,14,85],[16,7,48],[12,6,86],[5,17,72],[4,12,5],[12,10,23],[3,2,31],[12,7,5],[6,13,30],[6,7,88],[2,17,88],[6,8,98],[0,7,69],[10,15,13],[16,14,24],[1,17,24],[13,9,82],[13,6,67],[15,11,72],[12,0,83],[1,4,37],[12,9,36],[9,17,81],[9,15,62],[8,15,71],[10,12,25],[7,6,23],[16,5,76],[7,17,4],[3,11,82],[2,11,71],[8,4,11],[14,10,51],[8,10,51],[4,1,57],[6,16,68],[3,9,100],[1,14,26],[10,7,14],[8,17,24],[1,11,10],[2,9,85],[9,6,49],[11,4,95]]
7
2
6
`
	// 7 wanted
	_, adj, _, _, _ := parseInput(input)
	toString(adj)

	result := findCheapestPrice(parseInput(input))
	fmt.Println(result)
}

func parseInput(in string) (n int, adj [][]int, src int, dst int, K int) {
	lines := strings.Split(in, "\n")
	n = toInt(strings.TrimSpace(lines[1]))
	adj = toAdjArray(strings.TrimSpace(lines[2]))
	src = toInt(strings.TrimSpace(lines[3]))
	dst = toInt(strings.TrimSpace(lines[4]))
	K = toInt(strings.TrimSpace(lines[5]))
	return
}

func toString(adj [][]int) {
	for _, edge := range adj {
		fmt.Printf("%v %v %v\n", edge[0], edge[1], edge[2])
	}
}

func toAdjArray(s string) (adj [][]int) {
	s1 := strings.Split(s, "],[")
	for i, a := range s1 {
		if i == 0 {
			a = strings.Replace(a, "[", "", -1)
			adj = append(adj, splitter(a))
		} else if i == len(s1) - 1 {
			a = strings.Replace(a, "]", "", -1)
			adj = append(adj, splitter(a))
		} else {
			adj = append(adj, splitter(a))
		}
	}
	return adj
}

func splitter(s string) []int {
	arr := make([]int, 3)
	s1 := strings.Split(s, ",")
	for i, a := range s1 {
		a = strings.TrimSpace(a)
		n := toInt(a)
		arr[i] = n
	}
	return arr
}

func toInt(a string) int {
	n, _ := strconv.Atoi(a)
	return n
}
