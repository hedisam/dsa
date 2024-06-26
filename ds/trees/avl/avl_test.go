package avl

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

var keys = []int{2, 7, 6, 3, 10, 8, 40, 1}
var keysSorted = []int{1, 2, 3, 6, 7, 8, 10, 40}

func TestInsert(t *testing.T) {
	bt := buildAVL(keys)

	var output []int 
	w := func(k int) {
		output = append(output, k)
	}

	bt.InOrder(w)

	if !reflect.DeepEqual(keysSorted, output) {
		t.Errorf("bst insert failed:\nexpected: %v\ngot: %v", keys, output)
	}

}

func TestSearch(t *testing.T) {
	bt := buildAVL(keys)
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
	sbt := buildAVL(keys)

	testDelete(t, sbt)
}

func testDelete(t *testing.T, sbt *AVL) {
	key := sbt.root.Key
	sbt.Delete(key)

	// the bst must remain sorted 
	// and it should not find the key anymore
	var output []int 
	w := func(k int) {
		output = append(output, k)
		fmt.Printf("%d, ", k)
	}

	sbt.InOrder(w)

	if !sort.IntsAreSorted(output) {
		t.Errorf("bst delete failed: inorder traversal's output is not sorted anymore.")
	}

	if sbt.Search(key) {
		t.Errorf("bst delete failed: the deleted key was found after deletion")
	}

	testMaxBalance(t, sbt)
}

func TestMaxBalance(t *testing.T) {
	sbt := buildAVL(keys)
	testMaxBalance(t, sbt)
}

func testMaxBalance(t *testing.T, sbt *AVL) {
	maxBalance := sbt.MaxBalance()
	if maxBalance > 1 || maxBalance < -1 {
		t.Errorf("avl failed: max balance must be between -1 to 1, got: %d", maxBalance)
	}
}

func Test(t *testing.T) {
	arrrays := genRandomArrays(1000, 100)
	
	for _, arr := range arrrays {
		avlTree := buildAVL(arr)
		testMaxBalance(t, avlTree)
		testDelete(t, avlTree)
	}
}

func genRandomArrays(m, n int) [][]int {
	source := rand.NewSource(time.Now().Unix())
	rnd := rand.New(source)
	arrays := make([][]int, m)
	for i := 0; i < m; i++ {
		arrays[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arrays[i][j] = rnd.Int()
		}
	}

	return arrays
}

func buildAVL(keys []int) *AVL {
	sbt := &AVL{}

	for i := 0; i < len(keys); i++ {
		sbt.Insert(keys[i])
	}

	return sbt 
}









































