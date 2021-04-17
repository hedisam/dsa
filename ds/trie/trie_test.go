package trie

import (
	"reflect"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	t.Run("test content", func(t *testing.T) {
		tr := NewTrie()

		words := tr.Content()
		if len(words) > 0 {
			t.Errorf("trie Content failed: got %v, expected empty list", words)
		}

		input := []string{"Hida", "Hello", "Hi", "He", "What", "We", "With"}
		sort.Strings(input)

		for _, word := range input {
			tr.Insert(word)
		}

		words = tr.Content()
		sort.Strings(words)
		if !reflect.DeepEqual(words, input) {
			t.Errorf("trie content displaying failed: expected %v, got %v", input, words)
		}
	})

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