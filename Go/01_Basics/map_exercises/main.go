package main

import "fmt"

func printHeader(n int) {
	fmt.Printf("\n------------------------------%d. exercise------------------------------\n", n)
}

func main() {

	// BASIC METHODS OF WORKING WITH MAPS
	fmt.Println("------------------------------Map exercises in Go------------------------------\n")

	// 1. Create an empty map (users) with keys of type string and values of type int.
	printHeader(1)

	users := make(map[string]int)
	fmt.Println("Empty map users: ", users)

	// 2. Add elements to a map (ages[string]int) in two different ways.
	printHeader(2)

	// Way 1: Initialize the map with a literal
	ages := map[string]int{
		"Alice": 25,
	}

	// Way 2: Add elements by assignment
	ages["Charlie"] = 35

	fmt.Println("Ages map:", ages)

	// 3. Create a map (scores) with values "Alice": 90, "Bob": 85, "Charlie": 80.
	//    Remove the element with the key "Bob".
	printHeader(3)

	scores := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 80,
	}

	delete(scores, "Bob")
	fmt.Println("Scores map after deleting Bob:", scores)

	// 4. Create a map (settings) with values [string]bool "darkMode": true, "notifications": false.
	//    Check if the key "darkMode" exists in the map and print the value.
	printHeader(4)

	settings := map[string]bool{
		"darkMode":      true,
		"notifications": false,
	}

	if value, exists := settings["darkMode"]; exists {
		fmt.Println("The value of \"darkMode\":", value)
	} else {
		fmt.Println("The key \"darkMode\" does not exist in settings")
	}

	// 5. Create a map (products) with values [string]float64 "Laptop": 999.99, "Phone": 499.99.
	//    Iterate over all elements and print them.
	printHeader(5)

	products := map[string]float64{
		"Laptop": 999.99,
		"Phone":  499.99,
	}

	for key, value := range products {
		fmt.Printf("Product: %s, Price: %.2f\n", key, value)
	}

	// 6. Create a map (grades) with values [string]int "Math": 90, "Science": 85.
	//    Change the value of "Science" to 95.
	printHeader(6)

	grades := map[string]int{
		"Math":    90,
		"Science": 85,
	}

	grades["Science"] = 95
	fmt.Println("Grades map after updating Science:", grades)
}
