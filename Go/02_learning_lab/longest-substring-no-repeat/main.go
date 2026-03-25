package main

import "fmt"

/*
LeetCode 3: Longest Substring Without Repeating Characters
Quelle: LeetCode (https://leetcode.com)

Problem Statement:

Given a string s, find the length of the longest substring without repeating characters.

Examples:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
0 <= s.length <= 5 * 10^4
s consists of English letters, digits, symbols and spaces.

*/

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		if idx, exists := seen[s[right]]; exists {
			seen[s[right]] = idx
			if idx >= left {
				left = idx + 1
			}
		}
		seen[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	// Test case 1: s = "abcabcbb" -> longest substring without repeating characters should be 3
	s1 := "abcabcbb"
	result1 := lengthOfLongestSubstring(s1)
	fmt.Printf("Test 1 - s: \"%s\" -> Result: %d (Expected: 3)\n", s1, result1)

	// Test case 2: s = "bbbbb" -> all same characters -> should return 1
	s2 := "bbbbb"
	result2 := lengthOfLongestSubstring(s2)
	fmt.Printf("Test 2 - s: \"%s\" -> Result: %d (Expected: 1)\n", s2, result2)

	// Test case 3: s = "pwwkew" -> longest substring "wke" or "kew" -> should return 3
	s3 := "pwwkew"
	result3 := lengthOfLongestSubstring(s3)
	fmt.Printf("Test 3 - s: \"%s\" -> Result: %d (Expected: 3)\n", s3, result3)
}
