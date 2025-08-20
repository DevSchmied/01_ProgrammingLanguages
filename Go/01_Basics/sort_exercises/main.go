package main

import (
	"fmt"
	"sort"
)

// SORT PACKAGE TASKS

func main() {
	fmt.Println("Hello Go!")

	/* 1
	   Declare a slice of integers numbers containing the elements 5, 2, 9, 1, and 3. Sort this slice in ascending order.
	*/

	fmt.Println("------------------Exercise 1------------------")

	// numbersArr := [5]int{5, 2, 9, 1, 3} // array (with length)
	numbers := []int{5, 2, 9, 1, 3} // slice
	fmt.Println("Before sort:", numbers)
	fmt.Println("Is sorted initially:", sort.IntsAreSorted(numbers))

	sort.Ints(numbers)

	fmt.Println("After sort:", numbers)
	fmt.Println("Is sorted now:", sort.IntsAreSorted(numbers))

	/* 2.
	   Declare a slice of integers numbers containing the elements 1, 2, 3, 4, and 5. Check whether this slice is sorted in ascending order. Save the result in the variable isSorted.
	*/

	fmt.Println("\n------------------Exercise 2------------------")
	numbers = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice numbers:", numbers)
	isNumbersSorted := sort.IntsAreSorted(numbers)
	fmt.Println("Is sorted ascending?", isNumbersSorted)

	/* 3.
	   Declare a slice of strings words containing "banana", "apple", and "cherry". Sort the strings in alphabetical order.
	   Check whether this slice is sorted. Save the result in the variable isSorted.
	*/

	fmt.Println("\n------------------Exercise 3------------------")
	words := []string{"banana", "apple", "cherry"}
	fmt.Println("Slice words before sort:", words)
	sort.Strings(words)
	fmt.Println("Slice words after sort:", words)

	isWordsSorted := sort.StringsAreSorted(words)
	fmt.Println("Is variable words sorted ascending (in alphabetical order)?", isWordsSorted)

	/* 4.
	   Declare a slice of floating-point numbers floats containing 3.1, 2.2, 5.5, and 1.0. Sort the elements in ascending order.

	   Check whether this slice is sorted. Save the result in the variable isSorted.
	*/

	fmt.Println("\n------------------Exercise 4------------------")

	floats := []float64{3.1, 2.2, 5.5, 1.0}
	fmt.Println("slice floats before sort: ", floats)

	sort.Float64s(floats)

	fmt.Println("floats after sort: ", floats)

	isFloatsSorted := sort.Float64sAreSorted(floats)
	fmt.Println("are floats sorted ascending?", isFloatsSorted)

	/* 5.
	   Declare a struct Person with fields Name (string) and Age (integer). Create a slice people with three elements: Alice (30), Bob (25), Charlie (35). Sort the people by age in ascending order. Use an anonymous comparison function.
	*/

	fmt.Println("\n------------------Exercise 5------------------")

	type Person struct {
		Name string
		Age  int
	}

	/*
		person1 := Person{Name: "Alice", Age: 30}
		person2 := Person{Name: "Bob", Age: 25}
		person3 := Person{Name: "Charlie", Age: 35}

		people := []Person{person1, person2, person3}
	*/

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	fmt.Println("variable people before sort: ", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("variable people after sort: ", people)

	/* 6.
	   Declare a slice of integers numbersDesc containing 1, 4, 2, 5, and 3. Sort the elements in descending order. Use a custom comparison function.
	*/

	fmt.Println("\n------------------Exercise 6------------------")

	numbersDesc := []int{1, 4, 2, 5, 3}
	fmt.Println("numbersDesc before sort: ", numbersDesc)

	sort.Slice(numbersDesc, func(i, j int) bool {
		return numbersDesc[i] > numbersDesc[j]
	})

	fmt.Println("numbersDesc after sort, descending: ", numbersDesc)

	/* 7.
	   Declare a slice of strings wordsByLength containing "tree", "go", "algorithm", "hi". Sort the strings by their length (from shortest to longest). Use a custom comparison function.
	*/

	fmt.Println("\n------------------Exercise 7------------------")

	wordsByLength := []string{"tree", "go", "algorithm", "hi"}
	fmt.Println("wordsByLength before sort: ", wordsByLength)

	sort.Slice(wordsByLength, func(i, j int) bool {
		return len(wordsByLength[i]) < len(wordsByLength[j])
	})

	fmt.Println("wordsByLength after sort by their length (from shortest to longest): ", wordsByLength)

	repeat()

}

func repeat() {
	/*
			1. Declare a slice of integers numbers containing the elements 5, 2, 9, 1, and 3. Sort this slice
		in ascending order.
	*/
	fmt.Println("\n\n----------------------repeat()----------------------")

	numbers := []int{5, 2, 9, 1, 3}

	sort.Ints(numbers)

	fmt.Println("numbers: ", numbers)

	/*
	   2. Declare a slice of integers numbersSorted containing the elements 1, 2, 3, 4, and 5.
	   Check whether this slice is sorted in ascending order. Save the result in the variable isSorted.
	*/

	numbersSorted := []int{1, 2, 3, 4, 5}
	isSorted := sort.IntsAreSorted(numbersSorted)

	fmt.Println("isSorted: ", isSorted)
	/*
	   3. Declare a slice of strings words containing "banana", "apple", and "cherry". Sort the strings
	   in alphabetical order. Check whether this slice is sorted. Save the result in the variable areWordsSorted.
	*/
	words := []string{"banana", "apple", "cherry"}
	sort.Strings(words)
	areWordsSorted := sort.StringsAreSorted(words)

	fmt.Println("areWordsSorted: ", areWordsSorted)
	/*
	   4. Declare a slice of floating-point numbers floats containing 3.1, 2.2, 5.5, and 1.0. Sort the
	   elements in ascending order. Check whether this slice is sorted. Save the result in the variable areFloatsSorted.
	*/
	floats := []float64{3.1, 2.2, 5.5, 1.0}
	sort.Float64s(floats)

	areFloatsSorted := sort.Float64sAreSorted(floats)

	fmt.Println("areFloatsSorted: ", areFloatsSorted)
	/*
	   5. Declare a struct Person with fields Name (string) and Age (integer). Create a slice people
	   with three elements: Alice (30), Bob (25), Charlie (35). Sort the people by age in ascending order.
	   Use an anonymous comparison function.
	*/
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("people: ", people)
	/*
	   6. Declare a slice of integers numbersDesc containing 1, 4, 2, 5, and 3. Sort the elements in
	   descending order. Use a custom comparison function.
	*/
	numbersDesc := []int{1, 4, 2, 5, 3}

	sort.Slice(numbersDesc, func(i, j int) bool {
		return numbersDesc[i] > numbersDesc[j]
	})

	fmt.Println("numbersDesc: ", numbersDesc)
	/*
	   7. Declare a slice of strings wordsByLength containing "tree", "go", "algorithm", "hi". Sort the
	   strings by their length (from shortest to longest). Use a custom comparison function.
	*/
	wordsByLength := []string{"tree", "go", "algorithm", "hi"}

	sort.Slice(wordsByLength, func(i, j int) bool {
		return len(wordsByLength[i]) < len(wordsByLength[j])
	})

	fmt.Println(wordsByLength)

}
