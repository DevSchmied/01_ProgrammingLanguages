package main

import "fmt"

/*
LeetCode — Problem 268. Missing Number

Given an array nums containing n distinct numbers in the range [0, n],
return the only number in the range that is missing from the array.

Examples:
Input: nums = [3,0,1]  -> Output: 2
Input: nums = [0,1]    -> Output: 2
Input: nums = [9,6,4,2,3,5,7,0,1] -> Output: 8

Constraints:
- n == nums.length
- 1 <= n <= 10^4
- 0 <= nums[i] <= n
- All numbers in nums are unique.

Follow-up:
Can you implement a solution using only O(1) extra space and O(n) time?

Source: LeetCode
*/

// Solution 1 — Hash Map
func missingNumber(nums []int) int {
	seen := make(map[int]struct{})

	for _, num := range nums {
		seen[num] = struct{}{}
	}

	for i := 0; i <= len(nums); i++ {
		if _, exists := seen[i]; !exists {
			return i
		}
	}
	return -1
}

// Solution 2 — Brute Force
func missingNumberBruteForce(nums []int) int {
	for i := 0; i <= len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == nums[j] {
				break
			}
			if j == len(nums)-1 {
				return i
			}
		}
	}
	return -1
}

func main() {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{1, 2, 3}, 0},
		{[]int{0}, 1},
		{[]int{1}, 0},
	}

	fmt.Println("Testing HashMap solution:")
	for i, test := range tests {
		result := missingNumber(test.nums)
		fmt.Printf("Test %d: nums=%v -> result=%d (expected %d)\n",
			i+1, test.nums, result, test.want)
	}

	fmt.Println()

	fmt.Println("Testing Brute Force solution:")
	for i, test := range tests {
		result := missingNumberBruteForce(test.nums)
		fmt.Printf("Test %d: nums=%v -> result=%d (expected %d)\n",
			i+1, test.nums, result, test.want)
	}
}
