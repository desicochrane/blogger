---
title: Gopher Forth and Multiply
date: 22-Feb-2018
---
# Gopher Forth and Multiply

### Introduction
In this article we will use the "Binary Multiplication Algorithm" as a toy motivating example to play with recursion. We are going to approach the algorithm by first examining the problem it solves then building up to the final solution bit by bit. In this way I hope you gain an intuition for how the algorithm works and at the same time pick up some useful tactics for how to refactor recursive functions to improve performance.

We will be using golang for our implementation because I like it and because it has built in benchmarking for us to measure our performance - but the code examples should be straightforward enough to understand even if you've never seen any golang before. 

> In the spirit of learning and thinking more symbolically about the code we write I have not shied away from mathematical notation. But don't let that scare you off! For me at least, the satisfaction of reaching that "Aha!" moment is totally worth the time spent studying and staring down a chunk of intimidating math that I didn't immediately understand.

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

func RepeatedAddition(a int, n int) int {
  product := 0
  for i := 1; i <= n; i++ {
    product += a
  }

  return product
}
```

### Benchmarking
To get a feeling for how well our solution performs we can use golang's benchmarking tool to compare our solution with the native product operator:


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

> Don't worry if the benchmarking code is unfamiliar, all we are doing is getting an idea of how fast our code is by running it a bunch of times. You won't need to understand it to follow the rest of the article.

I ran this on my mac for various inputs for \\( (a,n) \\) and I get:


| \\( (a,n) \\)         | `NativeProduct`    | `RepeatedAddition` |
|-----------------------|--------------------|--------------------|
| \\( (17,28) \\)       | 2.4 ns/op          | 13.0 ns/op         |
| \\( (19998,12234) \\) | 2.4 ns/op          | 4,182 ns/op         |

We can see that our solution performs poorly compared to the native product operator and does not scale at all for larger inputs.

It seems we need to be more clever in our approach.

### Doubling and halving
We tried addition, but now let's explore bitwise shifts. We can use the fact that a performing a bitwise left-shift on a number \\(k\\) times is the same as multiplying by it by two, \\(k\\) times. That is:

$$ a \times 2^k \equiv a \ll k \text{ where } k \in \\{0,1,2,\\ldots\\} $$

> Note we use the notation \\(\ll\\) and \\(\gg\\) as left and right bitwise-shift operators

Then for the special case where \\(n\\) is a natural power of 2 we can shift \\(a\\) to the left \\(\log_2{n}\\) times to arrive at a correct solution:

$$ \\text{Multiply}(a,n) = a \ll log_2{n} \text{ where } n \in \\{2^0,2^1,2^2,\\ldots \\} $$

However this requires performing a logarithm operation to compute \\(\log_2{n}\\), which is not an operation we are allowed to use in our solution. To avoid logarithms we can turn to the properties of exponents and express our function recursively:

\\[ \\begin{aligned}
\\text{Multiply}(a, 2^k) &= a \times 2^k \\\\
                         &= a \times (2^1 \times 2^{k-1})            &(x^y \times x^z = x^{y+z}) \\\\
                         &= (a \times 2^1) \times 2^{k-1}            &(\text{associative law of multiplication}) \\\\
                         &= \\text{Multiply}(a \times 2^1, 2^{k-1})  &(\text{definition of Multiply}) \\\\
\\end{aligned} \\]

In other words we are doubling our first argument and halving our second argument (notice that \\(2^{k-1} = 2^k / 2 \\)
). By repeating this process \\(m\\) times we can observe the recursive pattern:

\\[ \\begin{aligned}
\\text{Multiply}(a, 2^k) &= \\text{Multiply}(a \times 2^1, 2^{k-1}) \\\\
                         &= \\text{Multiply}(a \times 2^2, 2^{k-2}) \\\\
                         &= \\text{Multiply}(a \times 2^3, 2^{k-3}) \\\\
                         & \ldots \\\\
                         &= \\text{Multiply}(a \times 2^m, 2^{k-m})
\\end{aligned} \\]

Notice that eventually we will reach the point where \\(m=k\\) whereby:

\\[ \\text{Multiply}(a \times 2^k, 2^{k-k}) = \\text{Multiply}(a \times 2^k, 1) = a \times 2^k \\]

Thus we can see that we are guaranteed our second argument will eventually reach \\(2^0 = 1\\), which we can use as our base case. For the non-base case, we will recursively call the function with double the first argument \\(a\\) and half the second argument \\(n\\), which we can perform by left and right-shifts respectively:

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

### Fixing for arbitrary \\(n\\)

If we compare our new approach to our first algorithm, we can see that we've made a big improvement in terms of operations performed. For `RepeatedAddition` we performed \\(n\\) additions, whereas `RecursiveDoubleHalf` performs \\(2\log_2{n}\\) shift operations, so we have gone from a worst case scenario of \\(O(n)\\) to \\(O(\log{n})\\). 

However our new algorithm depends on \\(n\\) being a power of 2, if we try \\(a = 17 \\) and \\(n = 28\\) we get:

```go
fmt.PrintLn(RecursiveDoubleHalf(17,28)) // 272
```

Which is incorrect.

To see why, notice that right-shifting an odd number means we lose any information that was in the least significant (rightmost) bit as it drops off the end after the shift:
 
\\[ \\begin{aligned}
17\_{\text{base} 10} \gg 1 &\to 10001\_{\text{base} 2} \gg 1 \\\\
         &\to 01000\_{\text{base} 2} \\\\
         &\to 8\_{\text{base} 10}
\\end{aligned} \\]

Thus we can define a right-shift as:

\\[
n \gg 1 = 
\\begin{cases}
  \frac{n}{2} & \\text{if } n \text{ is even} \\\\
  \frac{n-1}{2} & \\text{if } n \text{ is odd}
\\end{cases} \\]

Because our function depends on right-shifting being equivalent to halving, we require that \\(n\\) is even for every recursive call - which is only the case where \\(n\\) is a natural power of 2. If \\(n\\) is *not* even for even a single recursive call we will get an incorrect answer, because we are actually computing:

\\[ \\begin{aligned}
\\text{Multiply}(a,n) &= \\text{Multiply}(a\ll1,n\gg1) \\\\
                      &= \\text{Multiply}(2a,\frac{n-1}{2}) \\\\
\\end{aligned} \\]

We can calculate the correction we need to apply algebraically as:

\\[ \\begin{aligned}
\\text{correction} &= \text{correct} - \\text{computed} \\\\
              &= (a \times n)   - 2a\big( \frac{n-1}{2} \big) \\\\
              &= an - a(n-1) \\\\
              &= an - an+a \\\\
              &= a
\\end{aligned} \\]

So in the case where \\(n\\) is odd our computed solution will be \\(a\\) less than the correct answer, fixing our algorithm then means we need to add \\(a\\) to our result when \\(n\\) is odd. 

To determine if \\(n\\) is odd we just need to check if its least significant bit is a 1, which we can do by performing a logical `and` with \\(n\\) and 1 and seeing if the result is 1:

\\[ \\begin{aligned}
17\_{\text{base} 10} \\small{\\text{ AND }} 1\_{\text{base} 10} &\to 10001\_{\text{base} 2} \\small{\\text{ AND }} 00001\_{\text{base} 2} \\\\
                                     &\to 00001\_{\text{base} 2} \\\\
                    &\to 1\_{\text{base} 10}
\\end{aligned} \\]

Now we can fix our function:

```go
// multiply.go

