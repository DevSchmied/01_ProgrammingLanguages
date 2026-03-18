package main

import "fmt"

/*
LeetCode 409: Longest Palindrome

Problem Statement:

Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome.

Examples:
Input: s = "abccccdd"
Output: 7
Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.

Input: s = "a"
Output: 1
Explanation: The longest palindrome that can be built is "a", whose length is 1.

Constraints:
1 <= s.length <= 2000
s consists of lowercase and/or uppercase English letters only.

Note: This problem is from LeetCode and requires finding the maximum possible palindrome length using the given characters.
*/

func longestPalindrome(s string) int {

	mapString := make(map[rune]int)
	length := 0
	hasOdd := false
	for _, ch := range s {
		mapString[ch]++
	}

	for _, value := range mapString {
		length += (value / 2) * 2
		if value%2 == 1 {
			hasOdd = true
		}
	}
	if hasOdd {
		length++
	}
	return length
}

func main() {
	// Test case 1: s = "abccccdd" -> longest palindrome length should be 7
	s1 := "abccccdd"
	result1 := longestPalindrome(s1)
	fmt.Printf("Test 1 - s: %s -> Result: %d (Expected: 7)\n", s1, result1)

	// Test case 2: s = "a" -> single character -> should return 1
	s2 := "a"
	result2 := longestPalindrome(s2)
	fmt.Printf("Test 2 - s: %s -> Result: %d (Expected: 1)\n", s2, result2)
}
