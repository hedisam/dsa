package graph

type PathFinder struct {
	source int
	bfs *BFS
}

func NewPathFinder(g *Graph, source int) *PathFinder {
	path := &PathFinder{
		source: source,
	}
	bfs := NewBFS()
	bfs.Search(g, source)
	return path
}

func (p *PathFinder) HasPath(x int) bool {
	return p.bfs.parent[x] > -1
}

func (p *PathFinder) PathTo(x int) []int {
	var path []int
	p.pathTo(x, &path)
	return path
}

func (p *PathFinder) pathTo(x int, path *[]int) {
	//if bfs.source != x {
	//	bfs.pathTo(bfs.parent[x], path)
	//}
	//*path = append(*path, x)
	if p.bfs.parent[x] == -1 {
		// no parent, no path
		return
	}
	if p.source == x {
		*path = append(*path, x)
	} else {
		p.pathTo(p.bfs.parent[x], path)
		*path = append(*path, x)
	}
}