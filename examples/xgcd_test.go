package examples

import (
	"testing"
)

func TestGCD(t *testing.T) {
	samples := [][]int{
		{100, 0, 100},
		{16, 24, 8},
		{790933790547, 1849639579327, 3416723},
	}

	for _, gcd := range []GCD{
		//GCDNaive,
		GCDDifference,
		GCDDifferenceIterative,
		GCDMod,
		GCDModIterative,
	} {
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
func BenchmarkGCDDifference_large(b *testing.B) {
	benchmarkGCD(b, GCDDifference, 790933790547, 1849639579327)
}

// -----------------------------------------------------------------------------
func BenchmarkGCDDifferenceIterative_small(b *testing.B) {
	benchmarkGCD(b, GCDDifferenceIterative, 16, 24)
}
func BenchmarkGCDDifferenceIterative_large(b *testing.B) {
	benchmarkGCD(b, GCDDifferenceIterative, 790933790547, 1849639579327)
}

// -----------------------------------------------------------------------------
func BenchmarkGCDMod_small(b *testing.B) {
	benchmarkGCD(b, GCDMod, 16, 24)
}
func BenchmarkGCDMod_large(b *testing.B) {
	benchmarkGCD(b, GCDMod, 790933790547, 1849639579327)
}

// -----------------------------------------------------------------------------
func BenchmarkGCDModIterative_small(b *testing.B) {
	benchmarkGCD(b, GCDModIterative, 16, 24)
}
func BenchmarkGCDModIterative_large(b *testing.B) {
	benchmarkGCD(b, GCDModIterative, 790933790547, 1849639579327)
}

// -----------------------------------------------------------------------------
func TestXGCD(t *testing.T) {
	samples := [][]int{
		{100, 0, 1, 0},
		{24, 16, 1, -1},
		{240, 46, -9, 47},
		{10, 6, -1, 2},
		{7, 5, -2, 3},
		{391, 299, -3, 4},
		{239, 201, -37, 44},
		{1849639579327, 790933790547, -14500, 33909},
	}

	for _, xgcd := range []XGCD{
		XGCDRecursive,
		XGCDMatrix,
		XGCDAcc,
		XGCDAcc2,
		XGCDAcc3,
		XGCDWiki,
	} {
		for _, sample := range samples {
			if x, y := xgcd(sample[0], sample[1]); sample[2] != x || sample[3] != y {
				t.Fatalf("%d,%d != %d,%d", sample[2], sample[3], x, y)
			}
		}
	}
}

// -----------------------------------------------------------------------------
func benchmarkXGCD(b *testing.B, xgcd XGCD, m int, n int) {
	for i := 0; i < b.N; i++ {
		xgcd(m, n)
	}
}

// -----------------------------------------------------------------------------
func BenchmarkXGCDRecursive_small(b *testing.B) {
	benchmarkXGCD(b, XGCDRecursive, 16, 24)
}
func BenchmarkXGCDRecursive_large(b *testing.B) {
	benchmarkXGCD(b, XGCDRecursive, 790933790547, 1849639579327)
}
// -----------------------------------------------------------------------------
func BenchmarkXGCDMatrix_small(b *testing.B) {
	benchmarkXGCD(b, XGCDMatrix, 16, 24)
}
func BenchmarkXGCDMatrix_large(b *testing.B) {
	benchmarkXGCD(b, XGCDMatrix, 790933790547, 1849639579327)
}
// -----------------------------------------------------------------------------
func BenchmarkXGCDAcc_small(b *testing.B) {
	benchmarkXGCD(b, XGCDAcc, 16, 24)
}
func BenchmarkXGCDAcc_large(b *testing.B) {
	benchmarkXGCD(b, XGCDAcc, 790933790547, 1849639579327)
}

// -----------------------------------------------------------------------------
func BenchmarkXGCDAcc2_small(b *testing.B) {
	benchmarkXGCD(b, XGCDAcc2, 16, 24)
}
func BenchmarkXGCDAcc2_large(b *testing.B) {
	benchmarkXGCD(b, XGCDAcc2, 790933790547, 1849639579327)
}

// -----------------------------------------------------------------------------
func BenchmarkXGCDAcc3_small(b *testing.B) {
	benchmarkXGCD(b, XGCDAcc3, 16, 24)
}
func BenchmarkXGCDAcc3_large(b *testing.B) {
	benchmarkXGCD(b, XGCDAcc3, 790933790547, 1849639579327)
}
// -----------------------------------------------------------------------------
func BenchmarkXGCDWiki_small(b *testing.B) {
	benchmarkXGCD(b, XGCDWiki, 16, 24)
}
func BenchmarkXGCDWiki_large(b *testing.B) {
	benchmarkXGCD(b, XGCDWiki, 790933790547, 1849639579327)
}

