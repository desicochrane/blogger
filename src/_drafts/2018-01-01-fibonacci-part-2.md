---
title: [wip] The Fibonacci Problem - Part II
date: 01-Jan-2018
---
# The Fibonacci Problem - Part II

## Linear Algebra

\\[
\begin{pmatrix} \alpha & \beta \\\ \gamma & \delta \end{pmatrix} \begin{pmatrix} a \\\ b \end{pmatrix} = \begin{pmatrix} a+b \\\ a \end{pmatrix}
\\]

$$ \alpha a + \beta b = a + b $$
$$ \gamma a + \delta b = a $$

\\[
g: \begin{pmatrix} a \\\ b \end{pmatrix} \to \begin{pmatrix} 1 & 1 \\\ 1 & 0 \end{pmatrix} \begin{pmatrix} a \\\ b \end{pmatrix} = \begin{pmatrix} a+b \\\ a \end{pmatrix}
\\]

> todo: Any linear transformation can be expressed by matrix vector product

> todo: show a visualization (tweening) of the plane doing the transform, with a vector starting at (1,0) https://shadanan.github.io/MatVis/

> http://ncase.me/matrix/

https://www.geeksforgeeks.org/gcd-and-fibonacci-numbers/

Starting with 5, every second Fibonacci number is the length of the hypotenuse of a right triangle with integer sides, or in other words, the largest number in a Pythagorean triple. The length of the longer leg of this triangle is equal to the sum of the three sides of the preceding triangle in this series of triangles, and the shorter leg is equal to the difference between the preceding bypassed Fibonacci number and the shorter leg of the preceding triangle.

https://www.geeksforgeeks.org/cassinis-identity/
https://www.geeksforgeeks.org/program-check-plus-perfect-number/

fibonaccia as a vector operation, then efficient by https://en.wikipedia.org/wiki/Exponentiation_by_squaring.

* What is the Fibonacci series?
* Where does it show up?
* What is its relation to golden ratio?
* What is it *really* though? Why is it interesting?

linear recurrence
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
