---
title: Gopher Forth and Binary Multiply
date: 18-Feb-2018
---
# Gopher Forth and Binary Multiply

### The problem
Suppose we are tasked with implementing a function `Multiply` which takes two positive integers \\(a\\), \\(n\\) and outputs their product:

$$ \\text{Multiply}(a,n) = a \times n \text{ where } a,n \in \\{1,2,3,\\ldots \\} $$

We are not allowed to use the native product operator `*` in our solution, however bitwise operations, addition and subtraction are permitted. 

Because we will be comparing several different implementations, it is convenient to first define an interface for our solutions to adhere to:

```go
// multiply.go
package multiply

type Multiply func(a int, n int) int
```


> For the remainder of this article we assume \\(a\\) and \\(n\\) to be strictly positive integers.

### Multiplication as repeated addition
At risk of perpetuating [multiplication as repeated addition](https://www.maa.org/external_archive/devlin/devlin_06_08.html), we can take advantage of the fact that repeated addition for the natural numbers yields the same result as multiplication, that is:
 
 $$ \\text{Multiply}(a,n) = a \times n = \sum_{1}^n a  $$

When we translate this to golang we get:

```go
// multiply.go

type Multiply func(a int, n int) int

func RepeatedAddition(a int, n int) int {
  product := 0
  for i := 1; i <= n; i++ {
    product += a
  }

  return product
}
```

### Benchmarking
To get a feeling of how well our solution performs we can use golang's benchmarking tool to compare our solution with the native product operator:

```go
// multiply.go

func NativeProduct(a int, n int) int {
  return a * n
}
```

```go
// multiply_test.go
package multiply

// helper function (lowercase first letter) to 
// benchmark any implementation of "Multiply"
func benchmarkMultiply(b *testing.B, a int, n int, multiply Multiply) {
  for i := 0; i < b.N; i++ {
    multiply(a, n)
  }
}

func BenchmarkNativeProduct(b *testing.B) {
  benchmarkMultiply(b, 17, 28, NativeProduct)
}

func BenchmarkRepeatedAddition(b *testing.B) {
  benchmarkMultiply(b, 17, 28, RepeatedAddition)
}
```

I ran this on my mac for various inputs for \\( (a,n) \\) and I get:


| \\( (a,n) \\)         | `NativeProduct`    | `RepeatedAddition` |
|-----------------------|--------------------|--------------------|
| \\( (17,28) \\)       | 2.4 ns/op          | 13.0 ns/op         |
| \\( (19998,12234) \\) | 2.4 ns/op          | 4,182 ns/op         |

We can see that our solution performs poorly compared to the native product operator and does not scale at all for larger inputs. I also tried much larger inputs, but I lost my patience well before `RepeatedAddition` finished its computation.

It seems we need to be more clever in our approach.

### Shifting to double and halve
We tried addition, but now let's explore bitwise shifts. We can use the fact that a performing a bitwise left-shift on a number \\(k\\) times is the same as multiplying by it by two \\(k\\) times. That is:

$$ a \times 2^k \equiv a \ll k \text{ where } k \in \\{0,1,2,\\ldots\\} $$

Then for the special case where \\(n\\) is a natural power of 2 we can shift \\(a\\) to the left \\(\log_2{n}\\) times to arrive at a correct solution:

$$ \\text{Multiply}(a,n) = a \ll log_2{n} \text{ where } n \in \\{2^0,2^1,2^2,\\ldots \\} $$

However this requires performing a logarithm operation to compute \\(\log_2{n}\\), which is not an operation we are allowed to use in our solution. To avoid logarithms we can turn to the properties of exponents and express our function recursively:

\\[ \\begin{aligned}
\\text{Multiply}(a, 2^k) &= a \times 2^k \\\\
                         &= a \times (2^1 \times 2^{k-1})            &(x^y \times x^z = x^{y+z}) \\\\
                         &= (a \times 2^1) \times 2^{k-1}            &(\text{associative law of multiplication}) \\\\
                         &= \\text{Multiply}(a \times 2^1, 2^{k-1})  &(\text{definition of Multiply}) \\\\
\\end{aligned} \\]

By repeating this process \\(m\\) times we can observe the recursive pattern:

\\[ \\begin{aligned}
\\text{Multiply}(a, 2^k) &= \\text{Multiply}(a \times 2^1, 2^{k-1}) \\\\
                         &= \\text{Multiply}(a \times 2^2, 2^{k-2}) \\\\
                         &= \\text{Multiply}(a \times 2^3, 2^{k-3}) \\\\
                         & \ldots \\\\
                         &= \\text{Multiply}(a \times 2^m, 2^{k-m})
\\end{aligned} \\]

Notice that eventually we will reach the point where \\(m=k\\) whereby:

\\[ \\text{Multiply}(a \times 2^k, 2^{k-k}) = \\text{Multiply}(a \times 2^k, 1) = a \times 2^k \\]

Thus we can see if we recursively double \\(a\\) and halve \\(n\\) we will eventually reach the case where \\(n=2^0 = 1\\), which we will use as our base case.
 
Since doubling can be computed by a left shift and and halving an even number can be computed by a right shift (and since \\(n=2^k\\) is an even number) we can define the following recursive function:

\\[
\\text{Multiply}(a,n) =
\\begin{cases}
  a                             & \\text{if } n = 1 \\\\
  \\text{Multiply}(a\ll1,n\gg1) & \\text{if } n > 1 \\\\
\\end{cases}
\text{ where } n \in \\{2^0,2^1,2^2,\\ldots\\}
\\]


In golang this is as simple as:

```go
// multiply.go

func RecursiveDoubleHalf(a int, n int) int {
  if n == 1 {
    return a
  }

  return RecursiveDoubleHalf(a<<1, n>>1)
}
```


### Accounting for our error
> todo

### Understanding what's happening
> todo

### Tail recursion
> todo

### Fine Tuning
> todo

### Summary
> todo
