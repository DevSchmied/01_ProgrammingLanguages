package main

import "fmt"

func main() {
	// Task 1. What will be printed to the console?
	// Explain how string iteration works in Go when using range,
	// especially with UTF-8 encoded characters (runes).

	fmt.Println("Task 1")

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

	// Task 2. Create a small concurrent program in Go that demonstrates a data race
	// when multiple goroutines access and modify shared state without proper synchronization.

	fmt.Println("Task 2")

	m := map[int]*int{}
	val := 0
	m[0] = &val

	for i := 0; i < 30; i++ {
		go func(i int) {
			cur := m[0]
			*cur += i
		}(i)
	}

	fmt.Println(*m[0])
}
