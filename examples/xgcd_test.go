package examples

import (
	"testing"
)

func TestGCDNaive(t *testing.T) {
	samples := [][]int{
		{100, 0, 100},
		{16, 24, 8},
		{790933790547, 1849639579327, 8},
	}

	for _, gcd := range []GCD{GCDNaive, GCDDifference, GCDDifferenceIterative, GCDMod, GCDModIterative} {
		for _, sample := range samples {
			if result := gcd(sample[0], sample[1]); sample[2] != result {
				t.Fatalf("%d != %d", sample[2], result)
			}
		}
	}
}

// -----------------------------------------------------------------------------

func benchmarkGCD(b *testing.B, gcd GCD, m int, n int) {
	for i := 0; i < b.N; i++ {
		gcd(m, n)
	}
}

// -----------------------------------------------------------------------------
func BenchmarkGCDNaive_small(b *testing.B) {
	benchmarkGCD(b, GCDNaive, 16, 24)
}

// -----------------------------------------------------------------------------
func BenchmarkGCDDifference_small(b *testing.B) {
	benchmarkGCD(b, GCDDifference, 16, 24)
}

// -----------------------------------------------------------------------------
func BenchmarkGCDDifferenceIterative_small(b *testing.B) {
	benchmarkGCD(b, GCDDifferenceIterative, 16, 24)
}

// -----------------------------------------------------------------------------
func BenchmarkGCDMod_small(b *testing.B) {
	benchmarkGCD(b, GCDMod, 16, 24)
}

// -----------------------------------------------------------------------------
func BenchmarkGCDModIterative_small(b *testing.B) {
	benchmarkGCD(b, GCDModIterative, 16, 24)
}
