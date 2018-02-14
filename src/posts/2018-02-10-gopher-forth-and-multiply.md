---
title: Gopher Forth and Multiply
date: 10-Feb-2018
---
# Gopher Forth and Multiply

Suppose you are tasked with implementing multiplication for two non-negative integers, that is:

\\[ M(a,n) = a \times n  \text{ where } a,n \in \\{0,1,2,...\\} \\] 

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
         &= ((n+0)a + a \\\\
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
\end{aligned} \]
</div>

Here we can see the number of addition operations is 4. We can do better by making use of the associative law:

$$ a + (b + c) = (a+b)+c \text{  (associative law of addition) }$$

$$ \therefore M(a,4) = ((a+a)+a)+a = (a+a)+(a+a) $$

This means that we only need to compute \\(a+a\\) once and then reuse it. We can generalize this as:

$$ M(a,n) = M(a+a, \frac{n}{2}) $$

Notice that calculating\\(a+a\\)) and \\(\\frac{n}{2}\\) are as simple as a bitwise shift to the left and right respectively if \\(n\\) is an even number. In fact in the case where \\(n\\) is a power of two, that repeatedly dividing in half would eventually reduce the second argument to \\(1\\) which we could use as our inductive base case and we could construct the recursive algorithm as:

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

```go
fmt.PrintLn(Multiply1(2,28)) // 32
```

We can modify our function to perform an additional check for the case where \\(n\\) is odd and doing any additional calculation for that case.

To figure out what extra work is needed, we can observe that for the case where \\(n\\) is odd our bitwise.



