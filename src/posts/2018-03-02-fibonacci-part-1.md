---
title: The Fibonacci Problem - Part I
date: 02-Mar-2018
---

# The Fibonacci Problem - Part I

> This article is under active development, check back again in a week or so.

### Introduction
> todo

### The Problem
Over 800 years ago Leonardo of Pisa, nowadays known as **Fibonacci** (son of Bonacci), posed a question that went something like this:

> Suppose we have a newly-born pair of baby fluffies, a boy and a girl. Fluffies are magical creatures - they never die and when they are one month old they become adult fluffies. One month after a boy-girl pair of fluffies become adults, they will produce another newly-born pair of baby boy and girl fluffies, and then again every month after. Given this process repeats for every new pair of baby fluffies, how many pairs of fluffies will there be after one year has passed?

To simplify the language of the problem, we will consider each **pair** of fluffies as a single unit, so when we refer to the "number of fluffies" we are really referring to the "number of **pairs** of fluffies"

If we then denote the number of fluffies in a given month \\(n\\) as \\(F_n\\) then our task is to find \\( F\_{12} \\).

As a start, we make the observation that the number of fluffies in any given month is the sum of the adult fluffies and baby fluffies in that month. If we denote the number of **a**dult fluffies and **b**aby fluffies in month \\(n\\) as \\(a\_n\\) and \\(b\_n\\) respectively, we have that:

$$ F\_n = a\_n + b\_n $$

Let's first examine the number of adult fluffies. Observe that the number of adult fluffies in a given month will equal the number adult fluffies from the previous month **plus** all the baby fluffies from the previous month which have now become adults

\\[ \\begin{aligned}
a\_n &= a\_{n-1} + b\_{n-1} & \\\\
     &= F\_{n-1}            & (\text{above definition of } F\_n)
\\end{aligned} \\]

Next, let's examine the baby fluffies. Because every adult fluffy in a given month will each produce a baby fluffy the next month we have that the number of baby fluffies in a any given month must equal to the number of adult fluffies in the previous month:
 
\\[ \\begin{aligned}
b\_{n} &= a\_{n-1} \\\\
       &= F\_{n-2}   & (\text{above definition of } a\_n)
\\end{aligned} \\]


With this improved understanding of \\(a\_n\\) and \\(b\_n\\) we can re-write our definition of \\(F\_n\\):

\\[ \\begin{aligned}
F\_n &= a\_{n} + b\_{n} \\\\
     &= F\_{n-1} + F\_{n-2}
\\end{aligned} \\]

Which means that the number of fluffies in a given month is completely determined by the number of fluffies in the previous two months. Since we started the first month with one pair of fluffies, and before that (in month zero) we had no fluffies then we can define the number of fluffies as a recursive function:

\\[
F\_n =
\\begin{cases}
  0                   & \\text{if } n = 0 \\\\
  1                   & \\text{if } n = 1 \\\\
  F\_{n-1} + F\_{n-2} & \\text{if } n > 1 \\\\
\\end{cases}
\\]

Which is the definition of the famous fibonacci sequence, where each number is equal to the sum of the previous 2 numbers and the first two numbers are 0 then 1. 

Because our sequence starts at month zero, we can manually iterate over the first 13 numbers in the sequence to solve get the number of fluffies in month 12 and solve Fibonacci's problem:

\\[ 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144 \\]

\\[ \therefore F_{12} = 144 \\]

### Finding \\(F\_n\\) for arbitrary \\(n\\) 
> todo

### Improvement: Remember what you saw here
> todo

### Can we do better? Tail recursion revisited
> todo

### Strict tail recursion
> todo

### Summary
