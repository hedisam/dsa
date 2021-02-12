package graph

import (
	"fmt"
	"github.com/hedisam/algods/general/queue"
)

type BFS struct {
	processed []bool // which vertices have been processed
	discovered []bool // which vertices have been found
	parent []int // discovery relation

	earlyProcessor func(v int)
	edgeProcessor func(u, v int)
	lateProcessor func(v int)

	initialized bool
}

func NewBFS() *BFS {
	s := &BFS{
		earlyProcessor: func(v int) {},
		edgeProcessor: func(u, v int) {},
		lateProcessor: func(v int) {},
	}
	return s
}

func (bfs *BFS) Init(g *Graph) {
	if bfs.initialized {
		return
	}
	bfs.initialized = true

	bfs.processed = make([]bool, g.nVertices)
	bfs.discovered = make([]bool, g.nVertices)
	bfs.parent = make([]int, g.nVertices)
	for i := 0; i < g.nVertices; i++ {
		bfs.parent[i] = -1
	}
}

func (bfs *BFS) Search(g *Graph, source int) {
	bfs.Init(g)

	q := queue.NewArrayQueue(g.nVertices)

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

func (bfs *BFS) processVertexEarly(u int) {
	fmt.Println("BFS: Processed Vertex:", u)
}

func (bfs *BFS) processEdge(u, v int) {
	fmt.Printf("BFS: Processed Edge (%d, %d)\n", u, v)
}

func (bfs *BFS) processVertexLate(u int) {

}
