// leetcode
// Given an integer array nums and an integer val, remove all occurrences of val in
// nums in-place. The order of the elements may be changed. Then return the number
// of elements in nums which are not equal to val.

// Consider the number of elements in nums which are not equal to val be k, to get
// accepted, you need to do the following things:

// - Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
// - Return k.

package main

import "fmt"

func main() {
	fmt.Println("----------1. case----------")
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println("k = ", removeElement(nums, val))

	fmt.Println("----------2. case----------")
	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2
	fmt.Println("k = ", removeElement(nums2, val2))

	fmt.Println("----------repeat()----------")
	nums3 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val3 := 2
	fmt.Printf("nums3 before removal: %v\n", nums3)
	k := repeatRemoveElement(nums3, val3)
	fmt.Printf("nums3 after removal: %v\n", nums3[:k])
	fmt.Println("Number of elements (k):", k)

}

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	fmt.Println("nums:", nums)
	return k
}

func repeatRemoveElement(nums []int, val int) int {
	var k = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
