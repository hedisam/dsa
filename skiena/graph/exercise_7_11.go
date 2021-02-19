package graph

import (
	"fmt"
	"github.com/hedisam/algods/general/stack"
)

type ArtValue struct {
	isOperator bool
	valueStr   string
	valueFlo   float64
}

func (a *ArtValue) Evaluate(x, y float64) float64 {
	switch a.valueStr {
	case "+":
		return x + y
	case "-":
		return x - y
	case "/":
		return x / y
	case "*":
		return x * y
	}
	return 0
}

// solving exercise 7.11 from the Algorithm Design Manual 3rd edition
func SolveArithmetic() {
	// 2+3*4+(3*4)/5 = 16.4
	g := NewGraph(8, true)
	g.InsertEdge(0, 1)
	g.InsertEdge(0, 2)
	g.InsertEdge(1, 3)
	g.InsertEdge(1, 4)
	g.InsertEdge(2, 4)
	g.InsertEdge(2, 5)
	g.InsertEdge(4, 6)
	g.InsertEdge(4, 7)

	fmt.Println(g)

	values := make(map[int]*ArtValue)
	values[0] = &ArtValue{isOperator: true, valueStr: "+"}
	values[1] = &ArtValue{isOperator: true, valueStr: "+"}
	values[2] = &ArtValue{isOperator: true, valueStr: "/"}
	values[3] = &ArtValue{valueFlo: 2}
	values[4] = &ArtValue{isOperator: true, valueStr: "*"}
	values[5] = &ArtValue{valueFlo: 5}
	values[6] = &ArtValue{valueFlo: 3}
	values[7] = &ArtValue{valueFlo: 4}

	s := stack.NewFloatArrayStack(g.nVertices)

	dfs := NewDFS(g)
	dfs.lateProcessor = func(v int) {
		if !values[v].isOperator {
			s.Push(values[v].valueFlo)
		} else {
			x := s.Pop()
			y := s.Pop()
			result := values[v].Evaluate(x, y)
			values[v].valueFlo = result
			values[v].isOperator = false
			s.Push(result)
		}
	}

	dfs.edgeProcessor = func(u, v int) {
		if dfs.processed[v] && dfs.entryTime[u] > dfs.entryTime[v] {
			// cross edge
			s.Push(values[v].valueFlo)
		}
	}

	for i := 0; i < g.nVertices; i++ {
		if !dfs.discovered[i] {
			dfs.Search(i)
		}
	}

	fmt.Println("Expression:", values[0].valueFlo)
}
