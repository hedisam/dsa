package unionfind

type unionFind struct {
	// size of elements which need to be unified
	elemSize int
	// size of each component
	size []int
	parent []int
	// number of components
	compNum int
}

func NewUnionFind(size int) *unionFind {
	uf := &unionFind{
		elemSize: size,
		size:     make([]int, size),
		parent:   make([]int, size),
		compNum:  size,
	}
	for i := 0; i < size; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

// Find returns the component (the root node) that x belongs to, it will do path compression along the way.
func (uf *unionFind) Find(x int) int {
	if uf.parent[x] == x {
		return x
	}
	root := uf.Find(uf.parent[x])
	// path compression
	uf.parent[x] = root
	return root
}

// Unify the two components containing x and y
func (uf *unionFind) Unify(x, y int) {
	r1 := uf.Find(x)
	r2 := uf.Find(y)

	if r1 == r2 {return}

	if uf.size[r1] > uf.size[r2] {
		uf.size[r1] += uf.size[r2]
		uf.parent[r2] = r1
	} else {
		uf.size[r2] += uf.size[r1]
		uf.parent[r1] = r2
	}

	uf.compNum--
}

// Connected returns true if nodes x and y belong to the same component
func (uf *unionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Components returns the number of components, which could eventually goes down from uf.elemSize to 1
func (uf *unionFind) Components() int {
	return uf.compNum
}

// Size returns the number of elements that belong to the component containing node x
func (uf *unionFind) Size(x int) int {
	root := uf.Find(x)
	return uf.size[root]
}