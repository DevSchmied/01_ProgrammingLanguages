package main

import "fmt"

/*
LeetCode 383: Ransom Note

Note: This problem is from LeetCode.

Problem Statement:
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
Each letter in magazine can only be used once in ransomNote.

Examples:
Input: ransomNote = "a", magazine = "b"
Output: false

Input: ransomNote = "aa", magazine = "ab"
Output: false

Input: ransomNote = "aa", magazine = "aab"
Output: true

Constraints:
1 <= ransomNote.length, magazine.length <= 10^5
ransomNote and magazine consist of lowercase English letters only
*/

func canConstruct(ransomNote string, magazine string) bool {

	mapMagazine := make(map[rune]int)

	for _, v := range magazine {
		mapMagazine[v]++
		fmt.Println("v:", string(v))
	}
	fmt.Println(mapMagazine)

	for _, v := range ransomNote {
		fmt.Println("v2:", mapMagazine[v])
		_, exists := mapMagazine[v]
		if exists && mapMagazine[v] != 0 {
			mapMagazine[v]--
		} else {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1: ransomNote = "a", magazine = "b" -> should return false
	ransomNote1 := "a"
	magazine1 := "b"
	result1 := canConstruct(ransomNote1, magazine1)
	fmt.Printf("Test 1 - ransomNote: %s, magazine: %s -> Result: %v (Expected: false)\n",
		ransomNote1, magazine1, result1)

	// Test case 2: ransomNote = "aa", magazine = "ab" -> should return false
	ransomNote2 := "aa"
	magazine2 := "ab"
	result2 := canConstruct(ransomNote2, magazine2)
	fmt.Printf("Test 2 - ransomNote: %s, magazine: %s -> Result: %v (Expected: false)\n",
		ransomNote2, magazine2, result2)

	// Test case 3: ransomNote = "aa", magazine = "aab" -> should return true
	ransomNote3 := "aa"
	magazine3 := "aab"
	result3 := canConstruct(ransomNote3, magazine3)
	fmt.Printf("Test 3 - ransomNote: %s, magazine: %s -> Result: %v (Expected: true)\n",
		ransomNote3, magazine3, result3)
}
