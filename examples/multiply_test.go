package examples

import "testing"

func TestMultiply(t *testing.T) {
	samples := [][]int{
		{2, 8, 16},
		{9, 16, 144},
		{2, 28, 56},
	}

	for _, fn := range []Multiply{MultiplyNative, Multiply0, Multiply1} {
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

func BenchmarkMultiply0_small(b *testing.B) {
	benchmarkMultiply(b, 2, 16, Multiply0)
}
func BenchmarkMultiply0_large(b *testing.B) {
	benchmarkMultiply(b, 19998, 12234, Multiply0)
}

func BenchmarkMultiply1_small(b *testing.B) {
	benchmarkMultiply(b, 2, 16, Multiply1)
}

func BenchmarkMultiply1_large(b *testing.B) {
	benchmarkMultiply(b, 19998, 12234, Multiply1)
}

func benchmarkMultiply(b *testing.B, a int, n int, multiply Multiply) {
	for i := 0; i < b.N; i++ {
		multiply(a, n)
	}
}
