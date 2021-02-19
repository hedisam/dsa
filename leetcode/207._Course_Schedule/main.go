package main

import "fmt"

/////////////////////////// Graph /////////////////////////////

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

/////////////////////////////////////////////// DFS ///////////////////////////////////////


type DFS struct {
	// which vertices have been processed
	processed []bool
	// which vertices have been found
	discovered []bool
	// discovery relation: who's the parent of the ith vertex
	parent []int

	// edgeProcessor process a newly discovered edge
	edgeProcessor func(u, v int)

	// finished is a flag to stop the search if necessary
	finished bool

	// g is the graph to be processed
	g *Graph
}

func NewDFS(g *Graph) *DFS {
	dfs := &DFS{
		processed:      make([]bool, g.nVertices),
		discovered:     make([]bool, g.nVertices),
		parent:         make([]int, g.nVertices),
		edgeProcessor:  func(u, v int) {},
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
	// stop searching if something's wrong or if we've already achieved our solution.
	if dfs.finished {return}

	// this is the entry phase for the vertex
	dfs.discovered[v] = true

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
	dfs.processed[v] = true

}

////////////////////////// The problem /////////////////////////////////

func canFinish(numCourses int, prerequisites [][]int) bool {
	// basically we're dealing a digraph. we need to check if it's a DAG or not. To be a DAG, there must be no
	// cycles.

	g := NewGraph(numCourses, true)
	for _, edge := range prerequisites {
		// the edges are represented in a reverse way.
		g.InsertEdge(edge[1], edge[0])
	}

	dag := true
	dfs := NewDFS(g)
	dfs.edgeProcessor = func(u, v int) {
		if dfs.parent[v] != u {
			// found  a cycle
			dfs.finished = true
			dag = false
		}
	}
	for i := 0; i < numCourses; i++ {
		if !dfs.discovered[i] {
			dfs.Search(i)
		}
		if dfs.finished {
			break
		}
	}

	return dag
}

func main() {
	// 20
	//[[0,10],[3,18],[5,5],[6,11],[11,14],[13,1],[15,1],[17,4]]
	edges := [][]int{{0,10},{3,18},{5,5},{6,11},{11,14},{13,1},{15,1},{17,4}}
	result := canFinish(20, edges)
	fmt.Println("can finish:", result)
}