package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(in string) (n int, adj [][]int, src int, dst int, K int) {
	lines := strings.Split(in, "\n")
	n = toInt(strings.TrimSpace(lines[1]))
	adj = toAdjArray(strings.TrimSpace(lines[2]))
	src = toInt(strings.TrimSpace(lines[3]))
	dst = toInt(strings.TrimSpace(lines[4]))
	K = toInt(strings.TrimSpace(lines[5]))
	return
}

func toString(adj [][]int) {
	for _, edge := range adj {
		fmt.Printf("%v %v %v\n", edge[0], edge[1], edge[2])
	}
}

func toAdjArray(s string) (adj [][]int) {
	s1 := strings.Split(s, "],[")
	for i, a := range s1 {
		if i == 0 {
			a = strings.Replace(a, "[", "", -1)
			adj = append(adj, splitter(a))
		} else if i == len(s1) - 1 {
			a = strings.Replace(a, "]", "", -1)
			adj = append(adj, splitter(a))
		} else {
			adj = append(adj, splitter(a))
		}
	}
	return adj
}

func splitter(s string) []int {
	arr := make([]int, 3)
	s1 := strings.Split(s, ",")
	for i, a := range s1 {
		a = strings.TrimSpace(a)
		n := toInt(a)
		arr[i] = n
	}
	return arr
}

func toInt(a string) int {
	n, _ := strconv.Atoi(a)
	return n
}

