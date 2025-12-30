package main

/*
Task:

Write a function that receives a value of any type and prints the value’s runtime category along with the value itself.
The function should handle at least: whole numbers, text values, and decimal numbers.
If the value does not belong to these categories, print that the type is unknown.
*/

func printType(x any) {
	switch v := x.(type) {
	case int:
		println("int", v)
	case string:
		println("string", v)
	case float64:
		println("float64", v)
	default:
		println("unknown type", v)
	}
}

func main() {
	printType(1)
	printType("str")
	printType(6.4)
}
