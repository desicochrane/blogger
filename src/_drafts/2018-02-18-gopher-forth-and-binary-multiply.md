---
title: Gopher Forth and Binary Multiply
date: 18-Feb-2018
---

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
