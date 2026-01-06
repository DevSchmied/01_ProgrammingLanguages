package main

/*
Task: Shuffle String

Source: LeetCode (Problem 1528 — Shuffle String)

Description:
You are given a string s and an integer array indices of the same length.
The string s is shuffled such that the character at position i moves to
indices[i] in the shuffled string.

Your task is to return the resulting shuffled string.

Constraints:
- s.length == indices.length == n
- 1 <= n <= 100
- s consists only of lowercase English letters
- 0 <= indices[i] < n
- All values in indices are unique

Examples:

Input:
s = "codeleet"
indices = [4,5,6,7,0,2,1,3]

Output:
"leetcode"

Input:
s = "abc"
indices = [0,1,2]

Output:
"abc"
*/
import "fmt"

func restoreStringBySearch(s string, indices []int) string {
	if len(s) != len(indices) {
		return ""
	}

	resRunes := make([]rune, 0, len(s))

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if indices[j] == i {
				resRunes = append(resRunes, rune(s[j]))
				break
			}
		}
	}

	return string(resRunes)
}

func main() {
	fmt.Println("=======================1. Solution=======================")
	fmt.Println(restoreStringBySearch("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
	fmt.Println(restoreStringBySearch("abc", []int{0, 1, 2}))
}
