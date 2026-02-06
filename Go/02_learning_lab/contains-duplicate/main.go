package main

import "fmt"

/*
217. Contains Duplicate

Topics: Array, Hash Table, Sorting
Source: LeetCode (https://leetcode.com)

Given an integer array nums, return true if any value appears at least twice
in the array, and return false if every element is distinct.

Example 1:
Input: nums = [1,2,3,1]
Output: true
Explanation: The element 1 occurs at indices 0 and 3.

Example 2:
Input: nums = [1,2,3,4]
Output: false
Explanation: All elements are distinct.

Example 3:
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

Constraints:
- 1 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9
*/
func main() {
	// Test Example 1
	nums1 := []int{1, 2, 3, 1}
	result1 := containsDuplicate(nums1)
	fmt.Printf("Test 1 - Input: %v\n", nums1)
	fmt.Printf("Expected Output: true\n")
	fmt.Printf("Actual Output: %v\n", result1)

	// Test Example 2
	nums2 := []int{1, 2, 3, 4}
	result2 := containsDuplicate(nums2)
	fmt.Printf("Test 2 - Input: %v\n", nums2)
	fmt.Printf("Expected Output: false\n")
	fmt.Printf("Actual Output: %v\n", result2)

	// Test Example 3
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	result3 := containsDuplicate(nums3)
	fmt.Printf("Test 3 - Input: %v\n", nums3)
	fmt.Printf("Expected Output: true\n")
	fmt.Printf("Actual Output: %v\n", result3)

}

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, exists := numsMap[num]; exists {
			return true
		}
		numsMap[num] = struct{}{}

	}
	return false
}
