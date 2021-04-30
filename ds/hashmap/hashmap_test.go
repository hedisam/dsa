package hashmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	m := NewHashMap()

	m.Put("name", "Hidayat")
	assert.Equal(t, 1, m.Size())

	name, ok := m.Get("name")
	assert.True(t, ok)
	assert.Equal(t, "Hidayat", name)

	m.Put("last_name", "Khezri")
	assert.Equal(t, 2, m.Size())

	m.Put("name", "Eddie")
	assert.Equal(t, 2, m.Size())

	name, ok = m.Get("name")
	assert.True(t, ok)
	assert.Equal(t, "Eddie", name)

	for i := 0; i < 2; i++ {
		m.Delete("name")
		assert.Equal(t, 1, m.Size())

		name, ok = m.Get("name")
		assert.False(t, ok)
		assert.Nil(t, name)
	}

	m.Put("", "invalid key")
	assert.Equal(t, 1, m.Size())

	nilValue, ok := m.Get("")
	assert.False(t, ok)
	assert.Nil(t, nilValue)

	m.Delete("")
	assert.Equal(t, 1, m.Size())
}

func TestRedistributing(t *testing.T) {
	m := &HashMap{
		h:     &mockHash{}, // helps us trigger the redistributing & size doubling part
		slots: make([]*Slot, 2),
		size:  0,
	}

	m.Put("A", "ASCII 65")
	m.Put("B", "ASCII 66")
	m.Put("B", "repetitive")
	m.Put("B", "repetitive")
	m.Put("B", "repetitive")
	m.Put("B", "repetitive")
	m.Put("B", "repetitive")
	m.Put("B", "repetitive")
	m.Put("B", "repetitive")
	m.Put("B", "repetitive")
	m.Put("B", "last B")

	assert.Equal(t, 2, m.Size())

	B, ok := m.Get("B")
	assert.True(t, ok)
	assert.Equal(t, "last B", B)

	m.Put("C", "ASCII 67")
	m.Put("D", "ASCII 68")
	m.Put("E", "ASCII 69")
	m.Put("F", "ASCII 70")
	m.Put("G", "ASCII 71")
	m.Put("H", "ASCII 72")
	m.Put("I", "ASCII 73")
	m.Put("J", "ASCII 74")

}
