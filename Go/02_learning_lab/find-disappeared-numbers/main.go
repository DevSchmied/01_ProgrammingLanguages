package main

import "fmt"

/*
LeetCode 448: Find All Numbers Disappeared in an Array
Quelle: LeetCode (https://leetcode.com)

Problem Statement:

Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

Examples:
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
Explanation: The numbers 1 through 8 should appear, but 5 and 6 are missing.

Input: nums = [1,1]
Output: [2]
Explanation: The numbers 1 through 2 should appear, but 2 is missing.

Constraints:
n == nums.length
1 <= n <= 10^5
1 <= nums[i] <= n
*/

func findDisappearedNumbers(nums []int) []int {

	res := []int{}
	mapNums := make(map[int]struct{})

	for _, v := range nums {
		mapNums[v] = struct{}{}
	}

	for i := 1; i <= len(nums); i++ {
		if _, exists := mapNums[i]; !exists {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	// Test case 1: Standard case with multiple missing numbers
	nums1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result1 := findDisappearedNumbers(nums1)
	fmt.Printf("Test 1 - nums: %v -> Result: %v (Expected: [5 6])\n", nums1, result1)

	// Test case 2: Case with duplicate numbers
	nums2 := []int{1, 1}
	result2 := findDisappearedNumbers(nums2)
	fmt.Printf("Test 2 - nums: %v -> Result: %v (Expected: [2])\n", nums2, result2)
}
