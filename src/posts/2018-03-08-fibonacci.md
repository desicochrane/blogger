---
title: The Fibonacci Problem
date: 08-Mar-2018
---

# The Fibonacci Problem

### Introduction
In this article we will investigate the famous Fibonacci sequence from the perspective of tail-recursion. We will see that the Fibonacci sequence requires more advanced reasoning than the problem we explored in a [previous article](2018-02-22-gopher-forth-and-multiply.html) and we will practice new tactics in our pursuit of refactoring to a strict tail-recursive solution.

We will use Golang for its built-in benchmarking, but it should be easy to follow even if you've never seen Golang before.

### The Problem
Over 800 years ago Leonardo of Pisa posed a question that went something like this:

> Suppose we have a newly-born fluffy. Fluffies are magical creatures - they never die and when they are one month old they become adult fluffies. One month after a fluffy has become an adult it will produce another newly-born fluffy - and then again every month after. Given this process repeats for every newly-born fluffy, how many fluffies will there be after one year has passed?

If we then denote the number of fluffies in a given month \\(n\\) as \\(F\_n\\) then our task is to find \\( F\_{12} \\).

As a start, we make the observation that the number of fluffies in any given month is the sum of the adult fluffies and baby fluffies in that month. If we denote the number of Adult fluffies and Baby fluffies in month \\(n\\) as \\(A\_n\\) and \\(B\_n\\) respectively, we have that:

$$ F\_n = A\_n + B\_n $$

Let's first reason about the number of **adult** fluffies in a given month \\(n\\). Since fluffies never die, all the adult fluffies from the previous month \\( (n-1) \\) will still exist this month \\(n\\). Also, fluffies become adults after only one month, this means that all the baby fluffies from the previous month have now become adult fluffies this month. 

Thus we have that the number of adult fluffies in a given month \\(n\\) will be the sum of the adult fluffies and baby fluffies from the previous month \\((n-1)\\):

\\[ \\begin{aligned}
A\_n &= A\_{n-1} + B\_{n-1} & \\\\
     &= F\_{n-1}            & (\text{above definition of } F\_n)
\\end{aligned} \\]

Next, let's examine the **baby** fluffies in a given month. Because each adult fluffy from the previous month will produce a single baby fluffy the next month we have that the number of baby fluffies in a any given month must equal to the number of adult fluffies in the previous month:
 
\\[ \\begin{aligned}
B\_{n} &= A\_{n-1} \\\\
       &= F\_{n-2}   & (\text{above definition of } A\_n)
\\end{aligned} \\]

With this improved understanding of \\(A\_n\\) and \\(B\_n\\) we can re-write our definition of \\(F\_n\\):

\\[ \\begin{aligned}
F\_n &= A\_{n} + B\_{n} \\\\
     &= F\_{n-1} + F\_{n-2}
\\end{aligned} \\]

This means that the number of fluffies in a given month is determined by the number of fluffies in the previous two months. Thus if we know the number of fluffies in the first two months we can determine the number of fluffies in any subsequent month. We know that in the first month that we start with a single fluffy, that is \\(F\_1 = 1\\), and we can say that before the first month we had no fluffies at all so we know that \\(F\_0 = 0\\). 

This gives us all the pieces we need to define the \\(F_n\\) with a recursive relation:

\\[ \\begin{aligned}
F\_0 &= 0 \\\\
F\_1 &= 1 \\\\
F\_n &= F\_{n-1} + F\_{n-2} \\text{ where } n > 1
\\end{aligned} \\]

To solve Fibonacci's problem then we can use simply use the definition to compute from \\(F_0\\) to \\(F\_{12}\\):

\\[ 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144 \\]

\\[ \therefore F_{12} = 144 \\]

This famous sequence is known as the **Fibonacci sequence** and the \\(n^{th}\\) number in the sequence is known as the \\(n^{th}\\) **Fibonacci number**.

### Finding \\(F\_n\\) for \\(n > 12\\) 

Suppose then you are tasked with implementing a function to return \\(F_n\\) for arbitrary \\(n\\). To compare different implementations we will first define our interface as a function which takes an integer \\(n\\) and returns an integer:

```go
// fibonacci.go
package fibonacci

type Fib func(n int) int
```

As a first implementation it is tempting to codify our recursive definition directly as:
 
```go
// fibonacci.go

func FibNaive(n int) int {
    if n < 2 {
        return n
    }
    
    return FibNaive(n-1) + FibNaive(n-2)
}
```

