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

	bfs := NewBFS(g)
	bfs.edgeProcessor = func(x int, e *EdgeNode) {
		if b.colors[x] == b.colors[e.y] {
			fmt.Printf("Not bipartite becuase of edge (%d, %d)\n", x, e.y)
			bipartite = false
			return
		}
		b.colors[e.y] = b.complement(b.colors[x])
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
