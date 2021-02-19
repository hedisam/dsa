package graph

import (
	"fmt"
	"strings"
)

type CC struct {
	cNumber int // component number
	components []int
}

func NewCC(g *Graph) *CC {
	cc := &CC{cNumber: 0, components: make([]int, g.nVertices)}
	cc.process(g)
	return cc
}

func (cc *CC) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("The graph has %d connected components\n", cc.cNumber))
	c := 0
	for i := 0; i < len(cc.components); i++ {
		if cc.components[i] != c {
			c++
			s.WriteString("\n")
			s.WriteString(fmt.Sprintf("[component %d]: ", c))
		}
		s.WriteString(fmt.Sprintf("%d, ", i))
	}
	return s.String()
}

func (cc *CC) Connected(x, y int) bool {
	return cc.components[x] == cc.components[y]
}

func (cc *CC) Count() int {
	return cc.cNumber + 1
}

func (cc *CC) Id(v int) int {
	return cc.components[v]
}

func (cc *CC) process(g *Graph) {
	bfs := NewBFS(g)
	bfs.earlyProcessor = func(v int) {
		cc.components[v] = cc.cNumber
		fmt.Printf("CC: vertex %d belongs to component %d\n", v, cc.cNumber)
	}

	for i := 0; i < g.nVertices; i++ {
		if !bfs.discovered[i] {
			cc.cNumber++
			bfs.Search(g, i)
		}
	}
}