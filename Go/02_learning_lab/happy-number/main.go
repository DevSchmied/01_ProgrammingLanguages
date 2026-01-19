package main

import (
	"fmt"
)

/*
LeetCode 202 — Happy Number

Problem Description:
Write an algorithm to determine if a number n is happy.

A happy number is defined by the following process:
Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay),
or it loops endlessly in a cycle which does not include 1.

Those numbers for which this process ends in 1 are happy.

Return true if n is a happy number, and false if not.

Examples:

Example 1:
Input: n = 19
Output: true
Explanation:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1

Example 2:
Input: n = 2
Output: false

Constraints:
1 <= n <= 2^31 - 1

Source:
This is a problem from the LeetCode platform.
*/

func isHappy(n int) bool {
	seen := make(map[int]struct{})
	for n != 1 {
		if _, exists := seen[n]; exists {
			return false
		}
		seen[n] = struct{}{}

		res := 0

		for n > 0 {
			digit := n % 10
			res = res + (digit * digit)
			n /= 10
		}
		n = res
	}
	return true
}

func main() {
	fmt.Println("======================================")
	fmt.Println("Solution:")
	fmt.Println("======================================")

	fmt.Printf("Example 1 | Expected: %t | Got: %t\n", true,
		isHappy(19))
	fmt.Println()

	fmt.Printf("Example 2 | Expected: %t | Got: %t\n", false,
		isHappy(2))
	fmt.Println()
}
