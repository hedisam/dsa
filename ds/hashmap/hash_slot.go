package hashmap

type Slot struct {
	nodes []Node
}

func (s *Slot) put(node *Node) bool {
	// we'll remember one empty node (if any) to reuse its space.
	emptyIndex := -1
	for i := 0; i < len(s.nodes); i++ {
		if s.nodes[i].Key == "" { // this is an empty/nil node
			emptyIndex = i
			continue
		}
		if s.nodes[i].Key == node.Key {
			s.nodes[i].Value = node.Value
			return true
		}
	}
	if emptyIndex >= 0 {
		s.nodes[emptyIndex] = *node
		return false
	}
	s.nodes = append(s.nodes, *node)
	return false
}

func (s *Slot) get(key string) interface{} {
	for i := 0; i < len(s.nodes); i++ {
		if s.nodes[i].Key == key {
			return s.nodes[i].Value
		}
	}

	return nil
}

func (s *Slot) delete(key string) bool {
	for i := 0; i < len(s.nodes); i++ {
		if s.nodes[i].Key == key {
			s.nodes[i] = Node{}
			return true
		}
	}
	return false
}

func newSlot() *Slot {
	return &Slot{nodes: make([]Node, 1)}
}