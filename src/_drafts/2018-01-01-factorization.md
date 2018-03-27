---
title: [wip] Factorization
date: 01-Jan-2018
---
> for which integers \\(n > 1\\) is it possible to arrange \\(n\\) apples?

## Primes

## sieving?

if p is prime and p|ab then p|a or p|b. (euclids lemma)
proof:
consider
  gcd(p,a) | p (tautology)
but the only numbers that can divide a prime, are 1 and the prime itself.
 therefore gcd(p,a) = 1 or p
suppose 
  p does not divide a.
then gcd(p,a) = 1 
therefore there exists some x,y such that:
  ax + py = 1
and:
  ax \equiv 1 (mod p) for some x.
  
drumroll:
  p | ab 
implies that 
  ab \equiv 0 (mod p)    <no remainder>
  xab \equiv 0 (mod p)   <both sides times x>
  b \equiv 0 (mod p)     <x * a = 1 mod p>
therefore p|b

corollary:
if p | a1*a2...*ak then p divides at least 1 of a1*a2...*ak.

## fundamental theorem of arithmetic:

suppose there are 2 different prime representations of m:

m = p1*p2*p3...*pk = q1*q2...qn

we can cancel any common prime factors from both sides until all the primes on the left differ from the primes on the right.

assume p1 is a prime on the left, then p1|left
therefore p1 must divide the right
therefore p1 must divide one of the right primes
the only two factors of a prime are 1 and the prime. 
we know p1 is not 1, because p1 is prime.
therefore p1 must equal q1.
contradiction yo.



---
lemma: if a|n and b|n and gcd(a,b) = 1, then ab|n

if n \equiv 1 mod 6 then what can be n mod 4?
well, we know that n must be odd. so n mod 4 is either 1 or 3. so knowing the remainder mod 6 gave us some information about the remainder mod 4.


## chinese remainder theorem
> if gcd(a,b) = 1 then, then for any remainder r_a (n mod a) and any remainder r_b (m mod b) there exists some n such that
n mod a = n mod b, if n1 and n2 both satisfy this, then n1 = n2 mod ab


## primes
> 1. if n is *not* a prime, 2^{n}-1 is not a prime (contrapositive? p -> q !q -> !p)
> if n is a prime, 2^{n}-2 is a multiple of 2n
> if n is a prime, and p|2^{n}-1 then n|(p-1)

1. if 2^n - 1 is prime then n is prime
(proof) 
suppose n is not prime, then: 
  n = uv where u,v > 1
then
  2^n - 1 &= 2^{uv} - 1
          &= (2^u)^v - 1
          &= (2^u)^v - 1^v
          &= (2^u - 1)((2^u)^{v-1} + (2^u)^{v-2} + ... + 2^u + 1) <niceeeee trick!>
          
          which shows two factors greater than 1. therefore 2^n-1 is not prime.
          
          
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
  

 
