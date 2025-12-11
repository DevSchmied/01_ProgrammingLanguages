package main

/* Task. Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

=============

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

=============

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

=============

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

=============

Constraints:

2 <= nums.length <= 104
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
Only one valid answer exists.
*/

import "fmt"

func twoSum(nums []int, target int) []int {

	for idx1, num1 := range nums {
		for i := 1 + idx1; i < len(nums); i++ {
			if (num1 + nums[i]) == target {
				return []int{idx1, i}
			}
		}
	}
	return nil
}

func twoSumMapSolution(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		need := target - num

		if idx, found := m[need]; found {
			return []int{i, idx}
		}

		m[num] = i
	}

	return nil
}

func main() {

	fmt.Println("\n==========================Solution 1==========================\n")
	fmt.Println("=============Example 1=============")
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println(twoSum(nums1, target1))

	fmt.Println("=============Example 2=============")
	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println(twoSum(nums2, target2))

	fmt.Println("=============Example 3=============")
	nums3 := []int{3, 3}
	target3 := 6
	fmt.Println(twoSum(nums3, target3))

}
