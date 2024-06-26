package graph

import "testing"

func TestKruskal(t *testing.T) {
	g := NewGraph(7, false)
	g.InsertWeightedEdge(0, 1, 5)
	g.InsertWeightedEdge(0, 2, 7)
	g.InsertWeightedEdge(0, 3, 12)
	g.InsertWeightedEdge(1, 2, 9)
	g.InsertWeightedEdge(1, 4, 7)
	g.InsertWeightedEdge(2, 3, 4)
	g.InsertWeightedEdge(2, 4, 4)
	g.InsertWeightedEdge(2, 5, 3)
	g.InsertWeightedEdge(3, 5, 7)
	g.InsertWeightedEdge(4, 5, 2)
	g.InsertWeightedEdge(4, 6, 5)
	g.InsertWeightedEdge(5, 6, 2)

	mstWeight := Kruskal(g)

	want := 23

	if want != mstWeight {
		t.Errorf("mst kruskal failed, mst W wanted: %d, got: %d\n", want, mstWeight)
	}
}
