package trie 

// Node represents a trie node 
type Node struct {
	Children map[byte]*Node 
	IsWord bool 
}

func NewNode(isWord bool) *Node {
	return &Node{
		Children: make(map[byte]*Node),
		IsWord: isWord,
	}
}