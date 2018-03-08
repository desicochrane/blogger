package examples

import (
	"testing"
)

func TestFib(t *testing.T) {
	samples := [][]int{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{12, 144},
		{40, 102334155},
		{50, 12586269025},
		{60, 1548008755920},
		{70, 190392490709135},
		{80, 23416728348467685},
		{90, 2880067194370816120},
	}

	fibs := []Fib{FibVecSum, FibCached, FibTupleTailRecursive, FibTupleTailIterative, FibTupleTailIterative2, FibTupleTailIterativeStorage}

	for _, f := range fibs {
		for _, sample := range samples {
			if result := f(sample[0]); result != sample[1] {
				t.Fatalf("%d != %d", f(sample[0]), sample[1])
			}
		}
	}
}

func benchmarkFib(b *testing.B, f Fib, n int) {
	for i := 0; i < b.N; i++ {
		f(n)
	}
}

// -----------------------------------------------------------------------------
func BenchmarkFibNaive_12(b *testing.B) {
	benchmarkFib(b, FibNaive, 12)
}
func BenchmarkFibNaive_40(b *testing.B) {
	benchmarkFib(b, FibNaive, 40)
}

//func BenchmarkFibNaive_90(b *testing.B) {
//	benchmarkFib(b, FibNaive, 90)
//}

// -----------------------------------------------------------------------------
func BenchmarkFibCached_12(b *testing.B) {
	benchmarkFib(b, FibCached, 12)
}
func BenchmarkFibCached_40(b *testing.B) {
	benchmarkFib(b, FibCached, 40)
}
func BenchmarkFibCached_90(b *testing.B) {
	benchmarkFib(b, FibCached, 90)
}

// -----------------------------------------------------------------------------
func BenchmarkFibSumVec_12(b *testing.B) {
	benchmarkFib(b, FibVecSum, 12)
}
func BenchmarkFibSumVec_40(b *testing.B) {
	benchmarkFib(b, FibVecSum, 40)
}
func BenchmarkFibSumVec_90(b *testing.B) {
	benchmarkFib(b, FibVecSum, 90)
}
