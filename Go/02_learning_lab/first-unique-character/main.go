package main

import "fmt"

/*

LeetCode 387: First Unique Character in a String

Problem Statement:
- Given a string s, find the first non-repeating character in it and return its index.
- If it does not exist, return -1.

Examples:
Input: s = "leetcode"
Output: 0
Explanation: The character 'l' at index 0 is the first character that does not occur at any other index.

Input: s = "loveleetcode"
Output: 2

Input: s = "aabb"
Output: -1

Constraints:
1 <= s.length <= 10^5
s consists of only lowercase English letters

Note: This problem is from LeetCode and requires finding the first character that appears exactly once in the string.
*/

func firstUniqChar(s string) int {

	mapString := make(map[rune]int)
	for _, ch := range s {
		mapString[ch]++
	}

	for idx, ch := range s {
		if mapString[ch] == 1 {
			return idx
		}
	}

	return -1
}

func main() {
	// Test case 1: s = "leetcode" -> first unique character 'l' at index 0
	s1 := "leetcode"
	result1 := firstUniqChar(s1)
	fmt.Printf("Test 1 - s: %s -> Result: %d (Expected: 0)\n", s1, result1)

	// Test case 2: s = "loveleetcode" -> first unique character 'v' at index 2
	s2 := "loveleetcode"
	result2 := firstUniqChar(s2)
	fmt.Printf("Test 2 - s: %s -> Result: %d (Expected: 2)\n", s2, result2)

	// Test case 3: s = "aabb" -> no unique character -> should return -1
	s3 := "aabb"
	result3 := firstUniqChar(s3)
	fmt.Printf("Test 3 - s: %s -> Result: %d (Expected: -1)\n", s3, result3)
}
