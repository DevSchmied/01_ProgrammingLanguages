package main

/*
LeetCode 796 — Rotate String

Given two strings s and goal, return true if and only if s can become goal
after some number of shifts on s.

A shift on s consists of moving the leftmost character of s to the rightmost position.

Example:
s = "abcde" → "bcdea" after one shift.

Examples:
Input:  s = "abcde", goal = "cdeab"
Output: true

Input:  s = "abcde", goal = "abced"
Output: false

Constraints:
1 <= s.length, goal.length <= 100
s and goal consist of lowercase English letters.
*/

import (
	"fmt"
	"strings"
)

func rotateStringBruteForce(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	sRunes := []rune(s)
	tmp := make([]rune, len(sRunes))

	for i := 0; i < len(sRunes); i++ {
		copy(tmp, sRunes)

		for j := 0; j < len(sRunes); j++ {
			if j == len(sRunes)-1 {
				sRunes[0] = tmp[len(sRunes)-1]
				break
			}
			sRunes[j+1] = tmp[j]
		}

		if string(sRunes) == goal {
			return true
		}
	}

	return false
}

func rotateStringOptimized(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strings.Contains(s+s, goal)
}

func main() {
	fmt.Println("Brute-force solution:")
	fmt.Println("Expected: true,  Got:", rotateStringBruteForce("abcde", "cdeab"))
	fmt.Println("Expected: false, Got:", rotateStringBruteForce("abcde", "abced"))

	fmt.Println()

	fmt.Println("Optimized solution:")
	fmt.Println("Expected: true,  Got:", rotateStringOptimized("abcde", "cdeab"))
	fmt.Println("Expected: false, Got:", rotateStringOptimized("abcde", "abced"))
}
