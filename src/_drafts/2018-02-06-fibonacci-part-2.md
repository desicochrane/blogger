---
title: The Fibonacci Problem - Part II
date: 10-Feb-2018
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
