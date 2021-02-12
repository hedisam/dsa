package graph

import (
	"fmt"
	"github.com/hedisam/algods/general/queue"
)

type BFS struct {
	g *Graph
	processed []bool // which vertices have been processed
	discovered []bool // which vertices have been found
	parent []int // discovery relation
	source int // source vertex
}

func NewBFS(g *Graph, source int) *BFS {
	bfs := &BFS{g: g, source: source}
	bfs.init()
	return bfs
}

func (bfs *BFS) init() {
	bfs.processed = make([]bool, bfs.g.nVertices)
	bfs.discovered = make([]bool, bfs.g.nVertices)
	bfs.parent = make([]int, bfs.g.nVertices)
	for i := 0; i < bfs.g.nVertices; i++ {
		bfs.parent[i] = -1
	}
}

func (bfs *BFS) HasPath(x int) bool {
	return bfs.parent[x] > -1
}

func (bfs *BFS) PathTo(x int) []int {
	var path []int
	bfs.pathTo(x, &path)
	return path
}

func (bfs *BFS) pathTo(x int, path *[]int) {
	//if bfs.source != x {
	//	bfs.pathTo(bfs.parent[x], path)
	//}
	//*path = append(*path, x)
	if bfs.parent[x] == -1 {
		// no parent, no path
		return
	}
	if bfs.source == x {
		*path = append(*path, x)
	} else {
		bfs.pathTo(bfs.parent[x], path)
		*path = append(*path, x)
	}
}

func (bfs *BFS) Run() {
	q := queue.NewArrayQueue(bfs.g.nVertices)

	q.Enqueue(bfs.source)
	bfs.discovered[bfs.source] = true

	for !q.Empty() {
		u := q.Dequeue()
		bfs.processVertexEarly(u)
		bfs.processed[u] = true
		for p := bfs.g.edges[u]; p != nil; p = p.next {
			v := p.y
			if !bfs.processed[v] || bfs.g.directed {
				bfs.processEdge(u, v)
			}
			if !bfs.discovered[v] {
				q.Enqueue(v)
				bfs.discovered[v] = true
				bfs.parent[v] = u
			}
		}
		bfs.processVertexLate(u)
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
