package main

import (
	"fmt"
	"math"
)

/*
219. Contains Duplicate II
Difficulty: Easy

Source: GitHub (LeetCode problem #219)

Description:
Given an integer array nums and an integer k,
return true if there are two distinct indices i and j
such that:

    nums[i] == nums[j]
    and abs(i - j) <= k

Otherwise, return false.

Examples:
Input: nums = [1,2,3,1], k = 3
Output: true

Input: nums = [1,0,1,1], k = 1
Output: true

Input: nums = [1,2,3,1,2,3], k = 2
Output: false

Constraints:
1 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
0 <= k <= 10^5
*/

func main() {
	// Example 1
	nums1 := []int{1, 2, 3, 1}
	k1 := 3
	result1 := containsNearbyDuplicate(nums1, k1)
	fmt.Printf("Test 1 - Input: nums = %v, k = %d\n", nums1, k1)
	fmt.Printf("Expected: true | Got: %v\n\n", result1)

	// Example 2
	nums2 := []int{1, 0, 1, 1}
	k2 := 1
	result2 := containsNearbyDuplicate(nums2, k2)
	fmt.Printf("Test 2 - Input: nums = %v, k = %d\n", nums2, k2)
	fmt.Printf("Expected: true | Got: %v\n\n", result2)

	// Example 3
	nums3 := []int{1, 2, 3, 1, 2, 3}
	k3 := 2
	result3 := containsNearbyDuplicate(nums3, k3)
	fmt.Printf("Test 3 - Input: nums = %v, k = %d\n", nums3, k3)
	fmt.Printf("Expected: false | Got: %v\n", result3)
}

func containsNearbyDuplicate(nums []int, k int) bool {

	lastIndex := make(map[int]int, len(nums))

	for i, v := range nums {
		if idx, exists := lastIndex[v]; exists {
			if math.Abs(float64(idx-i)) <= float64(k) {
				return true
			}
		}
		lastIndex[v] = i
	}
	return false
}
