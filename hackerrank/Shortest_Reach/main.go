package Shortest_Reach

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
func (g *Graph) InsertEdge(x, y, weight int) {
	g.insertEdge(x, y, weight, g.directed)
}

func (g *Graph) insertEdge(x, y, weight int, directed bool) {
	edge := &EdgeNode{
		y:      y,
		weight: weight,
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
	g.insertEdge(y, x, weight, true)
}

// String returns a string displaying the graph.
func (g *Graph) String() string {
	var s strings.Builder
	var p *EdgeNode

	for i := 0; i < g.nVertices; i++ {
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
		edges: make([]*EdgeNode, n),
		degree: make([]int, n),
		nVertices: n,
		nEdges: 0,
		directed: directed,
	}
}

////////////////////////////// Queue Implementation ///////////////////////////////

type ArrayQueue struct {
	arr []int
}

func NewArrayQueue(cap int) *ArrayQueue {
	return &ArrayQueue{
		arr: make([]int, 0, cap),
	}
}

func (s *ArrayQueue) Enqueue(x int) {
	s.arr = append(s.arr, x)
}

func (s *ArrayQueue) Dequeue() int {
	x := s.arr[0]
	s.arr = s.arr[1:]
	return x
}

func (s *ArrayQueue) Top() int {
	return s.arr[0]
}

func (s *ArrayQueue) Empty() bool {
	return len(s.arr) == 0
}

func (s *ArrayQueue) Array() []int {
	return s.arr
}


//////////////////////////////// BFS ///////////////////////////

type BFS struct {
	processed []bool // which vertices have been processed
	discovered []bool // which vertices have been found
	parent []int // discovery relation

	earlyProcessor func(v int)
	edgeProcessor func(u, v int)
	lateProcessor func(v int)
}

func NewBFS(g *Graph) *BFS {
	s := &BFS{
		processed: make([]bool, g.nVertices),
		discovered: make([]bool, g.nVertices),
		parent: make([]int, g.nVertices),
		earlyProcessor: func(v int) {},
		edgeProcessor: func(u, v int) {},
		lateProcessor: func(v int) {},
	}
	s.init()
	return s
}

func (bfs *BFS) init() {
	for i := 0; i < len(bfs.parent); i++ {
		bfs.parent[i] = -1
	}
}

func (bfs *BFS) Search(g *Graph, source int) {
	q := NewArrayQueue(g.nVertices)

	q.Enqueue(source)
	bfs.discovered[source] = true

	for !q.Empty() {
		u := q.Dequeue()
		bfs.earlyProcessor(u)
		bfs.processed[u] = true
		for p := g.edges[u]; p != nil; p = p.next {
			v := p.y
			if !bfs.processed[v] || g.directed {
				bfs.edgeProcessor(u, v)
			}
			if !bfs.discovered[v] {
				q.Enqueue(v)
				bfs.discovered[v] = true
				bfs.parent[v] = u
			}
		}
		bfs.lateProcessor(u)
	}
}


////////////////////// the Problem /////////////////////////////////

// Complete the bfs function below.
func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	g := NewGraph(int(n), false)
	for _, edge := range edges {
		g.InsertEdge(int(edge[0]-1), int(edge[1]-1), 6)
	}

	levels := make([]int, n)
	dist := make([]int32, n)

	// vertices are indexed starting from 1, not zero
	s--
	search := NewBFS(g)
	search.earlyProcessor = func(v int) {
		if v != int(s) {
			levels[v] = levels[search.parent[v]] + 1
			dist[v] = int32(levels[v]) * 6
		}
	}

	search.Search(g, int(s))

	// removing the source from dist array
	dist = append(dist[:s], dist[s+1:]...)
	// setting unreachable nodes to -1
	for i := 0; i < len(dist); i++ {
		if dist[i] == 0 {
			dist[i] = -1
		}
	}
	return dist
}

func main() {

}