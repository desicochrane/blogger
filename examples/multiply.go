package examples

type Multiply func(a int, n int) int

// -----------------------------------------------------------------------------
func MultiplyNative(a int, n int) int {
	if n < 0 || a < 0 {
		panic("operand < 0")
	}

	return a * n
}

// -----------------------------------------------------------------------------
func Multiply0(a int, n int) int {
	if n < 0 || a < 0 {
		panic("operand < 0")
	}

	if n == 0 {
		return 0
	}

	return Multiply0(a, n-1) + a
}

// -----------------------------------------------------------------------------
func Multiply1(a int, n int) int {
	if n < 0 || a < 0 {
		panic("operand < 0")
	}

	if n == 0 {
		return 0
	}

	result := Multiply1(a<<1, n>>1)

	if n&0x1 == 1 {
		return result + a
	}

	return result
}

// -----------------------------------------------------------------------------
func Multiply2(a int, n int) int {
	if n < 0 || a < 0 {
		panic("operand < 0")
	}

	if n == 0 {
		return 0
	}

	var error int
	if n&0x1 == 1 {
		error = a
	} else {
		error = 0
	}

	a = a << 1
	n = n >> 1
	return error + Multiply2(a, n)
}

// -----------------------------------------------------------------------------
func Multiply2x(a int, n int) int {
	if n < 0 || a < 0 {
		panic("operand < 0")
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return a
	}

	var error int
	if n&0x1 == 1 {
		error = a
	} else {
		error = 0
	}

	a = a << 1
	n = n >> 1
	return error + Multiply2(a, n)
}

// -----------------------------------------------------------------------------
func Multiply3(a int, n int) int {
	if n < 0 || a < 0 {
		panic("operand < 0")
	}

	return Multiply3Acc(0, a, n)
}

func Multiply3Acc(result int, a int, n int) int {
	if n == 0 {
		return result
	}

	if n&0x1 == 1 {
		result += a
	}

	a = a << 1
	n = n >> 1
	return Multiply3Acc(result, a, n)
}

// -----------------------------------------------------------------------------
func Multiply4(a int, n int) int {
	if n < 0 || a < 0 {
		panic("operand < 0")
	}

	return Multiply4Acc(0, a, n)
}

func Multiply4Acc(result int, a int, n int) int {
	for {
		if n == 0 {
			return result
		}

		if n&0x1 == 1 {
			result += a
		}

		a = a << 1
		n = n >> 1
	}
}

// -----------------------------------------------------------------------------
func Multiply5(a int, n int) int {
	if n < 0 || a < 0 {
		panic("operand < 0")
	}

	return Multiply5Acc(0, a, n)
}

func Multiply5Acc(result int, a int, n int) int {
	if n == 0 {
		return 0
	}

	for {
		if n&0x1 == 1 {
			result += a
		}

		if n == 1 {
			return result
		}

		a = a << 1
		n = n >> 1
	}
}

// -----------------------------------------------------------------------------
func Multiply6(a int, n int) (result int) {
	if n < 0 || a < 0 {
		panic("operand < 0")
	}

	if n > a {
		a, n = n, a
	}

	if n == 0 {
		return 0
	}

	for {
		if n&0x1 == 1 {
			result += a

			if n == 1 {
				return result
			}
		}

		a = a << 1
		n = n >> 1
	}
}
