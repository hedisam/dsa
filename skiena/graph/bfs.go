package graph

import (
	"github.com/hedisam/algods/ds/queue"
)

type BFS struct {
	processed []bool // which vertices have been processed
	discovered []bool // which vertices have been found
	parent []int // discovery relation

	earlyProcessor func(v int)
	edgeProcessor func(x int, edge *EdgeNode)
	lateProcessor func(v int)
}

func NewBFS(g *Graph) *BFS {
	s := &BFS{
		processed: make([]bool, g.nVertices),
		discovered: make([]bool, g.nVertices),
		parent: make([]int, g.nVertices),
		earlyProcessor: func(v int) {},
		edgeProcessor: func(x int, edge *EdgeNode) {},
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
				bfs.edgeProcessor(u, p)
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
