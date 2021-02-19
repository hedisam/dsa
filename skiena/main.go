package main

import (
	"github.com/hedisam/algods/skiena/graph"
	"os"
)

func main() {
	f, err := os.Open("skiena/graph/dag.txt")
	if err != nil {
		panic(err)
	}
	_, err = graph.ReadGraph(f, true)
	if err != nil {
		panic(err)
	}

	//fmt.Println(g)

	graph.SolveArithmetic()
}