When I ran some benchmarks for this naive implementation for various values of \\(n\\) and got the following:
 
| \\( n \\)   | `Naive`           |
|-------------|-------------------|
| \\( 12 \\)  | 1,062 ns/op       |
| \\( 40 \\)  | 816,681,704 ns/op |
| \\( 90 \\)  | way too long      |

We can see that our solution is not scaling well at all - in fact for \\(n=90\\) my mac ran out of memory. What's going on here? We can "plug and chug" through our solution for \\(n = 5\\) to get a feel for what is happening:

\\[ \\begin{aligned}
F\_5 &= F\_3 + F\_4 \\\\
    &= (F\_1 + F\_2) + (F\_2 + F\_3) \\\\
    &= (1 + (F\_0 + F\_1)) + ((F\_0 + F\_1) + (F\_1 + F\_2)) \\\\
    &= (1 + (0 + 1)) + ((0 + 1) + (1 + (F\_0 + F\_1))) \\\\
    &= (1 + (0 + 1)) + ((0 + 1) + (1 + (0 + 1))) \\\\
    &= 5
\\end{aligned} \\]

Here we can see the problem - our solution seems to do a lot of duplicate work. We can see that in the case for \\(n=5\\) we are computing \\(F\_3\\) twice and \\(F\_2\\) three times! This is even more apparent when we visualize our algorithm with a graph:

<p style="text-align: center; margin: 50px 0;">
<img style="max-width: 500px;" src="../img/fibonacci-naive.dot.svg">
</p>

Our graph shows our algorithm behaves like a tree and seems to grow in an exponential manner. In fact we can see that as \\(n\\) increases our tree will have a height of approximately \\(n\\), with each layer having approximately double the nodes of the above layer. Since the addition operations happen in the leaves, we can see there will be \\(2^n\\) numbers to add up! Also each node in the entire graph roughly corresponds to another function call added to our call-stack which means we are taking exponential space as well! 

> It's no wonder then that I ran out of memory trying to perform `Fib(90)`. \\(2^{90}\\) is an obscenely large number and too much for my mac to handle.

### Remember what you saw here
The results of our first implementation do not seem to coincide with our own experience when we followed the recursive definition to compute the Fibonacci numbers by hand. This is because we did not need to perform duplicate computation, but rather we re-used our previous computation. The graph representation for our manual approach is then:

<p style="text-align: center; margin: 50px 0;">
<img style="max-width: 500px;" src="../img/fibonacci-cached.dot.svg">
</p>

To move our implementation more in alignment with this linear graph, we will cache our intermediate results so that we do not have to recompute them:

```go
// fibonacci.go

func FibCached(n int) int {
  return FibCachedHelper(n, make(map[int]int))
}

func FibCachedHelper(n int, cache map[int]int) int {
  if n < 2 {
    return n
  }

  a, cached := cache[n-2]
  if !cached {
    a = FibCachedHelper(n-2, cache)
  }

  b, cached := cache[n-1]
  if !cached {
    b = FibCachedHelper(n-1, cache)
  }

  result := a + b
  cache[n] = result

  return result
}
```

In our implementation you can see that in order to adhere to our interface we used a helper function which takes a cache as an argument. We then check this cache before performing any recursive calls and we cache our results after each addition. 

When I now run the benchmarks I get:

| \\( n \\)   | `Naive`           | `Cached`     |
|-------------|-------------------|--------------|
| \\( 12 \\)  | 1,062 ns/op       | 1,120 ns/op  |
| \\( 40 \\)  | 816,681,704 ns/op | 7,331 ns/op  |
| \\( 90 \\)  | way too long      | 14,488 ns/op |

This improved implementation has made calculating `Fib(90)` tractable and the performance appears to be scaling linearly with our input much like our graph suggested, and indeed our algorithm now gone from order \\(O(2^n)\\) to \\(O(n)\\).

> When I first implemented the cached version I used a global variable for the cache and my benchmarks were unreasonably fast. This is because the first time Golang executed `Fib` the final value was cached and available for subsequent executions. So calculating `Fib(90)` was as fast as cache lookup! Let that be a lesson to be careful of shared memory when benchmarking.

### Can we do better? Tail recursion revisited
In a [previous article](2018-02-22-gopher-forth-and-multiply.html) we saw how refactoring our recursive code to be **strict tail-recursive** resulted in a significant performance boost. However, unlike in that case here our deferred operation is **another** recursive call!

One tactic we can employ is "Replace \\(n\\) recursive calls with one". That is, we currently have two recursive calls to a function that returns a single integer and we can refactor this into a single new function that will return two integers. 

