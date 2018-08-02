package examples

import "fmt"

type Choose func(n, k int) int

// -----------------------------------------------------------------------------
func ChooseFactorialNaive(n, k int) int {
	// assumes n >= k >= 0
	if n < k || k < 0 {
		panic("invalid n,k")
	}

	return factorial(n) / (factorial(k) * factorial(n-k))
}

// -----------------------------------------------------------------------------
func ChooseKFacNaive(n, k int) int {
	// assumes n >= k >= 0
	if n < k || k < 0 {
		panic("invalid n,k")
	}

	return kfactorial(n, k) / (factorial(k))
}

// -----------------------------------------------------------------------------
func ChooseKFacCheckLarger(n, k int) int {
	// assumes n >= k >= 0
	if n < k || k < 0 {
		panic("invalid n,k")
	}

	s := n - k
	if s > k {
		return kfactorial(n, k) / factorial(k)
	} else {
		return kfactorial(n, s) / factorial(s)
	}
}

// -----------------------------------------------------------------------------
func ChooseRec(n, k int) int {
	// assumes n >= k >= 0
	if n < k || k < 0 {
		panic(fmt.Sprintf("invalid (n,k): (%d,%d)", n, k))
	}

	if k == 0 || n == k {
		return 1
	}

	return ChooseRec(n-1, k) + ChooseRec(n-1, k-1)
}

// -----------------------------------------------------------------------------
func ChooseRecCached(n, k int) int {
	return chooseRecCached(n, k, make(map[int]map[int]int))
}

func chooseRecCached(n, k int, cache map[int]map[int]int) int {
	if resultN, okN := cache[n]; okN {
		if resultK, okK := resultN[k]; okK {
			return resultK
		}
	} else {
		cache[n] = make(map[int]int)
	}

	if k == 0 || n == k {
		return 1
	}

	a, b := chooseRecCached(n-1, k, cache), chooseRecCached(n-1, k-1, cache)
	cache[n-1][k] = a
	cache[n-1][k-1] = b

	return a + b
}

// -----------------------------------------------------------------------------
func ChooseRec2(n, k int) int {
	// assumes n >= k >= 0
	if n < k || k < 0 {
		panic(fmt.Sprintf("invalid (n,k): (%d,%d)", n, k))
	}

	if k == 0 || n == k {
		return 1
	}

	return ChooseRec(n-1, k) + ChooseRec(n-1, k-1)
}

// -----------------------------------------------------------------------------
func ChooseBallsBoxes(n, k int) int {
	if n < k || k < 0 {
		panic(fmt.Sprintf("invalid (n,k): (%d,%d)", n, k))
	}
	boxes := k
	balls := n + 1 - k
	return chooseBallsBoxes(balls, boxes, map[int]map[int]int{})
}

func chooseBallsBoxes(ingredients, slots int, cache map[int]map[int]int) int {
	if slots == 1 {
		return ingredients
	}
	if ingredients == 1 {
		return 1
	}

	if c, ok := cache[ingredients]; ok {
		if d, ok := c[slots]; ok {
			return d
		}
	} else {
		cache[ingredients] = make(map[int]int)
	}

	sum := 1
	for i := 1; i <= slots; i++ {
		sum += chooseBallsBoxes(ingredients-1, i, cache)
	}

	cache[ingredients][slots] = sum

	return sum
}

// -----------------------------------------------------------------------------
func factorial(n int) int {
	if n < 0 {
		panic(fmt.Sprintf("invalid n: %d", n))
	}

	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}

	return result
}

func kfactorial(n, k int) int {
	if n < k || k < 0 {
		panic(fmt.Sprintf("invalid (n,k): (%d,%d)", n, k))
	}

	result := 1
	for i := 0; i < k; i++ {
		result *= n - i
	}
	return result
}

// -----------------------------------------------------------------------------
