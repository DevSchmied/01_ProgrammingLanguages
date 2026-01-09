package main

/*
LeetCode 844 — Backspace String Compare

Problem Description:
Given two strings s and t, return true if they are equal when both are typed
into empty text editors. The character '#' represents a backspace.

When a backspace is applied to an empty text editor, the text remains empty.

Examples:

Example 1:
Input:  s = "ab#c", t = "ad#c"
Output: true
Explanation:
Both strings are processed as follows:
- "ab#c" → "ac"
- "ad#c" → "ac"

Example 2:
Input:  s = "ab##", t = "c#d#"
Output: true
Explanation:
Both strings result in an empty string "".

Example 3:
Input:  s = "a#c", t = "b"
Output: false
Explanation:
- "a#c" → "c"
- "b"   → "b"

Constraints:
- 1 ≤ s.length, t.length ≤ 200
- s and t consist only of lowercase English letters and the '#' character

Follow-up:
Can you solve this problem in O(n) time and O(1) extra space?

Source:
LeetCode (https://leetcode.com/problems/backspace-string-compare/)
*/

func backspaceCompare(s string, t string) bool {

	return true
}
