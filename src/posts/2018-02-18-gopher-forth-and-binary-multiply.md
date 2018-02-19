---
title: Gopher Forth and Binary Multiply
date: 18-Feb-2018
---
# Gopher Forth and Binary Multiply

### The problem
Suppose we are tasked with implementing a function `Multiply` which takes two positive integers \\(a\\), \\(n\\) and outputs their product:

$$ \\text{Multiply} : (a,n) \to a \times n \text{ where } a,n \in \\{1,2,3,\\ldots \\} $$

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
package multiply

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
package multiply

// ...

// the native operate to compare our solutions against 
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

// NativeProduct
func BenchmarkNativeProduct(b *testing.B) {
	benchmarkMultiply(b, 17, 28, NativeProduct)
}

// RepeatedAddition
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

$$ \\text{Multiply} : (a,n) \to a \ll log_2{n} \text{ where } n \in \\{2^0,2^1,2^2,\\ldots \\} $$

However this requires performing a logarithm operation to compute \\(\log_2{n}\\), which is not an operation we are allowed to use in our solution. To avoid logarithms we can turn to the properties of exponents and express our function recursively:

\\[ \\begin{aligned}
\\text{Multiply}(a, 2^k) &= a \times 2^k \\\\
                         &= a \times (2^1 \times 2^{k-1}) \\\\
                         &= (a \times 2^1) \times 2^{k-1} \\\\
                         &= \\text{Multiply}(a \times 2^1, 2^{k-1}) \\\\
                         &= \\text{Multiply}(a \times 2^2, 2^{k-2}) \\\\
                         &= \\text{Multiply}(a \times 2^3, 2^{k-3}) \\\\
                         & \ldots \\\\
                         &= \\text{Multiply}(a \times 2^k, 2^0)
\\end{aligned} \\]

Recursively applying the above result means we are doubling \\(a\\) and halving \\(n\\) at each recursive call until eventually reaching the trivial case where \\(n=2^0 = 1\\). Since doubling and halving can be computed by left and right shifts we get the following recursive function definition:

\\[
\\text{Multiply}(a,n) =
\\begin{cases}
  a                             & \\text{if } n = 1 \\\\
  \\text{Multiply}(a\ll1,n\gg1) & \\text{if } n > 1 \\\\
\\end{cases}
\text{ where } n \in \\{2^0,2^1,2^2,\\ldots\\}
\\]


Implementing this in golang, we get:

```go
// multiply.go
package multiply

func MultiplyV1(a int, n int) int {
  if n == 1 {
    return a
  }

  return MultiplyV1(a<<1, n>>1)
}
```

### Accounting for our error

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

Can see that by right bit-wise shift, it will eventually hit zero, which will be our new base case. (otherwise zero would not have worked (todo add failing test))


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

The only time our function adds anything to the result is when it adds the error factor of \\(a\\), and it only does so when the last bit of \\(n\\) is a one. Also every time we shift, our \\(a\\) grows by a factor of two. So then what our algorithm effectively does is to double \\(a\\) every shift and sum all the cases where the bit is one. A way to visualize this is by first representing \\(a\\) in binary, and putting into a table as rows:

