package main

import "fmt"

/*
205. Isomorphic Strings

Source: LeetCode (https://leetcode.com)

Description:
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving
the order of characters. No two characters may map to the same character, but a character
may map to itself.

Examples:
1. Input: s = "egg", t = "add"
   Output: true
   Explanation: 'e' → 'a', 'g' → 'd'

2. Input: s = "foo", t = "bar"
   Output: false
   Explanation: 'o' needs to map to both 'a' and 'r'

3. Input: s = "paper", t = "title"
   Output: true

Constraints:
- 1 <= s.length <= 5 * 10^4
- t.length == s.length
- s and t consist of any valid ASCII character.

Note: The actual LeetCode problem may have a premium lock icon, but the problem
description and requirements are as stated above.
*/

func main() {
	// Test 1: Beispiel 1
	s1 := "egg"
	t1 := "add"
	fmt.Printf("Test 1: s = %q, t = %q\n", s1, t1)
	fmt.Printf("Erwartet: true\n")
	fmt.Printf("Ergebnis: %v\n\n", isIsomorphic(s1, t1))

	// Test 2: Beispiel 2
	s2 := "foo"
	t2 := "bar"
	fmt.Printf("Test 2: s = %q, t = %q\n", s2, t2)
	fmt.Printf("Erwartet: false\n")
	fmt.Printf("Ergebnis: %v\n\n", isIsomorphic(s2, t2))

	// Test 3: Beispiel 3
	s3 := "paper"
	t3 := "title"
	fmt.Printf("Test 3: s = %q, t = %q\n", s3, t3)
	fmt.Printf("Erwartet: true\n")
	fmt.Printf("Ergebnis: %v\n\n", isIsomorphic(s3, t3))
}

func isIsomorphic(s string, t string) bool {
	mapST := make(map[byte]byte, len(s))
	mapTS := make(map[byte]byte, len(s))

	for i := 0; i < len(s); i++ {
		c1 := s[i]
		c2 := t[i]

		if val, ok := mapST[c1]; ok && val != c2 {
			return false
		}

		if val, ok := mapTS[c2]; ok && val != c1 {
			return false
		}

		mapST[c1] = c2
		mapTS[c2] = c1
		// test
	}

	return true
}
