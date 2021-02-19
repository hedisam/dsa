package graph

import (
	"fmt"
	"github.com/hedisam/algods/general/stack"
)

func TopologicalSort(g *Graph) []int {
	s := stack.NewArrayStack(g.nVertices)

	dfs := NewDFS(g)
	dfs.lateProcessor = func(v int) {
		s.Push(v)
	}

	dfs.edgeProcessor = func(u, v int) {
		if dfs.parent[v] != u {
			dfs.finished = true
			fmt.Println("There's a cycle in the graph. not a DAG.")
		}
	}

	for i := 0; i < g.nVertices; i++ {
		if !dfs.discovered[i] {
			dfs.Search(i)
		}
	}

	return s.Array()
}
