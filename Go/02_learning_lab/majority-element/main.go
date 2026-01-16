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

	valuesMap := make(map[int]int)

	for _, v := range nums {
		valuesMap[v]++
	}

	counter := 0
	result := nums[0]
	for k, v := range valuesMap {
		if counter < v {
			counter = v
			result = k
		}
	}
	if counter > len(nums)/2 {
		return result
	}
	return 0
}

func majorityElementOptimized(nums []int) int {
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

	fmt.Println("======================================")
	fmt.Println("Solution 1: Hash Map")
	fmt.Println("======================================")

	fmt.Printf("Example 1 | Expected: %d | Got: %d\n", 3,
		majorityElement([]int{3, 2, 3}))
	fmt.Println()

	fmt.Printf("Example 2 | Expected: %d | Got: %d\n", 2,
		majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println()

	fmt.Printf("Example 3 | Expected: %d | Got: %d\n", 1,
		majorityElement([]int{2, 2, 1, 1, 1, 1, 1, 3, 3}))
	fmt.Println()

	fmt.Println("======================================")
	fmt.Println("Solution 2, optimized")
	fmt.Println("======================================")

	fmt.Printf("Example 1 | Expected: %d | Got: %d\n", 3,
		majorityElementOptimized([]int{3, 2, 3}))
	fmt.Println()

	fmt.Printf("Example 2 | Expected: %d | Got: %d\n", 2,
		majorityElementOptimized([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println()

	fmt.Printf("Example 3 | Expected: %d | Got: %d\n", 1,
		majorityElementOptimized([]int{2, 2, 1, 1, 1, 1, 1, 3, 3}))
}
