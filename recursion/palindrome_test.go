package recursion

import "testing"

func TestIsPalindrome(t *testing.T) {
	s := "tenet"
	wanted := true
	got := IsPalindrome(s)
	if wanted != got {
		t.Errorf("'%s': wanted %v, got %v\n", s, wanted, got)
	}
}
