---
title: Gopher Forth and Multiply
date: 14-Feb-2018
---
# Gopher Forth and Multiply

Suppose you are tasked with implementing multiplication for two non-negative integers, that is:

\\[ M(a,n) = a \times n  \text{ where } a,n \in \\{0,1,2,...\\} \\] 

(todo: explain simple case where multiplying by powers of 2 as a starting point)
(todo: discuss legal operations)

At risk of perpetuating [multiplication as repeated addition](https://www.maa.org/external_archive/devlin/devlin_06_08.html) we can use the fact that repeated addition for the natural numbers yields the same result as multiplication. Then we can express multiplication as recursive addition, by first defining the base case for \\(n = 0\\) and then for the arbitrary case where \\(n > 0\\):

\\[
M(a,n) = 
\\begin{cases}
  0          & \\text{if } n = 0 \\\\
  (n-1)a + a & \\text{if } n > 0
\\end{cases} \\]

Which although might seem obviously correct, we can prove it via induction:

\\[ \text{base case: }  M(a,0) = a \\times 0 = 0 \\]

\\[ \text{Suppose true for arbitrary n:} \\]

\\[ \\begin{aligned}
M(a,n+1) &= ((n+1)-1)a + a \\\\
         &= ((n-0)a + a \\\\
\\end{aligned} \\]

```go
// multiply.go

func Multiply0(a int, n int) int {
  if n < 0 || a < 0 {
    panic("operand < 0")
  }

  if n == 0 {
    return 0
  }

  return Multiply0(a, n-1) + a
}
```  

To get an idea of how well we might be doing, we can benchmark our implementation and compare it to the native mulitplication operator:

```go
// multiply.go

type Multiply func(a int, n int) int

func MultiplyNative(a int, n int) int {
  if n < 0 || a < 0 {
    panic("operand < 0")
  }

  return a * n
}
```

Now we can create generic test to benchmark for small and large operands:

```go
// multiply_test.go

func benchmarkMultiply(b *testing.B, a int, n int, multiply Multiply) {
  for i := 0; i < b.N; i++ {
    multiply(a, n)
  }
}
```

On my mac I get the following results:

| Method      | \\(a\\)  | \\(n\\)  | performance |
|-------------|----------|----------|-------------|
| `native`    | 2        | 16       | 2.80 ns/op  |
| `multiply0` | 2        | 16       | 9.00 ns/op  |
| `native`    | 19,998   | 12,234   | 2.46 ns/op  |
| `multiply0` | 19,998   | 12,234   | 4,143 ns/op  |

Observe the number of operations for our naive first implementation for the large numbers. So let's turn back to our symbolic representation of the problem to see how it behaves for case \\(n = 4\\):

<div katex>
\[ \begin{aligned}
M(a,4) &= (4-1)a+a \\
       &= ((3-1)a+a)+a \\
       &= (((2-1)a+a)+a)+a \\
       &= ((((1-1)a+a)+a)+a)+a \\
       &= (((0+a)+a)+a)+a \\
       &= ((a+a)+a)+a \\
       &= (a+a)+(a+a) \text{ (associative law of addition)} \\ 
       &= 2(a+a) \\
\end{aligned} \]
</div>

This means that we only need to compute \\(a+a\\) once and then reuse it. We can generalize this as:

$$ M(a,n) = M(a+a, \frac{n}{2}) $$

Notice that calculating \\(a+a\\)) and \\(\\frac{n}{2}\\) are as simple as a bitwise shift to the left and right respectively if \\(n\\) is an even number. In fact in the case where \\(n\\) is a power of two, that repeatedly dividing in half would eventually reduce the second argument to \\(1\\) which we could use as our inductive base case and we could construct the recursive algorithm as:

```go
// multiply.go

func Multiply1(a int, n int) int {
  if n < 0 || a < 0 {
    panic("operand < 0")
  }

  if n == 1 {
    return a
  }

  return Multiply1(a<<1, n>>1)
}
```

This is a big improvement in terms of operations, because in each successive iteration \\(n\\) decreased by a factor of 2 and we are only performing constant time operations in our function we have gone from a worst case scenario of \\(O(n)\\) to \\(O(\log{n})\\). 

However our algorithm assumes that we can divide \\(n\\) by a factor of two cleanly at every iteration, that is we assume that \\(n\\)) is a power of 2. If we invoked our new method with \\(n=28\\) for example we will get an incorrect result:

