---
title: [wip] Modular Exponentiation
date: 01-Jan-2018
---

\\( b^e \mod m \\)

can do fast modular exponentiation code.

> discrete log
> explain how not being monotonic means cannot perform binary search
> explain how log identities do not hold
> explain how not necessarily present or unique

## fermat's little theorem:
if prime p does not divide a, then 
a^{p-1} equiv 1 mod p
n^p - n | p
n^{p-1} - 1 | p

proof 1:
consider 1,2,...p-1 remainders (non zero) (why non-zero??)
???multiplying by a is invertible, since p does not divide a
???directed graph would be a 1to1 and onto a to ar, all incoming and outgoing edge degrees are 1
???this must be a cycle graph!
then a^k*r = r (mod p), which means a^k = 1 (mod p) 

note, with the graph, if we choose another remainder and do its cycle -> it cannot have any common nodes to other cycle otherwise would not be a cycle.
but both have length k.
therefore if there are C possible cycles, then c*k must equal p-1 (all the possible nodes)
therefore a^{p-1} = a^{c*k} = (a^k)^c = (a^k mod p)^c = 1^c mod p = 1 mod p

how to show we can speed up modular exponentiation with this trick? for mod p.

## euler's totient function
phi(n) counts the integers between 0 and n-1 which are co-prime with n
phi(1) = |{0}|   = 1 <gcd(1,0) = 1>
phi(2) = |{1}|   = 1 <gcd(2,0) = 2, gcd(2,1) = 1>
phi(3) = |{1,2}| = 2 <gcd(3,0) = 3, gcd(3,1) = 1, gcd(3,2) = 1>
phi(4) = |{1,3}| = 2 <gcd(4,0) = 4, gcd(4,1) = 1, gcd(4,2) = 2, gcd(4,3) = 1>

lemma: phi(prime) = prime - 1
lemma: phi(pq) = (p-1)(q-1) 
???-> chinese remainder theorem matrix pxq, only the 0 row and 0 col are not coprime, the remaining rectangle must be

> merry go round, prove it must hit every element.
r + ak = r mod p   where 0 < a < p
then
  p | (r + ak) - r
  p | ak
  a is less than p, so p cannot divide a.
  therefore
  p divides k
  therefore k = 0 or k = pn


## euler's theorem
if a is coprime with n then a^phi(n) \equiv 1 (mod n)
it's a generalization of fermat's little theorem

???can prove again with the graph cycle thing



n^2 - 1 is composite.


