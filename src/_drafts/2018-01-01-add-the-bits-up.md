---
title: [wip]  Bits Add
date: 01-Jan-2018
---
https://eater.net/8bit/

unsigned int myAdd(unsigned int a, unsigned int b)
{
    unsigned int carry = a & b;
    unsigned int result = a ^ b;
    while(carry != 0)
    {
        unsigned int shiftedcarry = carry << 1;
        carry = result & shiftedcarry;
        result ^= shiftedcarry;
    }
    return result;
}
