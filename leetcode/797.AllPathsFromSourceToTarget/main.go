package main

func allPathsSourceTarget(graph [][]int) [][]int {
	var paths [][]int
	dfs(graph, &paths, nil, 0, len(graph)-1)

	return paths
}

// dfs depth-first-search
func dfs(g [][]int, allPaths *[][]int, path []int, src, target int) {
	path = append(path, src)

	if src == target {
		myPath := make([]int, len(path))
		copy(myPath, path)
		*allPaths = append(*allPaths, myPath)
		return
	}

	for i := 0; i < len(g[src]); i++ {
		dfs(g, allPaths, path, g[src][i], target)
	}
}