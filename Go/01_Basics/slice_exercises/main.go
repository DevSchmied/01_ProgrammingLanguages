package main

import (
	"fmt"
	"sort"
)

func printHeader(n int) {
	fmt.Printf("\n------------------------------%d. exercise------------------------------\n", n)
}

func main() {
	fmt.Println("------------------------------Slice exercises in Go------------------------------\n")

	// 1. Declare a slice nums1 of type int with elements 1, 2, 3, 4.
	//    Determine its length and print the length to the screen.
	printHeader(1)
	nums1 := []int{1, 2, 3, 4}
	lengthNums1 := len(nums1)
	fmt.Println("The length of slice nums1:", lengthNums1)

	// 2. Create a slice nums2 with length 2 and capacity 5.
	//    Determine its capacity and print it to the screen.
	printHeader(2)
	nums2 := make([]int, 2, 5)
	capNums2 := cap(nums2)
	fmt.Println("The capacity of slice nums2:", capNums2)

	// 3. Declare a slice nums3 with elements 1, 2.
	//    Append 3 and 4 to it. Print the result.
	printHeader(3)
	nums3 := []int{1, 2}
	nums3 = append(nums3, 3, 4)
	fmt.Println("The slice nums3:", nums3)

	// 4. Create a slice src with elements 1, 2, 3.
	//    Create an empty slice dest.
	//    Copy src into dest in two ways: using copy and using append.
	//    Print both results.
	printHeader(4)
	src := []int{1, 2, 3}
	// way 1: copy
	dest1 := make([]int, len(src))
	copy(dest1, src)
	fmt.Println("dest1 (copy):", dest1)

	// way 2: append
	dest2 := []int{}
	dest2 = append(dest2, src...)
	fmt.Println("dest2 (append):", dest2)

	// 5. Declare a slice nums4 with elements 10, 20, 30, 40, 50.
	//    Create a new slice part containing elements from index 1 to 3 (inclusive).
	//    Print part.
	printHeader(5)
	nums4 := []int{10, 20, 30, 40, 50}
	part := nums4[1:4]
	fmt.Println("The slice part:", part)

	// 6. Declare a slice nums5 with elements 5, 2, 8, 1.
	//    Sort it and print the result.
	printHeader(6)
	nums5 := []int{5, 2, 8, 1}
	sort.Slice(nums5, func(i, j int) bool { return nums5[i] < nums5[j] })
	fmt.Println("The slice nums5 sorted:", nums5)

	// 7. Declare a slice nums6 with elements 10, 20, 30.
	//    Check whether 20 is contained in the slice and print the result.
	printHeader(7)
	nums6 := []int{10, 20, 30}
	contains := false
	for _, v := range nums6 {
		if v == 20 {
			contains = true
			break
		}
	}
	fmt.Println("Does the slice nums6 contain 20?", contains)

	// 8. Declare a slice nums7 with elements 3, 6, 9.
	//    Find the index of element 6 and print it.
	//    If element 5 is not in the slice, print -1.
	printHeader(8)
	nums7 := []int{3, 6, 9}
	idx := -1
	for i, v := range nums7 {
		if v == 6 {
			idx = i
			break
		}
	}
	fmt.Println("The index of the element 6:", idx)

	// 9. Declare a slice nums8 with elements 1, 2, 3, 4.
	//    Reverse it and print the result.
	printHeader(9)
	nums8 := []int{1, 2, 3, 4}
	for i, j := 0, len(nums8)-1; i < j; i, j = i+1, j-1 {
		nums8[i], nums8[j] = nums8[j], nums8[i]
	}
	fmt.Println("The slice nums8 after reverse:", nums8)

	// 10. Declare a slice nums9 with elements 10, 20, 30, 40.
	//     Remove the element at index 1 and print the result.
	printHeader(10)
	nums9 := []int{10, 20, 30, 40}
	idx = 1
	nums9 = append(nums9[:idx], nums9[idx+1:]...)
	fmt.Println("The slice nums9 after removing index 1:", nums9)

	// 11. Write Go code that removes the last element from a slice stack
	//     using the slicing operation.
	printHeader(11)
	stack := []int{1, 2, 3, 4}
	stack = stack[:len(stack)-1]
	fmt.Println("The stack after removing last element:", stack)

	repeat()
}
func repeat() {
	fmt.Println("\n\n------------------------------Repeat slice exercises in Go------------------------------\n")

	// 1. Declare a slice nums1 of type int with elements 1, 2, 3, 4.
	//    Determine its length and print the length to the screen.

	printHeader(1)
	nums1 := []int{1, 2, 3, 4}
	fmt.Println("Length of nums1:", len(nums1))

	// 2. Create a slice nums2 with length 2 and capacity 5.
	//    Determine its capacity and print it to the screen.

	printHeader(2)
	nums2 := make([]int, 2, 5)
	fmt.Println("Capacity of nums2:", cap(nums2))

	// 3. Declare a slice nums3 with elements 1, 2.
	//    Append 3 and 4 to it. Print the result.

	printHeader(3)
	nums3 := []int{1, 2}
	nums3 = append(nums3, 3, 4)
	fmt.Println("The slice nums3:", nums3)

	// 4. Create a slice src with elements 1, 2, 3.
	//    Create an empty slice dest.
	//    Copy src into dest in two ways: using copy and using append.
	//    Print both results.

	printHeader(4)
	src := []int{1, 2, 3}

	// 1. Using copy
	dest := make([]int, len(src))
	copy(dest, src)
	fmt.Println("The slice dest (copy):", dest)

	// 2. Using append
	dest = []int{}
	dest = append(dest, src...)
	fmt.Println("The slice dest (append):", dest)

	// 5. Declare a slice nums4 with elements 10, 20, 30, 40, 50.
	//    Create a new slice part containing elements from index 1 to 3 (inclusive).
	//    Print part.

	printHeader(5)
	nums4 := []int{10, 20, 30, 40, 50}
	part := nums4[1:4]
	fmt.Println("The slice part:", part)

	// 6. Declare a slice nums5 with elements 5, 2, 8, 1.
	//    Sort it and print the result.

	printHeader(6)
	nums5 := []int{5, 2, 8, 1}
	sort.Slice(nums5, func(i, j int) bool {
		return nums5[i] < nums5[j]
	})
	fmt.Println("The slice nums5 sorted:", nums5)

	// 7. Declare a slice nums6 with elements 10, 20, 30.
	//    Check whether 20 is contained in the slice and print the result.

	printHeader(7)
	nums6 := []int{10, 20, 30}
	number := 20
	found := false
	for idx, v := range nums6 {
		if v == number {
			fmt.Printf("%d is contained in the slice at index %d.\n", number, idx)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("%d is NOT contained in the slice.\n", number)
	}

	// 8. Declare a slice nums7 with elements 3, 6, 9.
	//    Find the index of element 6 and print it.
	//    If element 5 is not in the slice, print -1.

	printHeader(8)
	nums7 := []int{3, 6, 9}
	element := 6
	found = false
	for idx, v := range nums7 {
		if v == element {
			fmt.Printf("The index of the element %d is %d.\n", element, idx)
			found = true
			break
		}
	}
	if !found {
		fmt.Println(-1)
	}

	// 9. Declare a slice nums8 with elements 1, 2, 3, 4.
	//    Reverse it and print the result.

	printHeader(9)
	nums8 := []int{1, 2, 3, 4}
	length := len(nums8)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		nums8[i], nums8[j] = nums8[j], nums8[i]
	}
	fmt.Println("The slice nums8 reversed:", nums8)

	// 10. Declare a slice nums9 with elements 10, 20, 30, 40.
	//     Remove the element at index 1 and print the result.

	printHeader(10)
	nums9 := []int{10, 20, 30, 40}
	index := 1
	nums9 = append(nums9[:index], nums9[index+1:]...)
	fmt.Println("The slice nums9:", nums9)

	// 11. Write Go code that removes the last element from a slice stack
	//     using the slicing operation.

	printHeader(11)
	lgth := len(nums9)
	nums9 = nums9[:lgth-1]
	fmt.Println("The slice nums9:", nums9)
}
