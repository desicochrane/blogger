package examples

import (
	"math/rand"
	"fmt"
)

func RandomInt(min int, max int) int {
	return int((rand.Float64()*float64(max-min))+0.5) + min
}

func RandomPQ(min int, max int) (p int, q int) {
	for {
		if p = RandomInt(min, max); IsProbablePrime(p) {
			break
		}
	}

	for {
		if q = RandomInt(min, max); q != p && IsProbablePrime(q) {
			return p, q
		}
	}
}

func RandomCoprime(n int) int {
	for {
		if e := RandomInt(3, n); GCDModIter(e, n) == 1 {
			return e
		}
	}
}

func IsProbablePrime(p int) bool {
	if p < 2 {
		return false
	}

	if p == 2 || p == 3 {
		return true
	}

	guesses := 0
	for i := int(2); i < p && guesses < 50; i++ {
		guesses++

		k := int(rand.Float64()*float64(p-4)+0.5) + 2
		if k > p-2 || k < 2 {
			panic(fmt.Sprintf("%d %d", k, p))
		}
		if ModExpRecursive(k, p-1, p) != 1 {
			return false
		}
	}

	return true
}