func RecursiveDoubleHalf(a int, n int) int {
  if n == 1 {
    return a
  }

  result := RecursiveDoubleHalf(a<<1, n>>1)

  // if n is odd we need to add "a"
  if n&1 == 1 {
    return result + a
  }

  return result
}
```

When I run the benchmarking I get:

| \\( (a,n) \\)         | `NativeProduct`    | `RepeatedAddition` | `RecursiveDoubleHalf` |
|-----------------------|--------------------|--------------------|-----------------------|
| \\( (17,28) \\)       | 2.4 ns/op          | 13.0 ns/op         | 14.1 ns/op            |
| \\( (19998,12234) \\) | 2.4 ns/op          | 4,182 ns/op        | 45.6 ns/op            |

So we can see that while our performance has not changed much for small inputs, our new algorithm seems to scale much better for larger inputs.

### Understanding what's happening
We can observe that at each recursive step we are performing the following:

1. If the last bit of \\(n\\) is a 1, add \\(a\\) to our result
2. right-shift \\(n\\), left-shift \\(a\\)

And it repeats this process until \\(n\\) has been completely right-shifted, which we say takes \\(m\\) iterations where \\(m\\) is the number of bits in \\(n\\), i.e. \\(m = \lfloor \log_2{n} \rfloor\\).

Notice that the only time we update our result is when the least significant bit of \\(n\\) is a 1. Since our \\(n\\) is shifted right at every iteration, this is equivalent to a left-to-right bitwise scan of \\(n\\), that is for each iteration \\(k\\) we are checking if the \\(k^{\text{th}}\\) significant bit of \\(n\\), \\(\text{ksb}(n)\\).

So at each iteration \\(k\\) we have for \\(n=28\\):

$$ \\text{ksb}(n) = 0,0,1,1,1 $$

Which is the left-to-right binary representation of 28. 

To better understand then how our function works we can construct the following table to compute the solution by hand for \\(a=17\\) and \\(n=28\\):

| k | \\(\\text{col}\_{1} = a \times 2^k\\) | \\(\\text{col}\_{2} = \\text{ksb}(n)\\) | \\(\\text{col}\_{3} = \\text{col}\_{1} \times \\text{col}\_{2}\\) |
|---|------------------- |----------------------|----------------|
| 0 |  17                | 0                    | 0              |
| 1 |  34                | 0                    | 0              |
| 2 |  68                | 1                    | 68             |
| 3 | 136                | 1                    | 136            |
| 4 | 272                | 1                    | 272            |

$$ 17 \times 28 = \sum \\text{col}\_{3} = 68 + 136 + 272= 476  $$

> Note that \\( \text{col}\_{3} \\) can be calculated as: 
    \\( \text{col}\_{3} = \\begin{cases}
                            0               & \\text{if } \text{col}\_{2} = 0 \\\\
                            \text{col}\_{1} & \\text{if } \text{col}\_{2} = 1 \\\\
                          \\end{cases} \\)

This turns out to be a pretty quick way to perform multiplication by hand, and indeed this method was used as far back as ancient egypt.

If something about this algorithm seems vaguely familiar to you, it's because this is actually what is happening in the grade-school multiplication technique you likely learned when you were a child - just that this uses a base 2 numbering system. To see why we can try it by hand:

```txt
      10001 =  17
    x 11100 =  28
    -------  
      00000  
     00000   
    10001    
   10001    