We can return two integers from a single function by grouping them into a tuple or vector. We can describe this refactoring as:

\\[ \\begin{aligned}
F\_n &= F\_{n-1} + F\_{n-2} \\\\
\\text{ becomes: } F\_n &= \sum{\vec{v}} \\text{ where } \vec{v} = \\begin{bmatrix} F\_{n-1} \\\\ F\_{n-2} \\end{bmatrix}
\\end{aligned} \\]

As a convention we will refer to \\( \\begin{bmatrix} F\_n \\\\ F\_{n-1} \\end{bmatrix}\\) as the \\( n^{th} \\) **Fibonacci vector** denoted \\( \\vec{F}_n \\). Then the \\(n^{th}\\) Fibonacci number can be defined as the vector sum of the \\( (n-1)^{th} \\) Fibonacci vector:

\\[
F\_n = \sum{ \\begin{bmatrix} F\_{n-1} \\\\ F\_{n-2} \\end{bmatrix}} = \sum{ \\vec{F}\_{n-1} } \\text{ where } n > 1
\\]

> Notice that \\( \\vec{F}\_1 \\) is the smallest Fibonacci vector because \\( \vec{F}\_0 \\) would be \\( \begin{bmatrix} F\_0 \\\ F\_{(-1)} \end{bmatrix} \\) but \\( F\_{(-1)} \\) is undefined.

We can flesh out the boilerplate for this refactoring from our naive implementation as:

```go
// fibonacci.go

func FibVecSum(n int) int {
  if n < 2 {
    return n
  }

  a, b := FibVec(n - 1)

  return a + b
}

func FibVec(n int) (int, int) {
  // todo
}
```

Here we can see our new function `FibVecSum` works by calling a single function `FibVec` and summing the two integers it returns. This now begs the question of how to implement `FibVec`. 

Since the first item in the Fibonacci vector sequence is \\( \vec{F}\_1 = \begin{bmatrix} F\_1 \\\ F\_0 \end{bmatrix} = \begin{bmatrix} 1 \\\ 0 \end{bmatrix} \\) we can generate the rest of the sequence by hand:

$$ \vec{F}\_n := \begin{bmatrix} 1 \\\ 0 \end{bmatrix}, \begin{bmatrix} 1 \\\ 1 \end{bmatrix}, \begin{bmatrix} 2 \\\ 1 \end{bmatrix}, \begin{bmatrix} 3 \\\ 2 \end{bmatrix}, \begin{bmatrix} 5 \\\ 3 \end{bmatrix}, ... $$

To understand its recursive behaviour we can examine \\( \vec{F}\_n \\) vs \\( \vec{F}\_{n+1} \\):

\\[
\vec{F}\_n = \begin{bmatrix} F\_{n} \\\ F\_{n-1} \end{bmatrix}
\\]

\\[
\vec{F}\_{n+1} = \begin{bmatrix} F\_{n+1} \\\ F\_{n} \end{bmatrix} = \begin{bmatrix} F\_{n} + F\_{n-1} \\\ F\_{n-1} \end{bmatrix}
\\]

Notice the components of \\( \\vec{F}\_{n+1} \\) can be computed from the components of \\( \\vec{F}\_n \\) as so:

 $$ \\text{if } \\vec{F}\_n = \begin{bmatrix} a \\\ b \end{bmatrix} \\text{ then } \\vec{F}\_{n+1} = \begin{bmatrix} a + b \\\ a \end{bmatrix} $$ 

If we define a transform function \\( T\Big(\begin{bmatrix} a \\\ b \end{bmatrix}\Big) = \begin{bmatrix} a + b \\\ a \end{bmatrix}  \\) we can then define \\( \vec{F}_n \\) with a recursive relation:

\\[ \\begin{aligned}
\vec{F}\_1 &= \begin{bmatrix} 1 \\\ 0 \end{bmatrix} \\\\
\vec{F}\_n       &= T(F\_{n-1}) \\text{ where } n > 1
\\end{aligned} \\]

Translating this into Golang we can now implement `FibVec`:

```go
// fibonacci.go

func FibVecSum(n int) int {
  if n < 2 {
    return n
  }

  a, b := FibVec(n - 1)

  return a + b
}

func FibVec(n int) (int, int) {
  if n == 1 {
    return 1, 0
  }

  return FibVecTransform(FibVec(n - 1))
}

func FibVecTransform(a int, b int) (int, int) {
  return a + b, a
}
```

