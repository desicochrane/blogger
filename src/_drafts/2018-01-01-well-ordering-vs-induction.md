---
title: [wip] Well ordering vs induction
date: 01-Jan-2018
---

prove sum(1..n) = n(n+1)/2 with induction and with well-ordering, explain how the proofs are same but feel different.
https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-042j-mathematics-for-computer-science-spring-2015/readings/MIT6_042JS15_Session3.pdf

examples of well ordering:
- turtles all the way down
- every composite number can be expressed as a product of primes.

for me, "choosing the smallest" or assuming the smallest never clicked for me. I always though "well, maybe we were wrong that we chose the smallest?" so for me a better way is the "inductive" view. that it would repeat forever, in a finite set which cannot be.

> thm: the product of two positive integers smaller than a prime p is not divisible by p.
i.e: if p is prime and 0 < a,b < p then p does not divide a \times b.
ie. ab = qp + r where 0 < r < p

(proof) assume p|ab
then 
  ab = pk
divide p by b:
  p = bu + r
  ap = abu + ar
  ap - abu = ar
  ap - pku = ar
  p(a-ku) = ar
  
  whattt? this means that if p divides ab, (any ab) then it will also divide ar. where r is less than b. but this would repeat forever, which cannot be possible. (well ordering biatch!)
