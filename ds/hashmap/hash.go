package hashmap

import "fmt"

// Hash is implemented by our Hash function providers
type Hash interface {
	WriteString(s string) (int, error)
	Reset()
	Sum64() uint64
}

// mockHash only works with single characters. It hashes a single character to its ascii equivalent.
type mockHash struct {
	r rune
}

func (m *mockHash) WriteString(s string) (int, error) {
	if len(s) < 1 {
		return 0, fmt.Errorf("invalid string: empty string")
	}
	r := []rune(s)
	m.r = r[0]
	return 1, nil
}

func (m *mockHash) Reset() {
	m.r = 0
}

func (m *mockHash) Sum64() uint64 {
	return uint64(m.r)
}



