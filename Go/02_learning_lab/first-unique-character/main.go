package main

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
