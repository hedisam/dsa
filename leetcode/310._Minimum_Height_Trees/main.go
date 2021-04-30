package main

import (
	"fmt"
)

type EdgeNode struct {
	y int // adjacency info
	next *EdgeNode // next node in the adj list
}

type Graph struct {
	edges []*EdgeNode // adjacency info
	nVertices int // number of vertices in the graph
	directed bool // is the graph directed?
}

// InsertEdge from x to y.
func (g *Graph) InsertEdge(x, y int) {
	g.insertEdge(x, y, g.directed)
}

func (g *Graph) insertEdge(x, y int, directed bool) {
	edge := &EdgeNode{
		y:      y,
		// point next to the head of the x's adj list. then we assign this new edge to the head of the x's adj list.
		next:   g.edges[x],
	}

	// insert the new edge at thead of the adj list
	g.edges[x] = edge
	// increment the out-degree

	if directed {
		return
	}
	g.insertEdge(y, x, true)
}

// NewGraph returns a new empty graph with n vertices.
func NewGraph(n int, directed bool) *Graph {
	return &Graph{
		edges: make([]*EdgeNode, n),
		nVertices: n,
		directed: directed,
	}
}

type DFS struct {
	// which vertices have been processed
	processed []bool
	// which vertices have been found
	discovered []bool
	// discovery relation: who's the parent of the ith vertex
	parent []int

	// earlyProcessor process the vertex v on entry (discovery time)
	earlyProcessor func(v int)

	// g is the graph to be processed
	g *Graph
}

func NewDFS(g *Graph) *DFS {
	dfs := &DFS{
		processed:      make([]bool, g.nVertices),
		discovered:     make([]bool, g.nVertices),
		parent:         make([]int, g.nVertices),
		earlyProcessor: func(v int) {},
		g:              g,
	}
	dfs.init()
	return dfs
}

func (dfs *DFS) init() {
	for i := 0; i < len(dfs.parent); i++ {
		dfs.parent[i] = -1
	}
}

func (dfs *DFS) Search(v int) {
	// this is the entry phase for the vertex
	dfs.discovered[v] = true

	// do any early processing on the vertex if necessary
	dfs.earlyProcessor(v)

	// walking through v's adjacency list
	for e := dfs.g.edges[v]; e != nil; e = e.next {
		if !dfs.discovered[e.y] {
			dfs.parent[e.y] = v
			dfs.Search(e.y)
		}
	}

	dfs.processed[v] = true

}

func findPath(start, end int, parent []int) (path []int) {
	for x := end; x != start; x = parent[x] {
		path = append(path, x)
	}
	path = append(path, start)
	return
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if len(edges) == 0 {
		vs := make([]int, n)
		for i := 0; i < n; i++ {
			vs[i] = i
		}
		return vs
	}

	g := NewGraph(n, false)
	// inserting edges
	for i := 0; i < len(edges); i++ {
		g.InsertEdge(edges[i][0], edges[i][1])
	}

	levels := make([]int, n)
	maxLevel := 0
	deepestNode := -1

	dfs := NewDFS(g)

	earlyP := func(v int) {
		if dfs.parent[v] >= 0 {
			levels[v] = levels[dfs.parent[v]] + 1
			if levels[v] > maxLevel {
				maxLevel = levels[v]
				deepestNode = v
			}
		}
	}
	dfs.earlyProcessor = earlyP
	dfs.Search(0)

	source := deepestNode

	levels = make([]int, n)
	maxLevel = 0
	deepestNode = -1

	dfs = NewDFS(g)
	dfs.earlyProcessor = earlyP
	dfs.Search(source)

	path := findPath(source, deepestNode, dfs.parent)

	var vertices []int
	mid := len(path) / 2
	vertices = append(vertices, path[mid])
	if len(path) % 2 == 0 {
		vertices = append(vertices, path[mid-1])
	}

	return vertices
}

func main() {
	// [[3,0],[3,1],[3,2],[3,4],[5,4]]
	edges := [][]int{{3,0}, {3,1}, {3,2}, {3,4}, {5,4}}
	result := findMinHeightTrees(6, edges)
	fmt.Println(result)
}