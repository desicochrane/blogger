---
title: [wip] Bezouts Lemma
date: 01-Jan-2018
---

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


