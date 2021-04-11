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
	
	x.Height = max(height(x.Left), height(x.Right)) + 1 
	y.Height = max(height(y.Left), height(y.Right)) + 1 
	
	return y 
}

func (avl *AVL) balance(node *AVLNode) int {
	if node == nil {
		return 0 
	}

	return height(node.Left) - height(node.Right)
}

func (avl *AVL) Search(key int) bool {
	return avl.search(avl.root, key)
}

func (avl *AVL) search(root *AVLNode, key int) bool {
	if root == nil {
		return false 
	}

	if key < root.Key {
		return avl.search(root.Left, key)
	} else if key > root.Key {
		return avl.search(root.Right, key)
	} else {
		return true 
	}
}

func (avl *AVL) Insert(key int) {
	avl.root = avl.insert(avl.root, key)
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

func (avl *AVL) Delete(key int) {
	avl.root = avl.delete(avl.root, key)
}

func (avl *AVL) delete(root *AVLNode, key int) *AVLNode {
	// perform a std bst delete 
	if root == nil {return nil}

	if key < root.Key {
		root.Left = avl.delete(root.Left, key)
	} else if key > root.Key {
		root.Right = avl.delete(root.Right, key)
	} else {
		// found it
		if root.Right != nil && root.Left != nil {
			// having two children
			successorNode := avl.successorNode(root)
			root.Key = successorNode.Key
			root.Right = avl.delete(root.Right, root.Key)
		} else if root.Left != nil {
			// having one child (the left one)
			*root = *root.Left
		} else if root.Right != nil {
			// having one child (the right one)
			*root = *root.Right
		} else {
			// no children (so it's a leaf node)
			return nil 
		}
	}

	// update height of the current node 
	root.Height = max(height(root.Left), height(root.Right)) + 1

	// check if this subtree rooted at the current node is still balanced or not

	balance := avl.balance(root)

	// we will have 4 cases in case of having an unbalanced subtree 

	// left left case 
	if balance > 1 && avl.balance(root.Left) >= 0 {
		return avl.rotateRight(root)
	}
	
	// left right case 
	if balance > 1 && avl.balance(root.Left) < 0 {
		root.Left = avl.rotateLeft(root.Left)
		return avl.rotateRight(root)
	}

	// right right case 
	if balance < -1 && avl.balance(root.Right) <= 0 {
		return avl.rotateLeft(root)
	}

	// right left case 
	if balance < -1  && avl.balance(root.Right) > 0 {
		root.Right = avl.rotateRight(root.Right)
		return avl.rotateLeft(root)
	}

	return root 
}

// successorNode returns the next largest element of the node 
func (avl *AVL) successorNode(root *AVLNode) *AVLNode {
	if root == nil {return nil}
	var node *AVLNode
	for node = root.Right; node.Left != nil; node = node.Left {} 
	return node 
}

func (avl *AVL) InOrder(writer func(key int)) {
	avl.inOrder(avl.root, writer)
}

func (avl *AVL) inOrder(node *AVLNode, w func(key int)) {
	if node == nil {
		return
	}
	avl.inOrder(node.Left, w)
	w(node.Key)
	avl.inOrder(node.Right, w)
}

func (avl *AVL) MaxBalance() int {
	return avl.maxBalance(avl.root)
}

func (avl *AVL) maxBalance(root *AVLNode) int {
	rootBalance := avl.balance(root)
	leftBalance := avl.balance(root.Left)
	rightBalance := avl.balance(root.Right)

	max := rootBalance
	if leftBalance > max {
		max = leftBalance
	}
	if rightBalance > max {
		max = rightBalance
	}
	return max
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













