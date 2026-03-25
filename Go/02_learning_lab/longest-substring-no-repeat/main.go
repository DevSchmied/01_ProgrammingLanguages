package main

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

}
