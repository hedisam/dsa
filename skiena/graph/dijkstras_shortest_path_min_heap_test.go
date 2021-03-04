package graph

import (
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	g := NewGraph(5, false)
	g.InsertWeightedEdge(0, 1, 4)
	g.InsertWeightedEdge(0, 2, 2)
	g.InsertWeightedEdge(1, 2, 3)
	g.InsertWeightedEdge(1, 3, 4)
	g.InsertWeightedEdge(1, 4, 2)
	g.InsertWeightedEdge(2, 3, 8)
	g.InsertWeightedEdge(3, 4, 1)

	pathFinder := Dijkstra(g, 0)
	path, cost := pathFinder.PathTo(3)
	fmt.Printf("Shortest Path from %d->%d with total cost of %d: %v\n", 0, 3, cost, path)

	got := cost
	want := 7
	if got != want {
		t.Errorf("wrong total cost of the shortest path, want: %d, got: %d", want, got)
	}
}
