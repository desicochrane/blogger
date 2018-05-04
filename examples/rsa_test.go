package examples

import (
	"testing"
)

func TestFermatPrimalityTest(t *testing.T) {
	primes := []int{
		2, 3, 5, 7, 11,
	}

	for _, p := range primes {
		if !IsProbablePrime(p) {
			t.Errorf("%d is prime", p)
		}
	}

	notPrimes := []int{
		0, 1, 4, 6, 8, 9, 10,
	}

	for _, p := range notPrimes {
		if IsProbablePrime(p) {
			t.Errorf("%d is not prime", p)
		}
	}
	//
	//p, q := RandomPQ(10000, 20000)
	//fmt.Printf("p,q: %d,%d\n", p, q)
	//
	//n := p * q
	//fmt.Printf("n: %d\n", n)
	//
	//phiN := (p - 1) * (q - 1)
	//fmt.Printf("phi(n): %d\n", phiN)
	//
	//e := RandomCoprime(phiN)
	//fmt.Printf("e: %d\n", e)

	//d, y := XGCDAcc3(e, phiN)
	//fmt.Println(XGCDAcc3(526, 527))
	//fmt.Printf("x,y: %d,%d\n", d, y)

	// if d is less than zero, need to correct it
	//for {
	//	if d > 0 {
	//		break
	//	}
	//	d += phiN
	//}
	//fmt.Printf("d: %d\n", d)
	//
	//m := RandomInt(2, n-1)
	//fmt.Printf("m: %d\n", m)
	//
	//encrypted := ModExpIterative(m, e, n)
	//fmt.Printf("m': %d\n", encrypted)
	//
	//decrypted := ModExpIterative(encrypted, d, n)
	//fmt.Printf("m': %d\n", decrypted)
	//for i := 11; i < 100000; i++ {
	//	if IsProbablePrime(i) {
	//		fmt.Println(i)
	//	}
	//}
	//
	//message, _ := strconv.Atoi("desi")
	//fmt.Println(message)
}
