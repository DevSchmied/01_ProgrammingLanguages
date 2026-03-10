package main

import "fmt"

/*
LEETCODE 349: Intersection of Two Arrays

Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.

Constraints:
1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000

Source: LeetCode
*/

func intersection(nums1 []int, nums2 []int) []int {

	var mapArr map[int]struct{}
	var sliceInt []int

	if len(nums1) > len(nums2) {
		mapArr = make(map[int]struct{}, len(nums1))
		sliceInt = make([]int, 0, len(nums1))

		for _, val := range nums1 {
			if _, exists := mapArr[val]; !exists {
				mapArr[val] = struct{}{}
			}
		}

		for _, val := range nums2 {
			if _, exists := mapArr[val]; exists {
				sliceInt = append(sliceInt, val)
				delete(mapArr, val)
			}
		}

	} else {
		mapArr = make(map[int]struct{}, len(nums2))
		sliceInt = make([]int, 0, len(nums1))

		for _, val := range nums2 {
			if _, exists := mapArr[val]; !exists {
				mapArr[val] = struct{}{}
			}
		}

		for _, val := range nums1 {
			if _, exists := mapArr[val]; exists {
				sliceInt = append(sliceInt, val)
				delete(mapArr, val)
			}
		}
	}

	return sliceInt
}

func main() {

	// Testcase 1
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Printf("nums1 = %v, nums2 = %v\n", nums1, nums2)
	fmt.Printf("Intersection: %v\n\n", intersection(nums1, nums2))

	// Testcase 1
	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}
	fmt.Printf("nums1 = %v, nums2 = %v\n", nums3, nums4)
	fmt.Printf("Intersection: %v\n\n", intersection(nums3, nums4))
}
