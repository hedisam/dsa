package bst

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

var keys = []int{2, 7, 6, 3, 10, 8, 40, 1}
var keysSorted = []int{1, 2, 3, 6, 7, 8, 10, 40}



func TestInsert(t *testing.T) {
	bt := buildBST(keys)

	var output []int 
	w := func(k int) {
		output = append(output, k)
	}

	bt.Inorder(w)

	if !reflect.DeepEqual(keysSorted, output) {
		t.Errorf("bst insert failed:\nexpected: %v\ngot: %v", keys, output)
	}

}

func TestSearch(t *testing.T) {
	bt := buildBST(keys)
	t.Run("Key exists", func(t *testing.T) {
		key := 10 
		if !bt.Search(key) {
			t.Errorf("bst could not find key: %d", key)
		}
	})

	t.Run("Key doesn't exists", func(t *testing.T) {
		key := 150
		if bt.Search(key) {
			t.Errorf("bst found a non-existent key: %d", key)
		}
	})
}

func TestDelete(t *testing.T) {
	bt := buildBST(keys)

	key := 6
	bt.Delete(key)

	// the bst must remain sorted 
	// and it should not find the key anymore
	var output []int 
	w := func(k int) {
		output = append(output, k)
		fmt.Printf("%d, ", k)
	}

	bt.Inorder(w)

	if !sort.IntsAreSorted(output) {
		t.Errorf("bst delete failed: inorder traversal's output is not sorted anymore.")
	}

	if bt.Search(key) {
		t.Errorf("bst delete failed: the deleted key was found after deletion")
	}
}

func buildBST(keys []int) *BST {
	tree := &BST{}

	for i := 0; i < len(keys); i++ {
		tree.Insert(keys[i])
	}

	return tree 
}