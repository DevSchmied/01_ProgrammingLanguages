package main

import "fmt"

/*
LeetCode 169 — Majority Element

Problem Description:
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times.
You may assume that the majority element always exists in the array.

Examples:

Example 1:
Input:  nums = [3, 2, 3]
Output: 3

Example 2:
Input:  nums = [2, 2, 1, 1, 1, 2, 2]
Output: 2

Constraints:
- n == nums.length
- 1 <= n <= 5 * 10^4
- -10^9 <= nums[i] <= 10^9
- The input is generated such that a majority element will always exist.

Source:
This is a problem from the LeetCode platform.
*/

func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {

	fmt.Printf("Example 1 | Expected: %d | Got: %d\n", 3, majorityElement([]int{3, 2, 3}))

	fmt.Println()

	fmt.Printf("Example 2 | Expected: %d | Got: %d\n", 2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
