---
title: [wip] Bezouts Lemma
date: 01-Jan-2018
---
http://www.theoremoftheday.org/NumberTheory/BezoutId/TotDBezoutId.pdf
http://www.math.uconn.edu/~kconrad/blurbs/ugradnumthy/divgcd.pdf

- First show that if d|a and d|b and d = ax + by then d must be gcd(a,b)

- but will x,y  always exist? We can prove it by the "well ordering principle"

- look at all possible ax+by

- show that there must be at least some positive examples

- if there are some positive examples, there must be a smallest possible example.

- call this smallest example "s"

- show that "s" must divide all other elements of the set:

- pick any element of the set, call it "u"

- we need to prove that "s" divides "u"

- we know that u = s*q + r, all numbers have this property.

- we need to show that r must be zero to imply that s|u

- r = u - s*q
= (ax_u + by_u) - (ax_s + by_s)q
= ax_u + by_u - ax_sq - by_sq
= a(x_u - ax_sq) + b(y_u  - y_sq)
 but x_u - ax_sq and y_u  - y_sq are just integers
 so r is also of the form ax_r + by_r
 
 r by definition is 0 <= r < s
 so r is zero OR positive. 
 but we assumed that s is the smallest positive number in the set so cannot be less than s
so r must be zero.
therefore s exists and s divides every other element

then d | (ax + by) and d|a and d|b 

boom.

https://www.youtube.com/watch?v=OGxeVKApdLA
https://www.youtube.com/watch?v=9KM6bX2rud8

Other proof: is the inductive proof, like going through euclids algorithm and building up a recurrence relation
https://www.youtube.com/watch?v=9Y3dp_p_Z1Y

step 1: euclids algorithm by a = bq1 + r1 -> b = r1q2 + r2
a,b = 39,15 then:
39 = 15(2) + 9
15 = 9(1)  + 6
9  = 6(1)  + 3
6  = 3(2) + 0
therefore: 3

step 2: re-arrange for r (note to ignore last, because second last is our d and we need equation to be d = ax + by):
39 - 15(2) = 9
15 - 9(1)  = 6
9  - 6(1)  = 3

step 3: bottom up, substitute:
9  - 6(1)  = 3
-> 9  - (15 - 9(1))(1)  = 3
-> 9  - 15 + 9  = 3
-> 9(2) - 15  = 3

9(2) - 15  = 3
-> (39 - 15(2))(2) - 15  = 3
-> 39(2) - 15(5)  = 3
 boom.
- also show why all solutions follow the given pattern


explain for the case of inverse modular n!!
https://www.youtube.com/watch?v=shaQZg8bqUM

explain the gallon jugs: all possible states will be a linear combination of a,b,x,y.

4 = 3(-2) + 5(2)                 | m = as + bt
4 = 3(-2) + 5(2) + 3(5) - 3(5)   | m = as + bt + ab - ab
4 = 3(-2) + 5(2) + 3(5) - 5(3)   | m = as + bt + ab - ba
4 = 3(-2+5) + 5(2-3)             | m = as + ab + bt - ba
4 = 3(3) + 5(-1)                 | m = a(s + b) + b(t - a)


  m = as + bt
+     ab - ba
-------------
  m = a(s+b) + b(t-a)
  
can repeat k times:

  m =  as + bt
+     kab - kba
-------------
  m = a(s+kb) + b(t-ka) choosing k such that s + kb > 0

this proves that we can change s and t to get a new solution, but how can we get every solution? what is the pattern?

  m = ax + by = (ds)x + (dt)y
             +     st - ts
-----------------------------
  m =        s(dx + t) + t(dy - s)



s + kb > 0
kb > s
k > s/b if b > 0, k < s/b if b < 0 (undefined for b = 0)
s > s/b

notice: transferring between, i.e "pouring" does not change sum
so we can make a rule, we only "add" to one, and "subtract" from the other: sum = as -bt

then the times needed will be pouring s' times (we can make it positive) and pouring from bigger one t times!
(why does a need to be less than b?)


> invariant, at every step of state machine (euclids algorithm), gcd(x,y) = gcd(x,y) and both X and Y are linear combinations of a and b, that is x = as+bt and y = ap+bq
base case: trivial. gcd(a,b) = gcd(a,b) a = 1a + 0b, b = 0a + 1b
inductive: assume invariant true for gcd(x,y) then y = qx + r, r = y - qx. which is linear combination of y and x, which is are both a linear combination of a and b. 
therefore at each step it is true.

notice remainder gets smaller and smaller, must reach a minimum (well ordering)


