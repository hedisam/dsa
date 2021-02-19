package Shortest_Reach

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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



////////////////////// the Problem /////////////////////////////////

// Complete the bfs function below.
func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {

}

func main() {

}