+ 10001      
-----------  
  111011100 = 476
```

Which follows exactly the same pattern as our function - it scans through the each digit of \\(n\\) left-to-right, in the case where that digit is a zero we add nothing, in the case where that digit is a one we adds a left-shifted \\(a\\).

### Can we do better? Tail recursion
At this point our algorithm performs a number of operations, namely:

- \\(\lfloor \log_2{n} \rfloor\\) left-shifts; plus
- \\(\lfloor \log_2{n} \rfloor\\) right-shifts; plus 
- an addition operation every time a \\(1\\) is the least-significant-bit, which is just the number of 1's in the binary representation of \\(n\\), also known as the *population count* of \\(n\\) or \\(\\text{pop}(n) \\).

$$ \\text{operations} = 2\lfloor \log_2{n} \rfloor + \\text{pop}(n) $$

Which is still of the order \\(O(\log{n})\\).

However on top of these operations we are also performing \\(\lfloor \log_2{n} \rfloor\\) recursive function calls, which come with a performance cost. To get rid of these function calls we need to first refactor our solution to be tail-recursive.

```go
func TailRecursiveDoubleHalf(a int, n int) int {
  if n == 1 {
    return a
  }

  correction := 0
  if n&1 == 1 {
    correction = a
  }

  return correction + TailRecursiveDoubleHalf(a<<1,n>>1)
}
``` 

Notice that our recursive call is now the final operation in the function. We say that our function is in a *deferred* tail-recursive state. This is because after the recursive call returns we have some deferred work to do, namely we need to add the `correction` term.

What we want however is for our function to be in a *strict* tail-recursive state, that is, we want there to be no additional work to be done after the recursive step.

To see why, we can "plug and chug" our way through an example using \\(a=17\\) and \\(n=28\\) as our function currently is:

\\[ \\begin{aligned}
\\text{Multiply}(17,28) &= 0 + \\text{Multiply}(34,14) \\\\
                        &= 0 + 0 + \\text{Multiply}(68,7) \\\\
                        &= 0 + 0 + 68 + \\text{Multiply}(136,3) \\\\
                        &= 0 + 0 + 68 + 136 + \\text{Multiply}(272,1) \\\\
                        &= 0 + 0 + 68 + 136 + 272 \\\\
                        &= 0 + 0 + 68 + 408 \\\\
                        &= 0 + 0 + 476 \\\\
                        &= 0 + 476 \\\\
                        &= 476 \\\\
\\end{aligned} \\]

Here we can see clearly the addition operations that are deferred, separate from our recursive call. For our function to be strictly tail-recursive, we need to find a way to accumulate those deferred operations into our recursive function, that is we want our function to somehow behave more like this:

\\[ \\begin{aligned}
\\text{Multiply}(17,28) &= \\text{Multiply}^\prime(0,17,28) \\\\
                        &= \\text{Multiply}^\prime(0,34,14) \\\\
                        &= \\text{Multiply}^\prime(0,68,7) \\\\
                        &= \\text{Multiply}^\prime(68,136,3) \\\\
                        &= \\text{Multiply}^\prime(204,272,1) \\\\
                        &= 476 \\\\
\\end{aligned} \\]

This is our desired "plug and chug" output. Notice it uses a new function \\(\\text{Multiply}^\prime\\) which takes 3 arguments, the first of which is acts as an accumulator, capturing those deferred additions in our current implementation. This additional helper function is known as an *accumulator* function. 

Observe that at each recursive call of \\(\\text{Multiply}^\prime\\) the invariant \\(\text{acc} + a \times n \\) holds, which is the value of the final output. Our new accumulator function can then be defined as:
 $$ \\text{Multiply}^\prime(\text{acc},a,n) = \text{acc} + a \times n $$

By inspecting the desired "plug and chug" we can see how our accumulator ought to behave and thus these become our instructions as we implement it. At each step it checks if \\(n\\) is odd in which case it increases the accumulator argument by \\(a\\), and in the case where \\(n=1\\) it returns the accumulator as the solution:

```go
func StrictTailRecursiveDoubleHalf(a int, n int) int {
  return StrictTailRecursiveDoubleHalfAcc(0, a, n)
}

