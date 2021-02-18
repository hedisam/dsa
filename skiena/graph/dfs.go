package graph

type DFS struct {
	// which vertices have been processed
	processed []bool
	// which vertices have been found
	discovered []bool
	// discovery relation: who's the parent of the ith vertex
	parent []int

	// earlyProcessor process the vertex v on entry (discovery time)
	earlyProcessor func(v int)
	// edgeProcessor process a newly discovered edge
	edgeProcessor func(u, v int)
	// lateProcessor process the vertex v on exit (finishing time)
	lateProcessor func(v int)

	// time increments on entry and exit of each vertex. helps us to keep a record of entry/exit time of each vertex
	time int
	// entryTime is a vertex based index array that keeps track of vertices' entry (discovery) time
	entryTime []int
	// exitTime is a vertex based index array that keeps track of vertices' exit (finishing) time
	exitTime []int

	// finished is a flag to stop the search if necessary
	finished bool

	// g is the graph to be processed
	g *Graph
}

func NewDFS(g *Graph) *DFS {
	dfs := &DFS{
		processed:      make([]bool, g.nVertices),
		discovered:     make([]bool, g.nVertices),
		parent:         make([]int, g.nVertices),
		earlyProcessor: func(v int) {},
		edgeProcessor:  func(u, v int) {},
		lateProcessor:  func(v int) {},
		entryTime: make([]int, g.nVertices),
		exitTime: make([]int, g.nVertices),
		g:              g,
	}
	dfs.init()
	return dfs
}

func (dfs *DFS) init() {
	for i := 0; i < len(dfs.parent); i++ {
		dfs.parent[i] = -1
	}
}

func (dfs *DFS) Search(v int) {
	// stop searching if something's wrong or if we've already achieved our solution.
	if dfs.finished {return}

	// this is the entry phase for the vertex
	dfs.discovered[v] = true
	dfs.time += 1
	// keep track of vertex's entry time
	dfs.entryTime[v] = dfs.time

	// do any early processing on the vertex if necessary
	dfs.earlyProcessor(v)

	// walking through v's adjacency list
	for e := dfs.g.edges[v]; e != nil; e = e.next {
		if !dfs.discovered[e.y] {
			dfs.parent[e.y] = v
			// process the tree-edge if desired
			dfs.edgeProcessor(v, e.y)

			dfs.Search(e.y)
		} else if dfs.g.directed || !dfs.processed[e.y] && dfs.parent[v] != e.y {
			// if vertex y is discovered already, we process the edge (v, e.y) only if:
			// the graph is directed, so the edge is directed, and every directed edge is important and should be processed
			// the graph is undirected: we might be looking at an edge like (y, x) from (x, y, x) so we don't need to
			// process such edges because we've already processed (x, y). if not so, we only check the edge if y has not
			// finished processing, which denotes that we're backing to an ancestor. Therefore it's a back-edge. if
			// e.y has finished processing, then we should've already visited vertex v, thus, the undirected edge (e.y, v)
			// has already been processed.
			// is it necessary to check if e.y is finished or not???
			dfs.edgeProcessor(v, e.y)
		}

		if dfs.finished {return}
	}

	// exit phase for the vertex

	// process the vertex on finishing time if necassry
	dfs.lateProcessor(v)
	// keep track of its exit time
	dfs.time += 1
	dfs.exitTime[v] = dfs.time

	dfs.processed[v] = true

}
