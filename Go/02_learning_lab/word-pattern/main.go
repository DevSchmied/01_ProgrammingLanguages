package main

import "strings"

/*
Source: LeetCode

290. Word Pattern

Description:
Given a pattern and a string s, determine if s follows the same pattern.

Here "follow" means a full match such that there is a bijection between
a letter in pattern and a non-empty word in s.

Rules:
- Each letter in pattern maps to exactly one unique word in s.
- Each unique word in s maps to exactly one letter in pattern.
- No two letters map to the same word.
- No two words map to the same letter.

Examples:

Example 1:
Input:  pattern = "abba", s = "dog cat cat dog"
Output: true
Explanation:
'a' maps to "dog"
'b' maps to "cat"

Example 2:
Input:  pattern = "abba", s = "dog cat cat fish"
Output: false

Example 3:
Input:  pattern = "aaaa", s = "dog cat cat dog"
Output: false

Constraints:
- 1 <= pattern.length <= 300
- pattern contains only lowercase English letters
- 1 <= s.length <= 3000
- s contains only lowercase English letters and spaces
- s has no leading or trailing spaces
- words in s are separated by a single space
*/

func wordPattern(pattern string, s string) bool {
	letterMap := make(map[rune]string)
	wordMap := make(map[string]rune)
	letters := []rune(pattern)
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	for i := 0; i < len(words); i++ {
		if val, exists := letterMap[letters[i]]; exists {
			if val != words[i] {
				return false
			}
		}
		letterMap[letters[i]] = words[i]

		if val, exists := wordMap[words[i]]; exists {
			if val != letters[i] {
				return false
			}
		}
		wordMap[words[i]] = letters[i]
	}

	return true
}
