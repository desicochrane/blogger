---
title: [wip] GCD
date: 01-Jan-2018
---

### Extended

Suppose someone computed gcd of a and b. how can they convince you they have actually found the gcd?

they can present s and t. you can verify that sa + tb = d, and verify that d divides a and b. (todo: see if this is faster than just computing d yourself!)

Can prove why, that if \\( d = sa + tb \\) and \\(d \mid a \\) and \\(d \mid b) then d must be the gcd (algebraically and intuitively using shapes). 

> Lemma: if d|a and d|b and d=ax+by then d=gcd(a,b)
since d does divide a and b, we at least know it is a common divisor. since no divisor can be greater than the greatest common divisor (tautologies yo!) d must either be less than the gcd or equal to it: then d <= gcd(a,b)
> need to prove that any divisor of a and b is also a divisor of any linear combination of a and b.
> need to then show that d is a linear combination of a and b.
> then can show that the gcd must divide d so it cannot be bigger than d, which implies that gcd(a,b) <= d
> therefore, d is EQUAL to gcd. 

But what if there are NO s and t? But how to prove that such an s and t MUST exist? <bezouts lemma yo...>

can "prove" mechanically by solving for gcd and working backwards, and then showing that process would work for arbitrary a,b?
can "prove" by assuming true showing gcd(a,b) = gcd(b,a%b) therefore d = sa + tb = pb + q(a%b), then prove by induction?


```go
func XGCD(a int, b int) (int, int) {
  if b == 0 {
    return 1, 0
  }

  s, t := XGCD(b, a%b)
  return t, s - t*(a/b)
}
```

```go
func WikipediaXGCD(a int, b int) (int, int) {
  s, oldS, t, oldT, r, oldR := 0, 1, 1, 0, b, a
  for r != 0 {
    q := oldR / r
    oldR, r = r, oldR-q*r
    oldS, s = s, oldS-q*s
    oldT, t = t, oldT-q*t
  }
  return oldS, oldT
}
```

\\(d = \\text{gcd}(a,b)\\) is greatest common divisor of two integers \\(a\\) and \\(b\\)

$$ d = \\text{gcd}(a,b) $$

$$ d = \\text{gcd}(24,16) = 8$$

Given two integers a and b find the integers s and t such that \\( d = s \times a + t \times b \\)

$$ s, t = \text{xgcd}(a,b) $$

$$ s, t = \text{xgcd}(24,16) = 1, -1 $$

$$ 8 = 1 \times 24 + -1 \times 16 $$

```go
func XGCD(a int, b int) (int, int) {
	if b == 0 {
		return 1, 0
	}

	s, t := XGCD(b, a%b)
	return t, s - t*(a/b)
}
```

\\[ \\begin{aligned}
T\Big(\lfloor \frac{a}{b} \rfloor, \begin{bmatrix} s \\\ t \end{bmatrix}\Big) &= \begin{bmatrix} t \\\ s - t \times \lfloor \frac{a}{b} \rfloor \end{bmatrix} \\\\
&= \begin{bmatrix} 0 & 1 \\\ 1 & \lfloor -\frac{a}{b} \rfloor \end{bmatrix} \times \begin{bmatrix} s \\\ t \end{bmatrix} \\\\
&= \mathbf{A}\_{(a,b)} \times \begin{bmatrix} s \\\ t \end{bmatrix} \text{ where } \mathbf{A}\_{(a,b)} = \begin{bmatrix} 0 & 1 \\\ 1 & \lfloor -\frac{a}{b} \rfloor \end{bmatrix}
\\end{aligned} \\]


\\[ \\begin{aligned}
\\text{xgcd}\Big(\begin{bmatrix} 24 \\\ 16 \end{bmatrix}\Big) &= \mathbf{A}\_{(24,16)} \times \\text{xgcd}\Big(\begin{bmatrix} 16 \\\ 8 \end{bmatrix}\Big) \\\\
&= \mathbf{A}\_{(24,16)} \times \mathbf{A}\_{(16,8)} \times \\text{xgcd}\Big(\begin{bmatrix} 8 \\\ 0 \end{bmatrix}\Big) \\\\
&= \mathbf{A}\_{(24,16)} \times \mathbf{A}\_{(16,8)} \times \begin{bmatrix} 1 \\\ 0 \end{bmatrix} \\\\
&= \mathbf{A}\_{(24,16)} \times \begin{bmatrix} 0 \\\ 1 \end{bmatrix} \\\\
&= \begin{bmatrix} 1 \\\ -1 \end{bmatrix} \\\\
\\end{aligned} \\]

\\[ \\begin{aligned}
\\text{xgcd}\Big(\begin{bmatrix} 24 \\\ 16 \end{bmatrix}\Big) &= \\text{xgcd}^{\prime}\Big(\mathbf{A}\_{(24,16)}, \begin{bmatrix} 16 \\\ 8 \end{bmatrix}\Big) \\\\
&= \\text{xgcd}^{\prime}\Big(\mathbf{A}\_{(24,16)} \times \mathbf{A}\_{(16,8)}, \begin{bmatrix} 8 \\\ 0 \end{bmatrix}\Big) \\\\
&= \begin{bmatrix} 1 & -2 \\\ -1 & 3 \end{bmatrix} \times \begin{bmatrix} 1 \\\ 0 \end{bmatrix} \\\\
&= \begin{bmatrix} 1 \\\ -1 \end{bmatrix}
\\end{aligned} \\]

\\[
\text{xgcd}\Big(\begin{bmatrix} a \\\ b \end{bmatrix}\Big) = \\text{xgcd}^{\prime}\Big(I, \begin{bmatrix} a \\\ b \end{bmatrix}\Big)
\\]


\\[
\\text{xgcd}^{\prime}\Big(\mathbf{A}, \begin{bmatrix} a \\\ b \end{bmatrix} \Big) =
\\begin{cases}
  \mathbf{A} \times \begin{bmatrix} a \\\ b \end{bmatrix}    & \\text{if } b = 0 \\\\
  \\text{xgcd}^{\prime}\Big(\mathbf{A} \times \mathbf{A}\_{(a,b)}, \begin{bmatrix} b \\\ a \bmod b \end{bmatrix} \Big) & \\text{if } b \ne 0 \\\\
\\end{cases}
\\]

```go
func XGCD(a int, b int) (int, int) {
  m1, m2, m3, m4 := 1, 0, 0, 1

  for b != 0 {
    q := a / b
    m1, m2, m3, m4, a, b = m2, m1-q*m2, m4, m3-q*m4, b, a%b
  }
  return m1, m3
}
```


can check out https://www.coursera.org/learn/mathematical-foundations-cryptography/lecture/8XvAY/extended-euclidean-algorithm for an xgcd that optimises for case where gcd(a,b) = 1 for inverse checking
