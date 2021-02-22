package graph

import (
	"fmt"
	"io"
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
	g.insertEdge(x, y, 1, g.directed)
}

func (g *Graph) InsertWeightedEdge(x, y, weight int) {
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

// ReadGraph reads the graph's info from a io.Reader and returns the corresponding graph.
func ReadGraph(reader io.Reader, directed bool) (*Graph, error) {
	var g *Graph
	var n int // number of vertices
	var m int // number of edges

	_, err := fmt.Fscanf(reader, "%d %d\n", &n, &m)
	if err != nil {
		return nil, fmt.Errorf("bad format. expected first line to be [m n\n]: %w", err)
	}

	g = NewGraph(n, directed)

	var x, y int

	for i := 0; i < m; i++ {
		_, err = fmt.Fscanf(reader, "%d %d\n", &x, &y)
		if err != nil {
			return nil, fmt.Errorf("bad format. expected to have an edge as [x y]: %w", err)
		}
		g.InsertEdge(x, y)
	}

	return g, nil
}
