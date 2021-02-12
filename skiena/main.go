package main

import (
	"fmt"
	"github.com/hedisam/algods/skiena/graph"
	"os"
)

func main() {
	f, err := os.Open("skiena/graph/graph_source.txt")
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
}