thm: the gcd(a,b) is the smallest possible linear combination of a and b.


---
thm: ax + by = c has integer solutions for x,y iff gcd(a,b) | c
this claims two things: p -> q and q -> p

weaker claim: first let's prove that if ax + by has an integer solution, then the gcd(a,b) | c
if
 ax + by = c
 let gcd(a,b) = d
 a = ds
 b = dt
 then:
  (ds)x + (dt)y = c
  d(sx + ty) = c therefore d|c

weaker claim: first now prove that if gcd(a,b) | c then there is an integer solution to ax + by = c for x,y
we know by bezouts lemma that: 
 there exists some s,t such that as + bt = gcd(a,b)
 but c is a multiple of gcd(a,b) therefore c = k(gcd,ab) = k(as + bt) = a(ks) + b(kt) and x,y = ks,kt
--- BORING! it assumes on the interesting proof (bezouts lemma) as a given

so this gives us a solution to a diophantine equation, and a condition to show if a solution even exists. but are there more than one solution (yes)

// todo: can draw a line, showing the non integral solutions or all real solutions.

thm:
let gcd(a,b) = d
then: 
  a = dp <=> p = a/d 
  b = dq <=> q = b/d
  
if 
  x,y is a solution to ax + by = c
then
  *all* solutions have the form:
  a(x + tq) + b(y - tp) = c
  a(x + t(b/d)) + b(y - t(a/d)) = c for any integer t
  
  c = ax + by
      abt/d - bat/d    (note ab/d is lcm(a,b))
      -------
      a(x + tb/d) + b(y - ta/d) = c

so it is clear that all equations in that form are solutions. but is that all of them?

a = dp
b = dq

consider two solutions: x,y and s,t 
  ax + by = c
- as + bt = c
=============
a(x-s) + b(y-t) = 0

divide by d.
p(x-s) + q(y-t) = 0
p(x-s) = q(t-y)

so p is a factor of q(t-y), p | q(t-y). since p and q are coprime, p | (t-y), t-y = pk
p(x-s) + q(y-t) = 0
p(x-s) + q(pk) = 0
p(x-s) + q(pk) = 0 (divide by p) why can we do this? how do we know it's not zero?
x-s + qk = 0
x-s = -qk
therefore:

a(x-s) + b(y-t) = 0
a(qk) + b(-pk) = 0
a(kb/d) + b(-ka/d) = 0
since: 
ax + by            = c
a(kb/d) + b(-ka/d) = 0
----------------------
a(x + kb/d) + b(y-ka/d) = c

a(x + tb/d) + b(y - ta/d) = c

question, is the multiple of LCM k(ab/d) significant?

3x + 5y = 22

gcd(3,5) = 1
therefore all solutions have form:
a(x+kb/d) + b(y-ka/d) = c
3(x+k5/1) + 9(y-k3/1) = 22
3(x+5k) + 9(y-3k) = 22

to find x,y:
3(2) + 5(-1) = 1
3(44) + 5(-22) = 22
x,y = 44,-22

all solutions:
3(44+5k) + 9(-22-3k) = 22 for any integer k.
3(9+5k) + 9(-1-3k) = 22 for any integer k.

therefore :
x = 44 + 5k >= 0
y = -22 - 3k <= 0

we want x >= 0:
5k >= -44
k >= -44/5 
k >= -8.8
k >= -8

and y <= 0:
-22-3k <= 0
-3k <= 22
k >= 22/-3
k >= -7.333...
k >= -7


## Modular Division
\\( a \times a^{-1} \equiv 1 (\mod n) \\) then \\( \frac{\displaystyle b}{\displaystyle  a} \equiv b \times a^{-1} (\mod n) \\)

note \\( \big(b \times a^{-1}\big) \times a = b \times \big(a^{-1} \times a\big) = b \times 1 = b \\)

ax = 1 (mod b) iff ax + by = 1


ax + by = 1 implies:
  y is inverse of b (mod a)
  y is inverse of b (mod x)
  x is inverse of a (mod b)
  x is inverse of a (mod y)
  
  
example: what is 7 / 2 (mod 9) ?
1. does division by 2 (mod 9) exist? yes. gcd(2,9) = 1
2. what is 2x + 9y = 1 ? 2(5) + 9(-1) = 1, therefore 5 is 1/2 (mod 9)
3. check: 7/2 = 7 * 5 = 35 = 8 (mod 9)
4. check: 8 * 2 = 16 = 7 (mod 9)



> another intuition! if ax + by = 1, and a,b are not co-prime, then we can divide both sides by common divisor which would result in a fraction on the right which cannot be possible in an integer world.
