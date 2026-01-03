package main

import "fmt"

/*
	SlidCode Problem 344 — Reverse String

	Topic: Two Pointers, String Manipulation

	Description:
	Write a function that reverses a string.
	The input string is given as an array of characters s.

	You must do this by modifying the input array in-place
	with O(1) extra memory.

	Examples:

	Example 1:
	Input:  s = ["h","e","l","l","o"]
	Output: ["o","l","l","e","h"]

	Example 2:
	Input:  s = ["H","a","n","n","a","h"]
	Output: ["h","a","n","n","a","H"]

	Constraints:
	- 1 <= s.length <= 10^5
	- s[i] is a printable ASCII character

	Note:
	This is a SlidCode practice problem.
*/

func reverseString(s []byte) {
	length := len(s)

	for i := 0; i < length/2; i++ {
		tmp := s[i]
		s[i] = s[length-i-1]
		s[length-i-1] = tmp
	}
}

func main() {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}

	reverseString(s1)
	reverseString(s2)

	fmt.Println(string(s1))
	fmt.Println(string(s2))
}