At this point our refactored implementation of `Fib` now contains one recursive call per function. That is we have removed any deferred recursive call operations and we are one step closer to a strict tail-recursive implementation. 

Upon benchmarking this refactored implementation we get quite a performance boost:

| \\( n \\)   | `Naive`           | `Cached`     | `VecSum`    |
|-------------|-------------------|--------------|-------------|
| \\( 12 \\)  | 1,062 ns/op       | 1,120 ns/op  |  36.7 ns/op |
| \\( 40 \\)  | 816,681,704 ns/op | 7,331 ns/op  | 151 ns/op   |
| \\( 90 \\)  | way too long      | 14,488 ns/op | 405 ns/op   |

That's a whopping improvement! It seems we are on the right track.

### Strict tail-recursion

At the moment our recursive function `FibVec`  has no deferred **recursive calls** but still has a deferred operation, namely applying the `FibVecTransform`. If we want to refactor our solution to be strict tail-recursive we need to understand how this deferred operation behaves so that we can accumulate it into our recursive call.

If we plug and chug through the case where \\(n=5\\) we can get a better feeling for how this is behaving:

\\[ \\begin{aligned}
\\text{FibVec}(4) &= T \circ \\text{FibVec}(3) \\\\
                  &= T \circ T \circ \\text{FibVec}(2) \\\\
                  &= T \circ T \circ T \circ \\text{FibVec}(1) \\\\
                  &= T \circ T \circ T \Big( \begin{bmatrix} 1 \\\ 0 \end{bmatrix} \Big) \\\\
                  &= T \circ T \Big( \begin{bmatrix} 1 \\\ 1 \end{bmatrix} \Big) \\\\
                  &= T \Big( \begin{bmatrix} 2 \\\ 1 \end{bmatrix} \Big) \\\\
                  &= \begin{bmatrix} 3 \\\ 2 \end{bmatrix}
\\end{aligned} \\]

> Note we are using \\(T\\) to denote `FibVecTransform` and the notation \\(T \circ T(x) \\) to denote \\(T(T(x))\\)

Note that that each recursive call results in an additional deferred application of our transform function that we need to perform to our base case of \\( \vec{F}\_1 = \begin{bmatrix} 1 \\\ 0 \end{bmatrix} \\). Then by the time we reach our base case, we will always have accumulated \\( (n-1) \\) deferred functions to apply:
 
\\[
   \\text{FibVec}(n) = \underbrace{T \circ T \circ ... T}\_{n-1} \Big( \begin{bmatrix} 1 \\\ 0 \end{bmatrix} \Big)
\\]

The trick here is that we can create an accumulator argument to act as a count-down for the remaining times we need to apply our deferred operation. Then our desired "plug and chug" output will look like:


\\[ \\begin{aligned}
\\text{FibVec}(4) &= \\text{FibVec}^{\prime}\Big(4, \begin{bmatrix} 1 \\\ 0 \end{bmatrix}\Big) \\\\
                  &= \\text{FibVec}^{\prime}\Big(3, \begin{bmatrix} 1 \\\ 1 \end{bmatrix}\Big) \\\\
                  &= \\text{FibVec}^{\prime}\Big(2, \begin{bmatrix} 2 \\\ 1 \end{bmatrix}\Big) \\\\
                  &= \\text{FibVec}^{\prime}\Big(1, \begin{bmatrix} 3 \\\ 2 \end{bmatrix}\Big) \\\\
                  &= \begin{bmatrix} 3 \\\ 2 \end{bmatrix}
\\end{aligned} \\]

Using this as instructions we can define our new recursive function \\( \\text{FibVec}^{\prime} \\) as:

\\[
\\text{FibVec}^{\prime}\Big(\\text{acc}, \begin{bmatrix} a \\\ b \end{bmatrix} \Big) =
\\begin{cases}
  \begin{bmatrix} a \\\ b \end{bmatrix}                                    & \\text{if } \\text{acc} = 1 \\\\
  \\text{FibVec}^{\prime}\Big(\\text{acc}-1, \begin{bmatrix} a+b \\\ a \end{bmatrix} \Big) & \\text{if } \\text{acc} > 1 \\\\
\\end{cases}
\\]

> Notice this is more flexible than our original definition of \\(\\vec{F}\_n \\) because it allows for an arbitrary initial conditions of \\(a\\) and \\(b\\) in the base case.


We can directly implement this in Golang with strict tail-recursion as:

