package examples

import (
	"testing"
)

func TestModExp(t *testing.T) {
	samples := [][]int{
		{2, 2, 4, 0},
		{7, 4, 11, 3},
		{7, 128, 11, 9},
		{7, 13, 11, 2},
	}

	for _, mod := range []ModExp{ModExpRecursive, ModExpRecursivePow2, ModExpRecursivePow2R} {
		for _, s := range samples {
			result := mod(s[0], s[1], s[2])
			if result != s[3] {
				t.Fatalf("%d != %d", s[3], result)
			}
		}
	}
}

// -----------------------------------------------------------------------------
func benchmarkModExp(b *testing.B, mod ModExp, base, e, m int) {
	for i := 0; i < b.N; i++ {
		mod(base, e, m)
	}
}

// -----------------------------------------------------------------------------
func BenchmarkModExpNaive_small(b *testing.B) {
	benchmarkModExp(b, ModExpNaive, 7, 4, 11)
}

// -----------------------------------------------------------------------------
func BenchmarkModExpRecursive_small(b *testing.B) {
	benchmarkModExp(b, ModExpRecursive, 7, 4, 11)
}
func BenchmarkModExpRecursive_large(b *testing.B) {
	benchmarkModExp(b, ModExpRecursive, 7, 128, 11)
}
// -----------------------------------------------------------------------------
func BenchmarkModExpRecursivePow2_small(b *testing.B) {
	benchmarkModExp(b, ModExpRecursivePow2, 7, 4, 11)
}
func BenchmarkModExpRecursivePow2_large(b *testing.B) {
	benchmarkModExp(b, ModExpRecursivePow2, 7, 128, 11)
}
