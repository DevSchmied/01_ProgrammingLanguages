package main

import "fmt"

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

	if len(s) == 0 && len(t) == 0 {
		return true
	}

	return backspaceDelete(s) == backspaceDelete(t)
}

func backspaceDelete(s string) string {
	r := make([]rune, len(s))
	counter := 0

	for _, ch := range s {
		if ch != '#' {
			r[counter] = ch
			counter++
		} else if counter > 0 {
			counter--
		}
	}
	return string(r[:counter])
}

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a#c", "b"))
}
