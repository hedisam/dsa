package graph

import (
	"github.com/hedisam/algods/ds/unionfind"
	"sort"
)

// Kruskal finds the min spanning tree for graph g and returns the total W of the mst.
func Kruskal(g *Graph) (weight int) {
	edges := edgesToArray(g)
	// sort (asc.) edges based on their W
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].W < edges[j].W
	})

	uf := unionfind.NewUnionFind(g.nVertices)

	i := 0
	for uf.Components() > 1 && i < len(edges) {
		x := edges[i].X
		y := edges[i].Y
		if !uf.Connected(x, y) {
			weight += edges[i].W
			uf.Unify(x, y)
		}
		i++
	}

	return
}

func edgesToArray(g *Graph) []*EdgeInfo {
	edges := make([]*EdgeInfo, g.nEdges)
	var i int
	bfs := NewBFS(g)
	bfs.edgeProcessor = func(x int, e *EdgeNode) {
		edges[i] = &EdgeInfo{X: x, Y: e.y, W: e.weight}
		i++
	}
	for v := 0; v < g.nVertices; v++ {
		if !bfs.discovered[v] {
			bfs.Search(g, v)
		}
	}
	return edges
}
