package main

/*
Source: LeetCode — Problem 242. Valid Anagram

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Constraints:
1 <= s.length, t.length <= 5 * 10^4
s and t consist of lowercase English letters.
*/

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	length := len(s)
	if len(t) > length {
		length = len(t)
	}

	mapS := make(map[rune]int, length)
	mapT := make(map[rune]int, length)

	for _, v := range s {
		mapS[v] = mapS[v] + 1
	}

	for _, v := range t {
		mapT[v] = mapT[v] + 1
	}

	for k, valS := range mapS {
		valT, exists := mapT[k]
		if !exists || valT != valS {
			return false
		}
	}
	return true
}

func main() {
	// Test 1
	fmt.Println("Test 1:", isAnagram("anagram", "nagaram"), "expected: true")

	// Test 2
	fmt.Println("Test 2:", isAnagram("rat", "car"), "expected: false")
}
