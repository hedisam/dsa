package graph

import (
	"math"
)

type ShortestPath struct {
	source int
	parent []int
	dist []int
}

func (p *ShortestPath) PathTo(y int) ([]int, int ) {
	var path []int
	for x := y; x != -1; x = p.parent[x] {
		path = append(path, x)
	}
	return path, p.dist[y]
}

func Dijkstra(g *Graph, source int) *ShortestPath {
	pq := NewMinHeap(g.nEdges)
	parent := make([]int, g.nVertices)
	dist := make([]int, g.nVertices)
	visited := make([]bool, g.nVertices)

	for i := 0; i < g.nVertices; i++ {
		parent[i] = -1
		dist[i] = math.MaxInt64
	}

	dist[source] = 0
	v := source

	for {
		for p := g.edges[v]; p != nil; p = p.next {
			if !visited[p.y] && dist[p.y] > dist[v] + p.weight {
				pq.InsertKey(&EdgeInfo{X: v, Y: p.y, W: dist[v] + p.weight})
			}
		}

		visited[v] = true

	nextShortPath:
		if shortestPath, ok := pq.ExtractMin(); !ok {
			// we're done
			break
		} else if shortestPath.W >= dist[shortestPath.Y] {
			goto nextShortPath
		} else {
			dist[shortestPath.Y] = shortestPath.W
			parent[shortestPath.Y] = shortestPath.X
			v = shortestPath.Y
		}
	}

	return &ShortestPath{
		source: source,
		parent: parent,
		dist:   dist,
	}
}
