package main

import "fmt"

/*
LeetCode #541 — Reverse String II (Easy)

Problem statement:
Given a string s and an integer k, reverse the first k characters for every 2k characters
counting from the start of the string.

Rules:
- If there are fewer than k characters left, reverse all of them.
- If there are fewer than 2k but greater than or equal to k characters left,
  reverse the first k characters and leave the remaining characters unchanged.

Examples:
Input:  s = "abcdefg", k = 2
Output: "bacdfeg"

Input:  s = "abcd", k = 2
Output: "bacd"

Constraints:
- 1 <= s.length <= 10^4
- s consists of only lowercase English letters
- 1 <= k <= 10^4
*/

func reverseStr(s string, k int) string {
	r := []rune(s)
	n := len(r)

	for start := 0; start < n; start += 2 * k {
		left := start
		right := start + k - 1

		if right >= n {
			right = n - 1
		}

		for left < right {
			r[left], r[right] = r[right], r[left]
			left++
			right--
		}
	}

	return string(r)
}

func main() {
	fmt.Println("1. case")
	fmt.Println(reverseStr("abcdefg", 2))

	fmt.Println("2. case")
	fmt.Println(reverseStr("abcd", 2))
}
