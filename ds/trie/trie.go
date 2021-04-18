package trie

// Trie represents a trie (prefix tree) data structure
type Trie struct {
	root *Node 
}

// NewTrie returns a new trie object
func NewTrie() *Trie {
	return &Trie{
		root: NewNode(false),
	}
}

func (t *Trie) Insert(key string) {
	t.insert(t.root, key, 0)
}

func (t *Trie) insert(root *Node, key string, index int) {
	if index >= len(key) {
		return 
	}

	isWord := index == len(key) - 1

	_, ok := root.Children[key[index]]
	if !ok {
		root.Children[key[index]] = NewNode(isWord)
	} else if isWord {
		wordNode := root.Children[key[index]]
		wordNode.IsWord = true
	}

	t.insert(root.Children[key[index]], key, index+1)
}

func (t *Trie) Search(key string) bool {
	return t.search(t.root, key, 0)
}

func (t *Trie) search(root *Node, key string, index int) bool {
	_, ok := root.Children[key[index]]
	if !ok {
		return false 
	} else if index == len(key) - 1 { // are we at the last character of the key? if so, the node should be a word node to claim that we've found the key
		node := root.Children[key[index]]
		return node.IsWord
	}

	return t.search(root.Children[key[index]], key, index+1)
}

func (t *Trie) Delete(key string) {
	t.delete(t.root, key, 0)
}

func (t *Trie) delete(root *Node, key string, index int) {
	_, ok := root.Children[key[index]]
	if !ok {
		return 
	} else if index == len(key) - 1 {
		node := root.Children[key[index]]
		node.IsWord = false 
		return 
	}

	t.delete(root.Children[key[index]], key, index+1)
}

func (t *Trie) AutoComplete(query string) []string {
	return t.autoComplete(query)
}

func (t *Trie) autoComplete(query string) []string {
	// first we need to search for the query and find a node (the node with last char in the query) to start traversing from
	var root *Node = t.root 
	var ok bool 
	var index int 
search: 
	root, ok = root.Children[query[index]]
	if !ok {
		return []string{}
	} else if index == len(query) - 1 {
		// found the node to start with 		
	} else {
		index++
		goto search 
	}

	var words []string 
	t.content(root, &words, query)
	return words 
}

func (t *Trie) Content() []string {
	var words []string 
	t.content(t.root, &words, "") 
	return words 
}

func (t *Trie) content(root *Node, words *[]string, sb string) {
	for key, node := range root.Children {
		sb += string(key)
		if node.IsWord {
			*words = append(*words, sb)
		}
		t.content(node, words, sb)
		if len(root.Children) > 1 {
			sb = sb[:len(sb)-1]
		}
	}
}