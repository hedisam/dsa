package avl 

// AVLNode represents a node of an AVL tree
type AVLNode struct {
	Key int 
	Height int 
	Left *AVLNode
	Right *AVLNode
}

// AVL represents AVL self balancing search tree
type AVL struct {
	root *AVLNode
}

func height(node *AVLNode) int {
	if node == nil {
		return 0 
	}

	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (avl *AVL) rotateRight(y *AVLNode) *AVLNode {
	x := y.Left
	t2 := x.Right

	// perform rotation
	x.Right = y 
	y.Left = t2 

	// update heights 
	y.Height = max(height(y.Left), height(y.Right)) + 1 
	x.Height = max(height(x.Left), height(x.Right)) + 1 

	return x 
}

func (avl *AVL) rotateLeft(x *AVLNode) *AVLNode {
	y := x.Right
	t2 := y.Left

	y.Left = x 
	x.Right = t2 

	y.Height = max(height(y.Left), height(y.Right)) + 1 
	x.Height = max(height(x.Left), height(x.Right)) + 1 

	return y 
}

func (avl *AVL) balance(node *AVLNode) int {
	if node == nil {
		return 0 
	}

	return height(node.Left) - height(node.Right)
}

func (avl *AVL) Insert(key int) *AVLNode {
	return avl.insert(avl.root, key)
}

func (avl *AVL) insert(node *AVLNode, key int) *AVLNode {
	// performing normal bst insertion
	if node == nil {
		return NewAVLNode(key)
	}

	if key < node.Key {
		node.Left = avl.insert(node.Left, key)
	} else if key > node.Key {
		node.Right = avl.insert(node.Right, key)
	} else {
		// duplicate key 
		return node 
	}

	// update height of this ancestor node 
	node.Height = 1 + max(height(node.Left), height(node.Right))

	// check if this node has become unbalanced
	balance := avl.balance(node)

	// there would be 4 cases if the node is not balanced
	
	// left left case 
	if balance > 1 && key < node.Left.Key {
		return avl.rotateRight(node)
	}

	// right right case 
	if balance < -1 && key > node.Right.Key {
		return avl.rotateLeft(node)
	}

	// left right case 
	if balance > 1 && key > node.Left.Key {
		node.Left = avl.rotateLeft(node.Left)
		return avl.rotateRight(node)
	}

	// right left case 
	if balance < -1 && key < node.Right.Key {
		node.Right = avl.rotateRight(node.Right)
		return avl.rotateLeft(node)
	}

	// node is balanced
	return node 
}

func NewAVL(root int) *AVL {
	return &AVL{
		root: NewAVLNode(root),
	}
}

func NewAVLNode(key int) *AVLNode {
	return &AVLNode{
		Key: key,
		Height: 1,
	}
}













