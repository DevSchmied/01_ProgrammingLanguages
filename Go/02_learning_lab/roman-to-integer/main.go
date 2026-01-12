package main

/*
LeetCode — Roman to Integer

Description:
Roman numerals are represented by seven different symbols:

Symbol  Value
I       1
V       5
X       10
L       50
C       100
D       500
M       1000

Roman numerals are generally written from largest to smallest from left to right.
However, in some cases subtraction is used:

- I before V or X  → 4, 9
- X before L or C  → 40, 90
- C before D or M  → 400, 900

Given a valid Roman numeral string, convert it into its integer value.

Examples:
Input:  "III"
Output: 3

Input:  "LVIII"
Output: 58

Input:  "MCMXCIV"
Output: 1994

Constraints:
- 1 <= s.length <= 15
- s contains only characters: I, V, X, L, C, D, M
- s is guaranteed to be a valid Roman numeral in the range [1, 3999]
*/
