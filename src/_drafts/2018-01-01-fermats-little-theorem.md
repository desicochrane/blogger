---
title: [wip] fermat's little theorem
date: 01-Jan-2018
---
should include "encryption/decryption" m^ed mod prime! (special case of rsa)

> note: in all prime theorems/lemma explain **why** it only works for prime - give counter-examples
.

> Watson asks Sherlock: Given a string S of N 0's and M 1's, how many unique permutations of this string start with 1?
  Help Sherlock by printing the answer modulo (109+7).
  [hackerrank](https://www.hackerrank.com/challenges/sherlock-and-permutations/problem)
 
.
  
> A frequently used corollary of Fermat's Little Theorem is \\(a^p \equiv a \pmod {p} \\). As you can see, it is derived by multiplying both sides of the theorem by \\(a\\). The restated form is nice because we no longer need to restrict ourselves to integers \\(a\\) not divisible by \\(a\\) . [artofproblemsolving](https://artofproblemsolving.com/wiki/index.php?title=Fermat%27s_Little_Theorem)

.

> Some of the proofs of Fermat's little theorem given below depend on two simplifications. The first is that we may assume that \\(a\\) is in the range \\(0 \le a \le p - 1 \\). This is a simple consequence of the laws of modular arithmetic; we are simply saying that we may first reduce a modulo \\(p\\). [wikipedia](https://en.wikipedia.org/wiki/Proofs_of_Fermat%27s_little_theorem)

> Secondly, it suffices to prove that  \\( a^{p-1}\equiv 1 \pmod {p} \\) for a in the range \\(1 \le a \le p - 1 \\). Indeed, if the previous assertion holds for such a, multiplying both sides by a yields the original form of the theorem, \\( a^{p}\equiv a \pmod {p} \\). On the other hand, if \\(a = 0\\), the theorem holds trivially. [wikipedia](https://en.wikipedia.org/wiki/Proofs_of_Fermat%27s_little_theorem)

maybe interesting to first find the above, then show why should not be divisible by p because we "divide" by a.

### gcd(a,b) = 1 implies a^-1 exists (mod b)
**pre-requisites:**

- bezout's lemma
- pulverizer
- gcd
- divisibility mod n
- co-primes and inverse multiplication mod n ( to make it clear why it is cool to know things have gcd 1 )
- congruence preserved under multiplication

gcd(a,b) = c
ax + by = c
ax = c (mod b)
if c = 1 then:
ax = 1 (mod b)
then x is the multiplicative inverse of a. i.e.

ak = n (mod b)     <for any given k, n>
(ak)x = nx (mod b) <congruence preserved under multiplication>
(ka)x = nx (mod b) <multiplicative commutivity preserves congruence>
k(ax) = nx (mod b) <multiplicative associativity preserves congruence>
k(ax) = nx (mod b) <multiplicative associativity preserves congruence>
k(1) = nx (mod b)  <ax = 1 mod b>
k = nx (mod b)     <ax = 1 mod b>

effectively this means:
if a is co-prime with b, then you can divide by a in a mod b world.

---
3 people 32 minutes four walls
3 people 8 minutes for 1 wall
1 person 1 wall in 24 minutes

how many minutes 4 people to paint 7 walls.


18min 10mph

30mph

av speed

### gcd(p,a) = 1
**pre-requisites:**

- divisibility
- prime numbers
- modular arithmetic
- gcd
- co-primes and inverse multiplication mod n ( to make it clear why it is cool to know things have gcd 1 )

> For any prime \\(p\\) and integer \\(n\\): The **greatest common divisor** of \\(p\\) and \\(n\\) is \\(1\\) if and only if \\(n\\) is not a multiple of \\(p\\). 

\\[ \\begin{aligned}
p \nmid n \Longleftrightarrow \text{gcd}(p,n) = 1
\\end{aligned} \\]

*Proof*

- By definition, the gcd of \\(p\\) and \\(a\\) must divide \\(p\\)
- Because \\(p\\) is prime, by definition its only divisors are \\(1\\) and \\(p\\)
- Therefore the gcd of of \\(p\\) and \\(a\\) can only be \\(1\\) or \\(p\\)
- Since we know that \\(p\\) does not divide \\(a\\), the gcd can only be \\(1\\)

> Pro-tip: \\(a \mid b \\) is usually read "\\(a\\) divides \\(b\\)", but sometimes can be helpful to think of it as "\\(b\\) is a multiple of \\(a\\)". Likewise, \\(a \nmid b \\) can be read as "b is not a multiple of a"


> primality test, for fermat, maybe first sieve a bit

> fermat liars. how to construct them?


