package main

import "fmt"

/*
LeetCode — Palindrome Number

Quelle:
https://leetcode.com/problems/palindrome-number/


Task. Palindrome Number

Given an integer x, return true if x is a palindrome, and false otherwise.

==============

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

==============

Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

==============

Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	nums := []int{}
	modulo := 10
	for x != 0 {
		rest := x % modulo
		nums = append(nums, rest/(modulo/10))
		x -= rest
		modulo *= 10
	}
	if len(nums) == 1 {
		return true
	}

	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("============1. Example============")
	fmt.Println(isPalindrome(121))

	fmt.Println("============2. Example============")
	fmt.Println(isPalindrome(-121))

	fmt.Println("============3. Example============")
	fmt.Println(isPalindrome(10))
}
