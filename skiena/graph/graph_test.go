package graph

import (
	"os"
	"testing"
)

func TestReadGraph(t *testing.T) {
	f, err := os.Open("graph_source.txt")
	if err != nil {
		t.Error(err)
	}

	g, err := ReadGraph(f, true)
	if err != nil {
		t.Error(err)
	}

	if g.nVertices != 5 {
		t.Errorf("expected to have 5 vertices, got: %d", g.nVertices)
	}

	if g.nEdges != 6 {
		t.Errorf("expected to have 6 vertices, got; %d", g.nEdges)
	}
}
