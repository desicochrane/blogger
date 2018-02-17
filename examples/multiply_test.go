package examples

import "testing"

func TestMultiply(t *testing.T) {
	samples := [][]int{
		{0, 8, 0},
		{2, 8, 16},
		{9, 16, 144},
		{2, 28, 56},
		{17, 28, 476},
	}

	for _, fn := range []Multiply{
		MultiplyNative,
		Multiply0,
		Multiply1,
		Multiply2,
		Multiply2x,
		Multiply3,
		Multiply4,
		Multiply5,
		Multiply6,
	} {
		for _, sample := range samples {
			r := fn(sample[0], sample[1])
			if r != sample[2] {
				t.Fatalf("%d * %d = %d, got %d", sample[0], sample[1], sample[2], r)
			}
		}
	}
}

func BenchmarkMultiplyNative_small(b *testing.B) {
	benchmarkMultiply(b, 2, 16, MultiplyNative)
}
func BenchmarkMultiplyNative_large(b *testing.B) {
	benchmarkMultiply(b, 19998, 12234, MultiplyNative)
}
func BenchmarkMultiplyNative_large_n_small(b *testing.B) {
	benchmarkMultiply(b, 2, 12234, MultiplyNative)
}

func BenchmarkMultiply0_small(b *testing.B) {
	benchmarkMultiply(b, 2, 16, Multiply0)
}
func BenchmarkMultiply0_large(b *testing.B) {
	benchmarkMultiply(b, 19998, 12234, Multiply0)
}
func BenchmarkMultiply0_large_n_small(b *testing.B) {
	benchmarkMultiply(b, 2, 12234, Multiply0)
}

func BenchmarkMultiply1_small(b *testing.B) {
	benchmarkMultiply(b, 2, 16, Multiply1)
}
func BenchmarkMultiply1_large(b *testing.B) {
	benchmarkMultiply(b, 19998, 12234, Multiply1)
}
func BenchmarkMultiply1_large_n_small(b *testing.B) {
	benchmarkMultiply(b, 2, 12234, Multiply1)
}

func BenchmarkMultiply2_small(b *testing.B) {
	benchmarkMultiply(b, 2, 16, Multiply2)
}
func BenchmarkMultiply2_large(b *testing.B) {
	benchmarkMultiply(b, 19998, 12234, Multiply2)
}
func BenchmarkMultiply2_large_n_small(b *testing.B) {
	benchmarkMultiply(b, 2, 12234, Multiply2)
}

func BenchmarkMultiply3_small(b *testing.B) {
	benchmarkMultiply(b, 2, 16, Multiply3)
}
func BenchmarkMultiply3_large(b *testing.B) {
	benchmarkMultiply(b, 19998, 12234, Multiply3)
}
func BenchmarkMultiply3_large_n_small(b *testing.B) {
	benchmarkMultiply(b, 2, 12234, Multiply3)
}

func BenchmarkMultiply4_small(b *testing.B) {
	benchmarkMultiply(b, 2, 16, Multiply4)
}
func BenchmarkMultiply4_large(b *testing.B) {
	benchmarkMultiply(b, 19998, 12234, Multiply4)
}
func BenchmarkMultiply4_large_n_small(b *testing.B) {
	benchmarkMultiply(b, 2, 12234, Multiply4)
}

func BenchmarkMultiply5_small(b *testing.B) {
	benchmarkMultiply(b, 2, 16, Multiply5)
}
func BenchmarkMultiply5_large(b *testing.B) {
	benchmarkMultiply(b, 19998, 12234, Multiply5)
}
func BenchmarkMultiply5_large_n_small(b *testing.B) {
	benchmarkMultiply(b, 2, 12234, Multiply5)
}

func BenchmarkMultiply6_small(b *testing.B) {
	benchmarkMultiply(b, 2, 16, Multiply6)
}
func BenchmarkMultiply6_large(b *testing.B) {
	benchmarkMultiply(b, 19998, 12234, Multiply6)
}
func BenchmarkMultiply6_large_n_small(b *testing.B) {
	benchmarkMultiply(b, 2, 12234, Multiply6)
}

func benchmarkMultiply(b *testing.B, a int, n int, multiply Multiply) {
	for i := 0; i < b.N; i++ {
		multiply(a, n)
	}
}
