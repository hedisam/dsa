package graph

import (
	"math"
)

func Prim(g *Graph, source int) (weight int) {
	parent := make([]int, g.nVertices)
	// is the vertex in the tree
	intree := make([]bool, g.nVertices)
	// cost of adding to the tree
	distance := make([]int, g.nVertices)
	// cheapest cost to enlarge the tree
	var dist int

	for i := 0; i < g.nVertices; i++ {
		parent[i] = -1
		distance[i] = math.MaxInt64
	}

	distance[source] = 0

	for v := source; !intree[v]; {
		intree[v] = true
		weight += dist

		for p := g.edges[v]; p != nil; p = p.next {
			if distance[p.y] > p.weight && !intree[p.y] {
				distance[p.y] = p.weight
				parent[p.y] = v
			}
		}

		dist = math.MaxInt64
		for i := 0; i < g.nVertices; i++ {
			if !intree[i] && distance[i] < dist {
				dist = distance[i]
				v = i
			}
		}
	}

	return
}
