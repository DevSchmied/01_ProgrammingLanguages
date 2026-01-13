package main

import "fmt"

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

func romanToInt(s string) int {
	symbolValueMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0

	for idx := 0; idx < len(s); idx++ {
		if s[idx] == 'I' {
			if idx < len(s)-1 {
				if s[idx+1] == 'V' || s[idx+1] == 'X' {
					sum = sum + (symbolValueMap[s[idx+1]] - symbolValueMap[s[idx]])
					idx++
					continue
				}
			}
		}
		if s[idx] == 'X' {
			if idx < len(s)-1 {
				if s[idx+1] == 'L' || s[idx+1] == 'C' {
					sum = sum + (symbolValueMap[s[idx+1]] - symbolValueMap[s[idx]])
					idx++
					continue
				}
			}
		}
		if s[idx] == 'C' {
			if idx < len(s)-1 {
				if s[idx+1] == 'D' || s[idx+1] == 'M' {
					sum = sum + (symbolValueMap[s[idx+1]] - symbolValueMap[s[idx]])
					idx++
					continue
				}
			}
		}
		sum = sum + symbolValueMap[s[idx]]
	}

	return sum
}

func romanToIntOptimized(s string) int {
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && values[s[i]] < values[s[i+1]] {
			sum -= values[s[i]]
		} else {
			sum += values[s[i]]
		}
	}

	return sum
}

func main() {
	fmt.Println("===================1. Case, 1. Example===================")
	fmt.Println(romanToInt("III"))
	fmt.Println()

	fmt.Println("===================1. Case, 2. Example===================")
	fmt.Println(romanToInt("LVIII"))
	fmt.Println()

	fmt.Println("===================1. Case, 3. Example===================")
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println()
	fmt.Println()

	fmt.Println("===================2. Case, 1. Example===================")
	fmt.Println(romanToIntOptimized("III"))
	fmt.Println()

	fmt.Println("===================2. Case, 2. Example===================")
	fmt.Println(romanToIntOptimized("LVIII"))
	fmt.Println()

	fmt.Println("===================2. Case, 3. Example===================")
	fmt.Println(romanToIntOptimized("MCMXCIV"))

	fmt.Println("CI test")
}
