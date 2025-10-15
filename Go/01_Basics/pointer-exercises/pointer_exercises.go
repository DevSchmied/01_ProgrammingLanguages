package main

import "fmt"

// Task 3
type Person struct {
	ID   int
	Name string
}

func main() {

	/*
		Task 1 — Integer Pointer
			- Declare an integer variable.
			- Declare a pointer variable to an integer.
			- Assign a value to the integer variable.
			- Assign the address of the integer variable to the pointer.
			- Print the value of the integer variable.
			- Print the address of the integer variable.
			- Print the value of the pointer variable.
			- Print the value stored at the address the pointer points to.
	*/
	fmt.Println("---------------------Task 1---------------------")

	var num int
	var numPtr *int

	num = 42
	numPtr = &num

	fmt.Printf("Value of num: %d\n", num)
	fmt.Printf("Address of num: %p\n", &num)
	fmt.Printf("Value of numPtr (address stored): %p\n", numPtr)
	fmt.Printf("Value pointed to by numPtr: %d\n", *numPtr)

	/*
		Task 2 — String Pointer
			- Declare a string variable.
			- Declare a pointer variable to a string.
			- Assign a text value to the string variable.
			- Assign the address of the string variable to the pointer.
			- Print the value of the string variable.
			- Print the address of the string variable.
			- Print the value of the pointer variable.
			- Print the value stored at the address the pointer points to.
	*/

	fmt.Println()
	fmt.Println("---------------------Task 2---------------------")

	var str string
	var strPtr *string

	str = "Text"
	strPtr = &str

	fmt.Printf("Value of str: %s\n", str)
	fmt.Printf("Address of str: %p\n", &str)

	fmt.Printf("Value of strPtr (stored address): %p\n", strPtr)
	fmt.Printf("Value pointed to by strPtr: %s\n", *strPtr)

	/*
		Task 3 — Struct Pointer
		-	Define a simple struct with at least two fields.
		-	Declare a variable of that struct type.
		-	Declare a pointer variable to that struct type.
		-	Assign values to the struct fields.
		-	Assign the address of the struct variable to the pointer.
		-	Print the struct variable.
		-	Print the address of the struct variable.
		-	Print the pointer variable.
		-	Print the value stored at the address the pointer points to.
	*/

	fmt.Println()
	fmt.Println("---------------------Task 3---------------------")

	fmt.Println()
	fmt.Println("--------------------- Task 3 ---------------------")

	var person1 Person
	var person1Ptr *Person

	person1.ID = 1
	person1.Name = "Max Mustermann"
	person1Ptr = &person1

	fmt.Printf("Struct variable: %v\n", person1)
	fmt.Printf("Address of struct: %p\n", &person1)
	fmt.Printf("Pointer variable: %p\n", person1Ptr)
	fmt.Printf("Value pointed to by pointer: %v\n", *person1Ptr)
}
