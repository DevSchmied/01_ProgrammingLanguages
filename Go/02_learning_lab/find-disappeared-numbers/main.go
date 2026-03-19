package main

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
