package main

import "fmt"

/*
LeetCode 49: Group Anagrams
Quelle: LeetCode (https://leetcode.com)

Problem Statement:

Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Examples:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Explanation:
- There is no string in strs that can be rearranged to form "bat".
- The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
- The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.

Input: strs = [""]
Output: [[""]]
Explanation: The empty string is its own anagram.

Input: strs = ["a"]
Output: [["a"]]
Explanation: A single character string is trivially an anagram of itself.

Constraints:
1 <= strs.length <= 10^4
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, len(strs))
	resMap := make(map[string]int)
	idx := 0
	for _, str := range strs {
		key := bubbleSort(str)
		if i, exists := resMap[key]; exists {
			res[i] = append(res[i], str)
			continue
		}
		res[idx] = append(res[idx], str)
		resMap[key] = idx
		idx++
	}
	return res[:idx]
}

func bubbleSort(str string) string {
	result := []rune(str)
	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return string(result)
}

func main() {
	// Test case 1: Standard case with multiple anagram groups
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result1 := groupAnagrams(strs1)
	fmt.Printf("Test 1 - Multiple anagram groups -> Result: %v (Expected: [[bat] [nat tan] [ate eat tea]] or any order)\n", result1)

	// Test case 2: Single empty string
	strs2 := []string{""}
	result2 := groupAnagrams(strs2)
	fmt.Printf("Test 2 - Single empty string -> Result: %v (Expected: [[]])\n", result2)

	// Test case 3: Single character string
	strs3 := []string{"a"}
	result3 := groupAnagrams(strs3)
	fmt.Printf("Test 3 - Single character -> Result: %v (Expected: [[a]])\n", result3)
}
