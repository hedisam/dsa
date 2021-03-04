package graph

import (
	"fmt"
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

	heapCounter := 0
	for {
		for p := g.edges[v]; p != nil; p = p.next {
			if !visited[p.y] && dist[p.y] > dist[v] + p.weight {
				pq.InsertKey(&EdgeInfo{x: v, y: p.y, weight: dist[v] + p.weight})
			}
		}

		visited[v] = true

	nextShortPath:
		heapCounter++
		if shortestPath, ok := pq.ExtractMin(); !ok {
			// we're done
			break
		} else if shortestPath.weight >= dist[shortestPath.y] {
			goto nextShortPath
		} else {
			dist[shortestPath.y] = shortestPath.weight
			parent[shortestPath.y] = shortestPath.x
			v = shortestPath.y
		}
	}

	fmt.Printf("Edges: %d, HeapCounter: %d\n", g.nEdges, heapCounter)
	return &ShortestPath{
		source: source,
		parent: parent,
		dist:   dist,
	}
}
