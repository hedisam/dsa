package recursion

import "testing"

func TestRaise(t *testing.T) {
	base := 2
	exp := 4
	want := 16

	got := Raise(base, exp)
	if got != want {
		t.Errorf("Rase %d to %d failed, wanted: %d, got: %d\n", base, exp, want, got)
	}
}

func TestRaiseOpt(t *testing.T) {
	base := 2
	exp := 4
	want := 16

	got := RaiseOpt(base, exp)
	if got != want {
		t.Errorf("RaseOpt %d to %d failed, wanted: %d, got: %d\n", base, exp, want, got)
	}
}

func BenchmarkRaise(b *testing.B) {
	base := 2
	exp := 1000

	b.Run("regular raise", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Raise(base, exp)
		}
	})

	b.Run("optimized raise", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RaiseOpt(base, exp)
		}
	})
}