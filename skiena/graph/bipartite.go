package graph

import "fmt"

type color int
const (
	uncolored color = iota
	black
	white
)

type Bipartite struct {
	colors []color
}

func (b *Bipartite) complement(c color) color {
	if c == white {return black}
	if c == black {return white}
	return uncolored
}

func IsBipartite(g *Graph) bool {
	b := &Bipartite{colors: make([]color, g.nVertices)}
	bipartite := true

	bfs := NewBFS()
	bfs.Init(g)
	bfs.edgeProcessor = func(u, v int) {
		if b.colors[u] == b.colors[v] {
			fmt.Printf("Not bipartite becuase of edge (%d, %d)\n", u, v)
			bipartite = false
			return
		}
		b.colors[v] = b.complement(b.colors[u])
	}

	for i := 0; i < g.nVertices; i++ {
		if !bfs.discovered[i] {
			b.colors[i] = white
			bfs.Search(g, i)
		}
		if !bipartite {return false}
	}
	return bipartite
}
