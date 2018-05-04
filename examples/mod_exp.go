package examples

import "math"

type ModExp func(b, e, m int) int

func ModExpNaive(b, e, m int) int {
	return int(math.Pow(float64(b), float64(e))) % m
}

// -----------------------------------------------------------------------------
func ModExpRecursive(b, e, m int) int {
	// b**e % me
	// b**(e-1) % m * b % m
	if e == 1 {
		return b % m
	}

	return (ModExpRecursive(b, e-1, m) * b) % m
}

// -----------------------------------------------------------------------------
func ModExpRecursivePow2(b, e, m int) int {
	if e == 1 {
		return b % m
	}

	x := ModExpRecursive(b, e>>1, m)

	if e&1 == 1 {
		return (x * x * (b % m)) % m
	} else {
		return (x * x) % m
	}
}

// -----------------------------------------------------------------------------
func ModExpIterative(b, e, m int) int {
	// 1. Set c = 1, e′ = 0.
	c, ep := 1, 0

	for {
		// 2. Increase e′ by 1.
		ep++

		// 3. Set c = (b ⋅ c) mod m.
		c = (b * c) % m

		// 4. If e′ < e, goto step 2. Else, c contains the correct solution to c ≡ be mod m.
		if ep == e {
			return c
		}
	}
}

// -----------------------------------------------------------------------------
func ModExpRecursivePow2R(b, e, m int) int {
	if e == 0 {
		return 1
	}

	f := 1
	if e&1 == 1 {
		f = b % m
	}

	x := ModExpRecursive(b, e>>1, m)

	return ((x * x % m) * f) % m
}

//func ModExpAux(acc, b, e, m int) int {
//
//}