(todo: "// Run this program in the Go Playground." link)
```go
fmt.PrintLn(Multiply1(2,28)) // 32
```

We can modify our function to perform an additional check for the case where \\(n\\) is odd and doing any additional calculation for that case.

To figure out what extra work is needed, we can observe that for the case where \\(n\\) is odd our bitwise operation effectively ignores the last bit, thus:

\\[
n \gg 1 = 
\\begin{cases}
  \frac{n}{2} & \\text{if } n \text{ is even} \\\\
  \frac{n-1}{2} & \\text{if } n \text{ is odd}
\\end{cases} \\]

So we can calculate the error  in the case where \\(n\\) is odd as:

\\[ \\begin{aligned}
\\text{error} &= \text{correct} - \\text{result} \\\\
              &= M(a,n) - M(a+a,n \gg 1) \\\\
              &= M(a,n) - M(2a,\frac{n-1}{2}) \\\\
              &= an - 2a\big(\frac{n-1}{2}\big) \\\\
              &= an - a(n-1) \\\\
              &= an - an+a \\\\
              &= a
\\end{aligned} \\]

So we can see that all we need to do to adjust our algorithm is to add the error \\(a\\) in the case where \\(n\\) is negative. We can determine if \\(n\\) is negative by checking if its least significant bit is a \\(1\\) via a bitmask (todo: illustrate the bitmask), and result in the following corrected implementation.

Can see that by right bit-wise shift, it will eventually hit zero, which will be our new base case. (otherwise zero wouldn't have worked (todo add failing test))


```go
// multiply.go

func Multiply1(a int, n int) int {
  if n < 0 || a < 0 {
    panic("operand < 0")
  }

  if n == 0 {
    return 0
  }

  result := Multiply1(a<<1, n>>1)

  if n&0x1 == 1 {
    // account for error of "a" 
    // for case where "n" is odd
    return result + a 
  }

  return result
}
```

So what is actually happening?

`010110`

`001011`

`000101`

`000010`

`000001`

`000000`

The only time our function adds anything to the result is when it adds the error factor of \\(a\\), and it only does so when the last bit is a one. Also every time we shift, our \\(a\\) grows by a factor of two. So then what our algorithm effectively does is to double \\(a\\) every shift and sum all the cases where the bit is one. A way to visualize this is by first representing \\(a\\) in binary, and putting into a table as rows:

(todo: observe the last digit of \\(a\\) as function \\(\\text{FinalDigit}\\) of iteration \((k\\) corresponds to:)

$$ \\text{FinalDigit}(28, k) = 0,1,0,0,0,1 $$

(todo: go through algorithm by hand for case)

| \\(2^kn\\) | \\(\\text{FinalDigit}(a, k)\\) | \\(\\text{col}\_{1} \times \\text{col}\_{2}\\) |
|---------- |----------------------------------|----------------|
|   1       | 0                                | 0              |
|  28       | 1                                | 28             |
|  56       | 0                                | 0              |
| 112       | 0                                | 0              |
| 224       | 0                                | 0              |
| 448       | 1                                | 448            |

$$ \sum(\\text{col}\_{3}) = 28 + 448 = 476 = 17 \times 28 $$

This turns out to be a quick and easy way to perform multiplication by hand, and indeed this method was used as far back as ancient egypt.


On my mac I get the following results: (todo: split into two tables with captions OR use bar chart)

| Method      | \\(a\\)  | \\(n\\)  | performance |
|-------------|----------|----------|-------------|
| `native`    | 2        | 16       | 2.80 ns/op  |
| `multiply0` | 2        | 16       | 53.7 ns/op  |
| `multiply1` | 2        | 16       | 15.1 ns/op  |
| `native`    | 19,998   | 12,234   | 2.46 ns/op  |
| `multiply0` | 19,998   | 12,234   | 4,143 ns/op  |
| `multiply1` | 19,998   | 12,234   | 46.3 ns/op  |

Which is almost 2 orders of magnitude better than our first naive algorithm in the case where our inputs are even a little large! However we are still not even close to matching the native multiplication implementation in the language.




