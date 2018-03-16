---
title: [wip] Tail Recursion
date: 01-Jan-2018
---


```go
func Collatz(n int) bool {
  if n == 1 {
    return true
  }
  
  if n%2 == 0 { 
    n = n / 2
  } else { 
    n = 3*n + 1
  }
  
  return Collatz(n)
}
```

```go
func factorial(n int) {
  if n == 1 {
    return 1
  }
  
  return n * factorial(n-1)
}
```

```go
func factorial2(n int, acc int) {
  if n == 1 {
    return acc
  }
  
  return factorial2(n-1,n*acc)
}
```
