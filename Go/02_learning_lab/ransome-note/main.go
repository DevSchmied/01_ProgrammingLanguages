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