(todo: observe the last digit of \\(a\\) as function \\(\\text{lsb}\\) of iteration \((k\\) corresponds to:)

$$ \\text{lsb}(28, k) = 0,1,0,0,0,1 $$

(todo: go through algorithm by hand for case, show that it is faster because it is doing less additions and then discuss that it is really just an "addition chain" and discuss optimal addition chains, and relate to multiplication chains for exponentiation etc.)

(todo: discuss that the number of operations depends on the "population count" of the n)

| \\(2^ka\\) | \\(\\text{lsb}(n, k)\\) | \\(\\text{col}\_{1} \times \\text{col}\_{2}\\) |
|---------- |----------------------------------|----------------|
|  17       | 0                                | 0              |
|  34       | 0                                | 0              |
|  68       | 1                                | 68             |
| 136       | 1                                | 136            |
| 272       | 1                                | 272            |

$$ 17 \times 28 = \sum \\text{col}\_{3} = 68 + 136 + 272= 476  $$

This turns out to be a quick and easy way to perform multiplication by hand, and indeed this method was used as far back as ancient egypt.

On my mac I get the following results: (todo: split into two tables with captions OR use bar chart)

| Method      | \\(a\\)  | \\(n\\)  | performance |
|-------------|----------|----------|-------------|
| `native`    | 2        | 16       | 2.80 ns/op  |
| `multiply0` | 2        | 16       | 53.7 ns/op  |
| `multiply1` | 2        | 16       | 15.1 ns/op  |
| `native`    | 19,998   | 12,234   | 2.46 ns/op  |
| `multiply0` | 19,998   | 12,234   | 4,143 ns/op |
| `multiply1` | 19,998   | 12,234   | 46.3 ns/op  |

Which is almost 2 orders of magnitude better than our first naive algorithm in the case where our inputs are even a little large! However we are still not even close to matching the native multiplication implementation in the language.


Also note that the number of operations are not the only factor, we are also doing a lot of function calls. We can move this in the right direction step by step by first refactoring our code to be tail recursive. 

```go
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
``` 

Now it is in a deferred recursive state, we want to try to move it into a strict tail recurisive state. the benefit of tail recursion it that it easy to visualize what is happening with the plug and chug strategy. To understand how to move from this deferred tail recursion to strict tail recursion, we can "plug and chug" with an example to see how it behaves when \\(a = 17\\) and \\(n = 28\\):

\\[ \\begin{aligned}
M(17,28) &= 0 + M(34,14) \\\\
         &= 0 + 0 + M(68,7) \\\\
         &= 0 + 0 + 68 + M(136,3) \\\\
         &= 0 + 0 + 68 + 136 + M(272,1) \\\\
         &= 0 + 0 + 68 + 136 + 272 + M(544,0) \\\\
         &= 0 + 0 + 68 + 136 + 272 + 0 \\\\
         &= 0 + 0 + 68 + 136 + 272 \\\\
         &= 0 + 0 + 68 + 408 \\\\
         &= 0 + 0 + 476 \\\\
         &= 0 + 476 \\\\
         &= 476 \\\\
\\end{aligned} \\]


To make this strictly tail recursive then it seems like we need to pass forward the amounts being summed as an additional accumulator argument somehow to get a behaviour that is more like this:

\\[ \\begin{aligned}
M(17,28) &= M^\prime(0, 17, 28) \\\\
         &= M^\prime(0,34,14) \\\\
         &= M^\prime(0,68,7) \\\\
         &= M^\prime(68,136,3) \\\\
         &= M^\prime(204,272,1) \\\\
         &= M^\prime(476,544,0) \\\\
         &= 476 \\\\
\\end{aligned} \\]

(todo: optimisation for 0!)
(todo: demonstrate the the invariant is held for each iteration)
The code that achieves that looks like:

```go
func Multiply3(a int, n int) int {
  if n < 0 || a < 0 {
    panic("operand < 0")
  }

  if n > a {
    a, n = n, a
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
  
  a = a<<1
  n = n>>1
  return Multiply3Acc(result, a, n)
}
```

Now we can change the strict recursive tail call into a while true loop directly:

```go
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
```

### Small input \\(a=2, n=16\\):

| Method      | strategy                                  | performance |
|-------------|-------------------------------------------|-------------|
| `native`    | -                                         | 2.80 ns/op  |
| `multiply0` | add/subtract recursive                    | 55.3 ns/op  |
| `multiply1` | double/half recursive                     | 17.0 ns/op  |
| `multiply2` | double/half deferred tail-recursive       | 20.9 ns/op  |
| `multiply3` | double/half strict tail-recursive (accum) | 20.0 ns/op  |
| `multiply4` | double/half while loop (accum)            | 8.83 ns/op  |

### Larger input \\(a=19998, n=12234\\):

| Method      | strategy                                  | performance   |
|-------------|-------------------------------------------|---------------|
| `native`    | -                                         | 2.80 ns/op    |
| `multiply0` | add/subtract recursive                    | 5,3444 ns/op  |
| `multiply1` | double/half recursive                     | 47.9 ns/op    |
| `multiply2` | double/half deferred tail-recursive       | 53.0 ns/op    |
| `multiply3` | double/half strict tail-recursive (accum) | 53.1 ns/op    |
| `multiply4` | double/half while loop (accum)            | 16.7 ns/op    |


One stand out optimization is that it seems we are needlessly doubling a in the case where n has become zero after right shifting. We can remove this by changing our base case to \((n=1\\) and checking early if n is zero, though now that our base case is odd, we have to first.

```go
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
```


Also, another immediate optimization is we could check initially whether \\(n\\) is greater than \\(a\\), then if so we can switch them, then our algorithm will run in \\(O(\log{\text{min}(a,n)})\\).

our final:

```go
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
```


| \\(a\\) | \\(n\\) | `native`    | `Multiply`   |
|---------|---------|-------------|--------------|
| 2      | 16       |  2.92 ns/op | 5.92 ns/op   |
| 19998  | 12234    |  2.90 ns/op | 19.3 ns/op   |
| 2      | 12234    |  2.90 ns/op | 6.22 ns/op   |


(todo: benchmark with power of two, and one minus power of two)
(todo: how big can it handle?)
(todo: consistent benchmarks )
