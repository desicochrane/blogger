package examples

type GCD func(a int, b int) int

func GCDNaive(a int, b int) int {
	if b > a {
		a, b = b, a
	}

	if b == 0 {
		return a
	}

	for i := b; i > 0; i-- {
		if b%i == 0 && a%i == 0 {
			return i
		}
	}

	return 1
}

// -----------------------------------------------------------------------------

func GCDDifference(a int, b int) int {
	if b > a {
		a, b = b, a
	}

	if b == 0 {
		return a
	}

	return GCDDifference(b, a-b)
}

func GCDDifferenceIterative(a int, b int) int {
	for {
		if b > a {
			a, b = b, a
		}

		if b == 0 {
			return a
		}

		a, b = b, a-b
	}
}

// -----------------------------------------------------------------------------

func GCDMod(a int, b int) int {
	if b > a {
		a, b = b, a
	}

	if b == 0 {
		return a
	}

	return GCDMod(b, a%b)
}

// -----------------------------------------------------------------------------

func GCDModIterative(a int, b int) int {
	for {
		if b > a {
			a, b = b, a
		}
		if b == 0 {
			return a
		}
		a, b = b, a%b
	}
}

// -----------------------------------------------------------------------------

type XGCD func(a int, b int) (int, int)

func XGCDRecursive(a int, b int) (int, int) {
	if a > b {
		return XGCDRecursiveOrdered(a, b)
	} else {
		s, t := XGCDRecursiveOrdered(b, a)
		return t, s
	}
}
func XGCDRecursiveOrdered(a int, b int) (int, int) {
	if b == 0 {
		return 1, 0
	}

	s, t := XGCDRecursiveOrdered(b, a%b)
	// d = s*b + t*(a%b)
	// d = s*b + t*(a-b*(a/b))
	// d = s*b + t*a - t*b*(a/b)
	// d = b(s - t*(a/b)) + t*a
	// d = t*a + (s - t*(a/b))b
	return t, s - t*(a/b)
}

// -----------------------------------------------------------------------------
type Matrix2x2 [4]int
type Vector2x1 [2]int

func TransformMatrix(v Vector2x1) Matrix2x2 {
	q := -1 * v[0] / v[1]
	return Matrix2x2([4]int{0, 1, 1, q})
}

func MatrixVecMul(m Matrix2x2, v Vector2x1) Vector2x1 {
	a, b := m[0], m[1]
	c, d := m[2], m[3]

	s, t := v[0], v[1]

	return [2]int{a*s + b*t, c*s + d*t}
}

func MatrixMulMatrix(x Matrix2x2, y Matrix2x2) Matrix2x2 {
	a, b := x[0], x[1]
	c, d := x[2], x[3]

	e, f := y[0], y[1]
	g, h := y[2], y[3]

	return [4]int{
		e*a + g*b,
		f*a + h*b,
		e*c + g*d,
		f*c + h*d,
	}
}

// -----------------------------------------------------------------------------

func XGCDMatrix(a int, b int) (int, int) {
	if b == 0 {
		return 1, 0
	}

	T := TransformMatrix([2]int{a, b})

	s, t := XGCDMatrix(b, a%b)
	v := Vector2x1{s, t}

	r := MatrixVecMul(T, v)

	return r[0], r[1]
}

// -----------------------------------------------------------------------------
func XGCDAcc(a int, b int) (int, int) {
	v := XGCDAccPrime([4]int{1, 0, 0, 1}, [2]int{a, b})
	return v[0], v[1]
}

func XGCDAccPrime(m Matrix2x2, v Vector2x1) Vector2x1 {
	if v[1] == 0 {
		return MatrixVecMul(m, Vector2x1{1, 0})
	}

	t := TransformMatrix(v)

	r := MatrixMulMatrix(m, t)

	a, b := v[0], v[1]
	return XGCDAccPrime(r, [2]int{b, a % b})
}

// -----------------------------------------------------------------------------
func XGCDAcc2(a int, b int) (int, int) {
	v := XGCDAccPrime2([4]int{1, 0, 0, 1}, [2]int{a, b})
	return v[0], v[1]
}

func XGCDAccPrime2(m Matrix2x2, v Vector2x1) Vector2x1 {
	for {
		if v[1] == 0 {
			return MatrixVecMul(m, Vector2x1{1, 0})
		}

		t := TransformMatrix(v)

		r := MatrixMulMatrix(m, t)

		a, b := v[0], v[1]

		m, v = r, [2]int{b, a % b}
	}
}

// -----------------------------------------------------------------------------
func XGCDAcc3(a int, b int) (int, int) {
	m1, m2, m3, m4 := 1, 0, 0, 1

	for b != 0 {
		q := a / b
		m1, m2, m3, m4, a, b = m2, m1-q*m2, m4, m3-q*m4, b, a%b
	}
	return m1, m3
}

// -----------------------------------------------------------------------------
func XGCDWiki(a int, b int) (int, int) {
	s := 0
	oldS := 1
	t := 1
	oldT := 0
	r := b
	oldR := a
	for r != 0 {
		q := oldR / r
		oldR, r = r, oldR-q*r
		oldS, s = s, oldS-q*s
		oldT, t = t, oldT-q*t
	}
	return oldS, oldT
}
