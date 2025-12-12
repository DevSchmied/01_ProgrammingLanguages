package main

import "fmt"

func main() {
	// Task: What will be printed to the console?
	// Explain how string iteration works in Go when using range,
	// especially with UTF-8 encoded characters (runes).

	str1 := "hello 🙂💪"
	str2 := "привет 🙂💪"

	for i, char := range str1 {
		fmt.Println(i, ": ", char)
	}

	fmt.Println()
	for i, char := range str1 {
		fmt.Printf("%d: %c\n", i, char)
	}

	fmt.Println()
	fmt.Println()

	for i, char := range str2 {
		fmt.Println(i, ": ", char)
	}

	fmt.Println()
	for i, char := range str2 {
		fmt.Printf("%d: %c\n", i, char)
	}

}
