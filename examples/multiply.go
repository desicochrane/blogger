package examples

// -----------------------------------------------------------------------------
type Multiply func(a int, n int) int

func NativeProduct(a int, n int) int {
	return a * n
}

// -----------------------------------------------------------------------------

func RepeatedAddition(a int, n int) int {
	product := 0
	for i := 1; i <= n; i++ {
		product += a
	}

	return product
}

// -----------------------------------------------------------------------------
func RecursiveDoubleHalf(a int, n int) int {
	if n == 1 {
		return a
	}

	result := RecursiveDoubleHalf(a<<1, n>>1)

	// if n is odd we need to add "a"
	if n&1 == 1 {
		return result + a
	}

	return result
}

// -----------------------------------------------------------------------------
func TailRecursiveDoubleHalf(a int, n int) int {
	if n == 1 {
		return a
	}

	error := 0
	if n&1 == 1 {
		error = a
	}

	return error + TailRecursiveDoubleHalf(a<<1, n>>1)
}

// -----------------------------------------------------------------------------
func StrictTailRecursiveDoubleHalf(a int, n int) int {
	return StrictTailRecursiveDoubleHalfAcc(0, a, n)
}

func StrictTailRecursiveDoubleHalfAcc(acc int, a int, n int) int {
	if n&1 == 1 {
		acc += a

		if n == 1 {
			return acc
		}
	}

	a = a << 1
	n = n >> 1
	return StrictTailRecursiveDoubleHalfAcc(acc, a, n)
}

// -----------------------------------------------------------------------------
func IterativeDoubleHalf(a int, n int) int {
	acc := 0

	for {
		if n&1 == 1 {
			acc += a

			if n == 1 {
				return acc
			}
		}

		a = a << 1
		n = n >> 1
	}
}
