---
title: WIP Number theory 
date: 10-Feb-2018
---
# Fibonacci
* What is the Fibonacci problem?

> quote the original problem

Comes from famous book "Liber Abaci" written by Leonardo De Piza (check out Kieth Devlin video) written a long time ago (1202). He asked the following problem about rabbits:

Month | # Adult Rabbits | # Baby Rabbits |
----|---|---|
1     | 0 | 1 |

* What is the Fibonacci series?
* Where does it show up?
* What is its relation to golden ratio?
* What is it *really* though? Why is it interesting?

# tail recursion

linear recursion
( just anything where you can demonstrate same recurrence and same initial conditions e.g pascals triangle)
(number of tilings, can multiply by scalar, and add multiple together to show they also satisfy the recurrence
... but show that the initial 2 numbers uniquely determine that specific recurrence)
(number of ways of representing n as sum of 1's and 2's)
(vending machine problem!!!!) i.e. 5c, 10c, 25c, think of A(n) = #{ways if first coin is 5c} + #{first coin in 10c} etc
-> clever side, mapping 5c to 5n is a good example to illustrate power of bijection
Linear recurrence of "order n" A(n) = c_1*A(n-1) + c_2*A(n-2) ... ("c" for coefficient)
-> is uniquely determined by its first n terms
 
 WHY called linear? because if one set of "solutions" satisfy the recurrence (a0,a1,a2)
 ... and another solution a0',a1',a2', then any linear combination (i.e. adding and applying scalar)
 ... is also a solution to the recurrence! (i.e. also c*a1,c*a2,c*a3 etc)
 ... i.e. it forms a vector space! the fuck is a vector space yo?
 
 
 how about the baby case? i.e. A(n) = cA(n-1) => can see it is just a geometric progression!
 (note: being uniquerly determined -> there are no other sequences that determine it)
 
 idea: if teh baby case is a geometric progression, why not all?

### need to show:
gcd(a, b) = gcd(a, a mod b)

### can show:
if 
  d|a and d|b 
then
  d|(r = a mod b)

proof numerically:
if d|a then a = d*m, where m ∈ Int
if d|b then b = d*n, where n ∈ Int

then by definition:
a = b*k + r, where k ∈ Int
r = a - b*k
  = (d*m) - (d*n)*k
  = d(m - n*k)

since m, n, k ∈ Int, d|r

visually this is obvious:
|========================| a = 24
|================| b = 16
|========| d = gcd(24,16) = 8

|========|========|========| a = 24
|========|========| b = 16
|========| d = gcd(24,16) = 8


### maybe can prove by contradiction? if there was a greater number that divides r? then it would be greater than the gcd(a,b) ??

<article>
  <h2>What is number theory?</h2>
  <p>the study of the integers.</p>
  <p>...,-2,-1,0,1,2,3...</p>
  <p>but why tho? crypto?</p>
</article>
<article>
  <h2>what is crypto?</h2>
  <p>hiding numbers! (mit open courseware Lecture 4: Number Theory I https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-042j-mathematics-for-computer-science-fall-2010/video-lectures/lecture-4-number-theory-i/. )</p>
  <p>hiding information</p>
  <p>communication with an adversary (Lecture notes on cryptography pdf)</p>
</article>

definition of m|a
: m divides a. iff there exists some integer k such that m*k = a
e.g. 3|6
if a = 0, then any m|a simply by choosing k = 0
thereform for all m, m|0

thm. using jugs of water:
for some jug with "a" gallons, and another with "b" gallons, assuming a >= b
if m|a, and m|b, then m|(any possible jug filling)

consider state machine representing jugs!
<article>
  <h2>1. What is the GCD?</h2>
  <p>...</p>
</article>

<article>

</article>
1. Need to prove: gcd(a,b) = gcd(b, a - b)<br>
2. Need to prove that the sequence:
A_n = B_{n-1}
and
B_n = A_{n-1} - B_{n-1}

```js
const a = "yyolo"
console.log(a)
```