func StrictTailRecursiveDoubleHalfAcc(acc int, a int, n int) int {
  if n&1 == 1 {
    acc += a

    if n == 1 {
      return acc
    }
  }

  a = a << 1
  n = n >> 1
  return StrictTailRecursiveDoubleHalfAcc(acc, a, n)
}
```

Now our solution is **strictly tail-recursive** - its has no deferred operations after its recursive call and it simply calls itself with the same arguments. This is effectively the same as a `GOTO` statement, telling the function to start at the top again.

When I benchmark this new implementation against our previous, we see a performance *decrease*:

| \\( (a,n) \\)         | `RecursiveDoubleHalf` | `StrictTailRecursiveDoubleHalf` |
|-----------------------|-----------------------|---------------------------------|
| \\( (17,28) \\)       | 14.1 ns/op            | 17.6 ns/op                      |
| \\( (19998,12234) \\) | 45.6 ns/op            | 50.7 ns/op                      |

This performance hit is due to the golang compiler not providing automatic tail-recursion optimization (at least at time of writing). Fortunately, when our code is in a strict-tail-recursive state, optimization is as straight-forward as replacing the recursive call with a `while(true)` loop or a `GOTO` statement - the rest of the code can be left as-is!

In golang we emulate a `while(true)` loop by using a `for` loop without any condition. Also once we have replaced the recursive call with a loop, we no longer need the accumulator function. The final result comes out as:

```go
func IterativeDoubleHalf(a int, n int) int {
  acc := 0

  for {
    if n&1 == 1 {
      acc += a

      if n == 1 {
        return acc
      }
    }

    a = a << 1
    n = n >> 1
  }
}
```

Which yields a significant performance boost:

| \\( (a,n) \\)         | `RecursiveDoubleHalf` | `StrictTailRecursiveDoubleHalf` | `IterativeDoubleHalf` |
|-----------------------|-----------------------|---------------------------------|-----------------------|
| \\( (17,28) \\)       | 14.1 ns/op            | 17.6 ns/op                      | 9.18 ns/op            |
| \\( (19998,12234) \\) | 45.6 ns/op            | 50.7 ns/op                      | 21.9 ns/op            |

### Summary
So that's that! In the end our solution is still much slower than the native operator, but we have come a long way. We've seen how to reason about and make solid claims about our recursive code. We have spent some time thinking about our algorithm from different perspectives to get deeper intuition for what is going on. We've seen how we can refactor to strict tail recursion via the "plug-and-chug" strategy. We've seen how useful it is to benchmark our code to point us in the right direction.

The final implementation still has a lot of wiggle room for optimization, and there are many other multiplication algorithms out there. Exploring those is left as an exercise to the reader.

