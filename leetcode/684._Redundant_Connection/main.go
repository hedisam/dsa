package main

import (
	"fmt"
	"strings"
)

type EdgeNode struct {
	y int // adjacency info
	weight int // edge weigh, if any
	next *EdgeNode // next node in the adj list
}

type Graph struct {
	edges []*EdgeNode // adjacency info
	degree []int // out-degree of each vertex
	nVertices int // number of vertices in the graph
	nEdges int // number of edges in the graph
	directed bool // is the graph directed?
}

// InsertEdge from x to y.
func (g *Graph) InsertEdge(x, y int) {
	g.insertEdge(x, y, g.directed)
}

func (g *Graph) insertEdge(x, y int, directed bool) {
	edge := &EdgeNode{
		y:      y,
		weight: 1,
		// point next to the head of the x's adj list. then we assign this new edge to the head of the x's adj list.
		next:   g.edges[x],
	}

	// insert the new edge at thead of the adj list
	g.edges[x] = edge
	// increment the out-degree
	g.degree[x]++

	if directed {
		g.nEdges++
		return
	}
	g.insertEdge(y, x, true)
}

// String returns a string displaying the graph.
func (g *Graph) String() string {
	var s strings.Builder
	var p *EdgeNode

	for i := 1; i <= g.nVertices; i++ {
		s.WriteString(fmt.Sprintf("[%d] -> ", i))
		p = g.edges[i]
		for p != nil {
			s.WriteString(fmt.Sprintf("%d, ", p.y))
			p = p.next
		}
		s.WriteString("\n")
	}
	return s.String()
}

// NewGraph returns a new empty graph with n vertices.
func NewGraph(n int, directed bool) *Graph {
	return &Graph{
		edges: make([]*EdgeNode, n+1),
		degree: make([]int, n+1),
		nVertices: n,
		nEdges: 0,
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
	// edgeProcessor process a newly discovered edge
	edgeProcessor func(u, v int)
	// lateProcessor process the vertex v on exit (finishing time)
	lateProcessor func(v int)

	// time increments on entry and exit of each vertex. helps us to keep a record of entry/exit time of each vertex
	time int
	// entryTime is a vertex based index array that keeps track of vertices' entry (discovery) time
	entryTime []int
	// exitTime is a vertex based index array that keeps track of vertices' exit (finishing) time
	exitTime []int

	// finished is a flag to stop the search if necessary
	finished bool

	// g is the graph to be processed
	g *Graph
}

func NewDFS(g *Graph) *DFS {
	dfs := &DFS{
		processed:      make([]bool, g.nVertices+1),
		discovered:     make([]bool, g.nVertices+1),
		parent:         make([]int, g.nVertices+1),
		earlyProcessor: func(v int) {},
		edgeProcessor:  func(u, v int) {},
		lateProcessor:  func(v int) {},
		entryTime: make([]int, g.nVertices+1),
		exitTime: make([]int, g.nVertices+1),
		g:              g,
	}
	dfs.init()
	return dfs
}

func (dfs *DFS) init() {
	for i := 1; i < len(dfs.parent); i++ {
		dfs.parent[i] = -1
	}
}

func (dfs *DFS) Search(v int) {
	// stop searching if something's wrong or if we've already achieved our solution.
	if dfs.finished {return}

	// this is the entry phase for the vertex
	dfs.discovered[v] = true
	dfs.time += 1
	// keep track of vertex's entry time
	dfs.entryTime[v] = dfs.time

	// do any early processing on the vertex if necessary
	dfs.earlyProcessor(v)

	// walking through v's adjacency list
	for e := dfs.g.edges[v]; e != nil; e = e.next {
		if !dfs.discovered[e.y] {
			dfs.parent[e.y] = v
			// process the tree-edge if desired
			dfs.edgeProcessor(v, e.y)

			dfs.Search(e.y)
		} else if dfs.g.directed || !dfs.processed[e.y] && dfs.parent[v] != e.y {
			dfs.edgeProcessor(v, e.y)
		}

		if dfs.finished {return}
	}

	// exit phase for the vertex

	// process the vertex on finishing time if necessary
	dfs.lateProcessor(v)
	// keep track of its exit time
	dfs.time += 1
	dfs.exitTime[v] = dfs.time

	dfs.processed[v] = true

}

func FindRedundantConnection(edges [][]int) []int {
	n := edges[len(edges)-1][1]
	g := NewGraph(n, false)
	for i := len(edges)-1; i >= 0; i-- {
		g.InsertEdge(edges[i][0], edges[i][1])
	}

	backEdge := make([]int, 2)

	dfs := NewDFS(g)
	dfs.edgeProcessor = func(u, v int) {
		if dfs.parent[v] != u {
			// back-edge
			if u < v {
				backEdge[0] = u
				backEdge[1] = v
				return
			}
			backEdge[0] = v
			backEdge[1] = u
		}
	}
	dfs.Search(1)

	return backEdge
}

func main() {
	// todo: it's not just about finding a back-edge. the question requires us to return the last occurring edge
	// if there are multiple answers.
	edges := [][]int{{1,2}, {1,3}, {2,3}}
	result := FindRedundantConnection(edges)
	fmt.Println(result)
}

