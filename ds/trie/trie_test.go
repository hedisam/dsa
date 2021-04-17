package trie

import "testing"

func Test(t *testing.T) {
	tr := NewTrie()

	tr.Insert("Hidayat")
	
	if !tr.Search("Hidayat") {
		t.Errorf("trie search failed: 'Hidayat' should've been found by Search")
	}

	tr.Insert("Hida")
	if !tr.Search("Hida") {
		t.Errorf("trie search failed: 'Hida' should've been found by Search")
	}

	tr.Delete("OK")

	if tr.Search("Hello") {
		t.Errorf("trie search failed: 'Hello' should've not been found by Search")
	}

	tr.Insert("Hello")

	if !tr.Search("Hello") {
		t.Errorf("trie search/insert failed: 'Hello' should be in the trie ")
	}

	tr.Delete("Hidayat")
	if tr.Search("Hidayat") {
		t.Errorf("trie delete failed: 'Hidayat' should've been deleted")
	}
}