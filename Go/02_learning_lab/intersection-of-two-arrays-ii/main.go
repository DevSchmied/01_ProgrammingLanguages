package main

/*

LeetCode 350: Intersection of Two Arrays II

Source: LeetCode

Problem Statement:
- Given two integer arrays nums1 and nums2, return an array of their intersection.
- Each element in the result must appear as many times as it shows in both arrays, and you may return the result in any order.

Examples:
- Input: nums1 = [1,2,2,1], nums2 = [2,2]
- Output: [2,2]

- Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
- Output: [4,9] (or [9,4])

Constraints:
- 1 <= nums1.length, nums2.length <= 1000
- 0 <= nums1[i], nums2[i] <= 1000

*/

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	counts := make(map[int]int)
	res := make([]int, 0, len(nums1))

	for _, v := range nums1 {
		counts[v]++
	}

	for _, v := range nums2 {
		if counts[v] > 0 {
			res = append(res, v)
			counts[v]--
		}
	}

	return res
}
