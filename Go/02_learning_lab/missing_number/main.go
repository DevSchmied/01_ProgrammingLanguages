package main

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
