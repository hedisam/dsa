package main

import (
	"fmt"
	"github.com/hedisam/algods/skiena/graph"
	"os"
)

func main() {
	f, err := os.Open("skiena/graph/twocolors.txt")
	if err != nil {
		panic(err)
	}
	g, err := graph.ReadGraph(f, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(g)

	cc := graph.NewCC(g)
	fmt.Println(cc)

	fmt.Println("Is Bipartite:", graph.IsBipartite(g))

	x := 0
	y := 6
	fmt.Printf("Inserting edge (%d, %d)\n", x, y)
	g.InsertEdge(x, y)

	fmt.Println("Is Bipartite:", graph.IsBipartite(g))
}