```go
// fibonacci.go

func FibTailVecSum(n int) int {
  if n < 2 {
    return n
  }

  a, b = FibTailVec(n-1, 1, 0)

  return a + b
}

func FibTailVec(acc int, a int, b int) (int, int) {
  if acc == 1 {
    return a, b
  }

  acc, a, b = acc-1, a+b, a
  return FibTailVec(acc, a, b)
}
```

Before we refactor this to iterative implementation let's look at our benchmarks for comparison:


| \\( n \\)   | `Naive`           | `Cached`     | `VecSum`    | `TailVecSum` |
|-------------|-------------------|--------------|-------------|--------------|
| \\( 12 \\)  | 1,062 ns/op       | 1,120 ns/op  |  36.7 ns/op | 36.0 ns/op   |
| \\( 40 \\)  | 816,681,704 ns/op | 7,331 ns/op  | 151 ns/op   | 145 ns/op    |
| \\( 90 \\)  | way too long      | 14,488 ns/op | 405 ns/op   | 395 ns/op    |

We can see that our refactorings made very little difference, which is a good sign.

### Iteration
Now that our function for computing \\(\\vec{F}_n \\) is in a strict tail-recursive form, we can directly replacing our  recursive call with `for` loop without changing any other code:


```go
// fibonacci.go

func FibIterative(n int) int {
  if n < 2 {
    return n
  }

  a, b := FibIterativeVec(n-1, 1, 0)

  return a + b
}

func FibIterativeVec(acc int, a int, b int) (int, int) {
  for {
    if acc == 1 {
      return a, b
    }

    acc, a, b = acc-1, a+b, a
  }
}
```

Here we can see that instead of returning \\(a\\) and \\(b\\) as a tuple we could simply return their sum thus removing the deferred operation in the main function:

```go
// fibonacci.go

func FibIterative(n int) int {
  if n < 2 {
    return n
  }

  return FibIterativeVec(n-1, 1, 0)
}

func FibIterativeVec(acc int, a int, b int) int {
  for {
    if acc == 1 {
      return a + b // return the sum
    }

    acc, a, b = acc-1, a+b, a
  }
}
```

Now that the main function has no deferred operations and our helper function is iterative, we can inline our logic to a single function, which also means we can use \\(n\\) as our accumulator. The final, elegant solution is:

```go
// fibonacci.go

func FibIterative(n int) int {
  if n < 2 {
    return n
  }

  a, b := 1, 0

  for {
    if n == 2 {
      return a + b
    }

    n, a, b = n-1, a+b, a
  }
}
```


Looking at the final benchmarks we can give ourselves a pat on the back for our improvements:

| \\( n \\)   | `Naive`           | `Cached`     | `VecSum`    | `TailVecSum` | `Iterative` |
|-------------|-------------------|--------------|-------------|--------------|-------------|
| \\( 12 \\)  | 1,062 ns/op       | 1,120 ns/op  |  36.7 ns/op | 36.0 ns/op   | 7.67 ns/op  |
| \\( 40 \\)  | 816,681,704 ns/op | 7,331 ns/op  | 151 ns/op   | 145 ns/op    | 30.2 ns/op  |
| \\( 90 \\)  | way too long      | 14,488 ns/op | 405 ns/op   | 395 ns/op    | 51.1 ns/op  |


### Summary
In this article we spent some time reasoning about recursive functions with a deferred recursive operation. We saw how to understand why a recursive function was super slow and how caching intermediate results made our problem tractable.  We saw that a tactic for simplifying recursive functions can be to replace multiple recursive calls with a new recursive function that returns a vector. We saw how the "plug and chug method" can help us to understand our functions deferred behaviour, and we saw that when our deferred operations are repeated function calls we can treat the final function call count as our accumulator when refactoring to strict tail-recursion.

In terms of performance our final result ended up significantly better than our initial naive implementation. In a future article I would like to explore how various properties of vectors might help us further improve our solution as well as exploring the Fibonacci sequence in a more general context as a linear recurrence.

### Additional resources
If you're interested in the Fibonacci sequence and linear recurrences, I can highly recommend checking out the following resources:

- [Finding Fibonacci - Keith Devlin](https://www.youtube.com/watch?v=vzKRV7AoLKQ)
- [Introduction to Enumerative Combinatorics](https://www.coursera.org/learn/enumerative-combinatorics)
- [Mathematics for Computer Science (Weeks 12 - 15)](https://www.coursera.org/learn/enumerative-combinatorics)
- [Fibonacci Numbers Wikipedia](https://en.wikipedia.org/wiki/Fibonacci_number)
