package examples

import "testing"

func TestMultiply(t *testing.T) {
	samples := [][]int{
		{2, 8, 16},
		{9, 16, 144},
		{2, 28, 56},
		{17, 28, 476},
		{19998, 12234, 244655532},
	}

	for _, fn := range []Multiply{
		NativeProduct,
		RepeatedAddition,
	} {
		for _, sample := range samples {
			r := fn(sample[0], sample[1])
			if r != sample[2] {
				t.Fatalf("%d * %d = %d, got %d", sample[0], sample[1], sample[2], r)
			}
		}
	}
}

func benchmarkMultiply(b *testing.B, a int, n int, multiply Multiply) {
	for i := 0; i < b.N; i++ {
		multiply(a, n)
	}
}

// NativeProduct
func BenchmarkNativeProduct_smallInput(b *testing.B) {
	benchmarkMultiply(b, 17, 28, NativeProduct)
}
func BenchmarkNativeProduct_largerInput(b *testing.B) {
	benchmarkMultiply(b, 19998, 12234, NativeProduct)
}

// RepeatedAddition
func BenchmarkRepeatedAddition_smallInput(b *testing.B) {
	benchmarkMultiply(b, 17, 28, RepeatedAddition)
}
func BenchmarkRepeatedAddition_largerInput(b *testing.B) {
	benchmarkMultiply(b, 19998, 12234, RepeatedAddition)
}
