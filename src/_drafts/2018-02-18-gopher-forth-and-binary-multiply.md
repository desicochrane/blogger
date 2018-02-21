---
title: Gopher Forth and Binary Multiply
date: 18-Feb-2018
---

### Understanding what's happening

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


### Tail recursion
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
