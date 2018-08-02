package examples

import "testing"

func TestChoose(t *testing.T) {
	samples := [][]int{
		{5, 5, 1},
		{6, 4, 15},
		{100, 3, 161700},
		{100, 97, 161700},
		{40, 20, 137846528820},
	}

	for _, fn := range []Choose{ChooseBallsBoxes} {
		for _, s := range samples {
			if result := fn(s[0], s[1]); result != s[2] {
				t.Errorf("%d!=%d", s[2], result)
			}
		}
	}
}

// -----------------------------------------------------------------------------
func benchmarkChoose(b *testing.B, n int, k int, fn Choose) {
	for i := 0; i < b.N; i++ {
		fn(n, k)
	}
}

// -----------------------------------------------------------------------------
func BenchmarkChooseNaive_a(b *testing.B) {
	n, k := 5, 5
	benchmarkChoose(b, n, k, ChooseFactorialNaive)
}

//func BenchmarkChooseNaive_b(b *testing.B) {
//n, k := 100, 3
//benchmarkChoose(b, n, k, ChooseFactorialNaive)
//}

// -----------------------------------------------------------------------------
func BenchmarkChooseKFacNaive_a(b *testing.B) {
	n, k := 5, 5
	benchmarkChoose(b, n, k, ChooseKFacNaive)
}
func BenchmarkChooseKFacNaive_b(b *testing.B) {
	n, k := 100, 3
	benchmarkChoose(b, n, k, ChooseKFacNaive)
}
//func BenchmarkChooseKFac_c(b *testing.B) {
//	n, k := 100, 97
//	benchmarkChoose(b, n, k, ChooseKFacNaive)
//}
// -----------------------------------------------------------------------------
func BenchmarkChooseKFacCheckLarger_a(b *testing.B) {
	n, k := 5, 5
	benchmarkChoose(b, n, k, ChooseKFacCheckLarger)
}
func BenchmarkChooseKFacCheckLarger_b(b *testing.B) {
	n, k := 100, 3
	benchmarkChoose(b, n, k, ChooseKFacCheckLarger)
}
func BenchmarkChooseKFacCheckLarger_c(b *testing.B) {
	n, k := 100, 97
	benchmarkChoose(b, n, k, ChooseKFacCheckLarger)
}
// -----------------------------------------------------------------------------
func BenchmarkChooseRec_a(b *testing.B) {
	n, k := 5, 5
	benchmarkChoose(b, n, k, ChooseRec)
}
func BenchmarkChooseRec_b(b *testing.B) {
	n, k := 100, 3
	benchmarkChoose(b, n, k, ChooseRec)
}
func BenchmarkChooseRec_c(b *testing.B) {
	n, k := 100, 97
	benchmarkChoose(b, n, k, ChooseRec)
}
// -----------------------------------------------------------------------------
func BenchmarkChooseRecCached_a(b *testing.B) {
	n, k := 5, 5
	benchmarkChoose(b, n, k, ChooseRecCached)
}
func BenchmarkChooseRecCached_b(b *testing.B) {
	n, k := 100, 3
	benchmarkChoose(b, n, k, ChooseRecCached)
}
func BenchmarkChooseRecCached_c(b *testing.B) {
	n, k := 100, 97
	benchmarkChoose(b, n, k, ChooseRecCached)
}
func BenchmarkChooseRecCached_d(b *testing.B) {
	n, k := 40, 20
	benchmarkChoose(b, n, k, ChooseRecCached)
}
// -----------------------------------------------------------------------------
func BenchmarkChooseBallsBoxes_a(b *testing.B) {
	n, k := 5, 5
	benchmarkChoose(b, n, k, ChooseBallsBoxes)
}
func BenchmarkChooseBallsBoxes_b(b *testing.B) {
	n, k := 100, 3
	benchmarkChoose(b, n, k, ChooseBallsBoxes)
}
func BenchmarkChooseBallsBoxes_c(b *testing.B) {
	n, k := 100, 97
	benchmarkChoose(b, n, k, ChooseBallsBoxes)
}
func BenchmarkChooseBallsBoxes_d(b *testing.B) {
	n, k := 40, 20
	benchmarkChoose(b, n, k, ChooseBallsBoxes)
}
