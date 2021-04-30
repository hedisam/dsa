package hashmap

import "hash/maphash"

const (
	loadFactorThreshold = 0.75
)

type Node struct {
	Key string
	Value interface{}
}

type HashMap struct {
	// h provides the hash function
	h Hash
	// slots saves our hash nodes, each slot contains a chained list of nodes
	slots []*Slot
	// size is the number of items/nodes saved in the HashMap which is useful to calculate the load factor
	size int
}

func NewHashMap() *HashMap {
	return &HashMap{
		h: &maphash.Hash{},
		slots: make([]*Slot, 32),
	}
}

// Get returns the value corresponding the key saved in the HashMap. It'll return (nil, false) if the key's not in the
// HashMap, otherwise (the value, true) will be returned.
func (m *HashMap) Get(key string) (interface{}, bool) {
	if key == "" {
		return nil, false
	}

	index := m.index(key)

	slot := m.slots[index]
	if slot != nil {
		value := slot.get(key)
		if value != nil {
			return value, true
		}
	}

	return nil, false
}

// Size returns the number of key value pairs saved in the hash map.
func (m *HashMap) Size() int {
	return m.size
}

func (m *HashMap) Delete(key string) {
	if key == "" {
		return
	}

	index := m.index(key)
	slot := m.slots[index]
	if slot != nil {
		if slot.delete(key) {
			m.size--
		}
	}
}

// Put saves a (key, value) pair in the hash map. An existing value with the same key will get replaced by this new value.
// key-value pairs with key strings will be ignored.
func (m *HashMap) Put(key string, value interface{}) {
	if key == "" {
		return
	}

	m.put(&Node{Key: key, Value: value})

	// calculate the load factor to see if we need to grow the capacity of our hashmap
	loadFactor := float32(m.size) / float32(cap(m.slots))
	if loadFactor > loadFactorThreshold {
		m.doubleAndRedistribute()
	}
}

func (m *HashMap) put(node *Node) {
	index := m.index(node.Key)
	if m.slots[index] == nil {
		m.slots[index] = newSlot()
	}
	keyDuplicate := m.slots[index].put(node)
	if !keyDuplicate {
		m.size++
	}
}

func (m *HashMap) doubleAndRedistribute() {
	oldSlots := m.slots
	m.slots = make([]*Slot, cap(oldSlots) * 2)

	m.size = 0

	for _, slot := range oldSlots {
		if slot != nil {
			for _, node := range slot.nodes {
				m.put(&node)
			}
		}
	}
}

func (m *HashMap) index(key string) int {
	_, _ = m.h.WriteString(key)
	defer m.h.Reset()

	return int(m.h.Sum64() % uint64(cap(m.slots)))
}
