package examples

type Multiply func(a int, n int) int

func MultiplyNative(a int, n int) int {
	if n < 0 || a < 0 {
		panic("operand < 0")
	}

	return a * n
}

func Multiply0(a int, n int) int {
	if n < 0 || a < 0 {
		panic("operand < 0")
	}

	if n == 0 {
		return 0
	}

	return Multiply0(a, n-1) + a
}

func Multiply1(a int, n int) int {
	if n < 0 || a < 0 {
		panic("operand < 0")
	}

	if n == 1 {
		return a
	}

	result := Multiply1(a<<1, n>>1)

	if n&0x1 == 1 {
		return result + a
	}

	return result
}
