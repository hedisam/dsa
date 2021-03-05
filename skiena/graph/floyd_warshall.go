package graph

import "math"

type AdjacencyMatrix struct {
	// adjacency weight info. +Inf denotes having no edge.
	weight [][]int
	// number of vertices
	nVertices int
}

func Floyd(g *Graph) [][]int {
	adjM := ToAdjMatrix(g)

	for k := 0; k < adjM.nVertices; k++ {
		for i := 0; i < adjM.nVertices; i++ {
			for j := 0; j < adjM.nVertices; j++ {
				throughK := adjM.weight[i][k] + adjM.weight[k][j]
				if throughK < adjM.weight[i][j] {
					adjM.weight[i][j] = throughK
				}
			}
		}
	}
	return adjM.weight
}

func ToAdjMatrix(g *Graph) *AdjacencyMatrix {
	adjMatrix := &AdjacencyMatrix{
		weight:    make([][]int, g.nVertices),
		nVertices: g.nVertices,
	}

	for i := 0; i < g.nVertices; i++ {
		adjMatrix.weight[i] = make([]int, g.nVertices)
		for j := 0; j < g.nVertices; j++ {
			if i == j {
				adjMatrix.weight[i][j] = 0
				continue
			}
			adjMatrix.weight[i][j] = math.MaxInt64
		}
	}

	for i := 0; i < g.nVertices; i++ {
		for p := g.edges[i]; p != nil; p = p.next {
			adjMatrix.weight[i][p.y] = p.weight
		}
	}

	return adjMatrix
}


