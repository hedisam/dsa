package bst

import (
	. "github.com/hedisam/algods/ds/trees"
)

// BST represents a binary search tree
type BST struct {
	root *BinaryNode
}

func NewBST(root int) *BST {
	return &BST{root: NewBinaryNode(root)}
}

func (bst *BST) Search(key int) bool {
	return bst.search(bst.root, key)
}

func (bst *BST) search(node *BinaryNode, key int) bool {
	if node == nil {
		return false
	}

	if key < node.Key {
		return bst.search(node.Left, key)
	} else if key > node.Key {
		return bst.search(node.Right, key)
	}

	return true
}

func (bst *BST) Insert(key int) {
	bst.root = bst.insert(bst.root, key)
}

func (bst *BST) insert(node *BinaryNode, key int) *BinaryNode {
	if node == nil {
		return NewBinaryNode(key)
	}

	if key < node.Key {
		node.Left = bst.insert(node.Left, key)
	} else if key > node.Key {
		node.Right = bst.insert(node.Right, key)
	}

	return node
}

func (bst *BST) Delete(key int) {
	bst.root = bst.delete(bst.root, key)
}

func (bst *BST) delete(node *BinaryNode, key int) *BinaryNode {
	if node == nil {
		return nil
	}

	if key < node.Key {
		node.Left = bst.delete(node.Left, key)
	} else if key > node.Key {
		node.Right = bst.delete(node.Right, key)
	} else {
		// we found the node with our key, now delete it
		// node with only one child or no child
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// for nodes with two children we need to find the node's (inorder) successor (next biggest node after this node) which is the smallest node in its right subtree
		succKey := bst.successor(node)

		// replace current node's key with its successor's key, this way the target key will be removed (replaced) from the tree
		node.Key = succKey

		// now we have duplicate nodes, delete the successor node
		node.Right = bst.delete(node.Right, succKey)
	}

	return node
}

func (bst *BST) successor(root *BinaryNode) int {
	node := root.Right
	successorKey := node.Key

	for node.Left != nil {
		node = node.Left
		successorKey = node.Key
	}

	return successorKey
}

func (bst *BST) Inorder(writer func(key int)) {
	bst.inorder(bst.root, writer)
}

func (bst *BST) inorder(node *BinaryNode, w func(key int)) {
	if node == nil {
		return
	}
	bst.inorder(node.Left, w)
	w(node.Key)
	bst.inorder(node.Right, w)
}
