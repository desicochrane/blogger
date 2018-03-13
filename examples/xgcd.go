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
	if b > a {
		a, b = b, a
	}

	for {
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
	if b > a {
		a, b = b, a
	}
	for {
		if b == 0 {
			return a
		}
		a, b = b, a%b
	}
}
