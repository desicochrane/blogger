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

	fibs := []Fib{FibCached, FibSumOfTuple, FibTupleTailRecursive, FibTupleTailIterative, FibTupleTailIterative2, FibTupleTailIterativeStorage}

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
func BenchmarkFibBySumOfTuple_12(b *testing.B) {
	benchmarkFib(b, FibSumOfTuple, 12)
}
func BenchmarkFibBySumOfTuple_40(b *testing.B) {
	benchmarkFib(b, FibSumOfTuple, 40)
}
func BenchmarkFibBySumOfTuple_90(b *testing.B) {
	benchmarkFib(b, FibSumOfTuple, 90)
}

// -----------------------------------------------------------------------------
func BenchmarkFibTupleTailRecursive_12(b *testing.B) {
	benchmarkFib(b, FibTupleTailRecursive, 12)
}
func BenchmarkFibTupleTailRecursive_40(b *testing.B) {
	benchmarkFib(b, FibTupleTailRecursive, 40)
}
func BenchmarkFibTupleTailRecursive_90(b *testing.B) {
	benchmarkFib(b, FibTupleTailRecursive, 90)
}

// -----------------------------------------------------------------------------
func BenchmarkFibTupleTailIterative_12(b *testing.B) {
	benchmarkFib(b, FibTupleTailIterative, 12)
}
func BenchmarkFibTupleTailIterative_40(b *testing.B) {
	benchmarkFib(b, FibTupleTailIterative, 40)
}
func BenchmarkFibTupleTailIterative_90(b *testing.B) {
	benchmarkFib(b, FibTupleTailIterative, 90)
}

// -----------------------------------------------------------------------------
func BenchmarkFibTupleTailIterative2_12(b *testing.B) {
	benchmarkFib(b, FibTupleTailIterative2, 12)
}
func BenchmarkFibTupleTailIterative2_40(b *testing.B) {
	benchmarkFib(b, FibTupleTailIterative2, 40)
}
func BenchmarkFibTupleTailIterative2_90(b *testing.B) {
	benchmarkFib(b, FibTupleTailIterative2, 90)
}

// -----------------------------------------------------------------------------
func BenchmarkFibTupleTailIterativeStorage_12(b *testing.B) {
	benchmarkFib(b, FibTupleTailIterativeStorage, 12)
}
func BenchmarkFibTupleTailIterativeStorage_40(b *testing.B) {
	benchmarkFib(b, FibTupleTailIterativeStorage, 40)
}
func BenchmarkFibTupleTailIterativeStorage_90(b *testing.B) {
	benchmarkFib(b, FibTupleTailIterativeStorage, 90)
}

// -----------------------------------------------------------------------------
func BenchmarkFibTupleTailIterativeStorage2_12(b *testing.B) {
	benchmarkFib(b, FibTupleTailIterativeStorage2, 12)
}
func BenchmarkFibTupleTailIterativeStorage2_40(b *testing.B) {
	benchmarkFib(b, FibTupleTailIterativeStorage2, 40)
}
func BenchmarkFibTupleTailIterativeStorage2_90(b *testing.B) {
	benchmarkFib(b, FibTupleTailIterativeStorage2, 90)
}
