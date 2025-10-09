package main

import (
	"errors"
	"fmt"
	"sync"
)

/*
Goal

In this exercise, you will practice how to create, wrap, check, and unwrap errors in Go.
You will also learn how to define a custom error type and extract it using errors.As().

Step 1 — Create and return basic errors
Write a function:
- func readFile(name string) (string, error)
 * If name is an empty string (""), return an error using errors.New("filename is missing").
 * If the file name is "config.txt", return the string "file loaded successfully" and nil.
 * Otherwise, return a new error using fmt.Errorf("file %s not found", name).

Step 2 — Wrap an error using fmt.Errorf()
Create a new function:
    func loadConfig(name string) error

Inside this function:
- Call readFile(name).
- If an error is returned, wrap it with a new message using:
      fmt.Errorf("failed to load configuration: %w", err)
- Return the wrapped error.
- If no error occurs, return nil.

Step 3 — Compare wrapped errors
In your main() function:
- Call loadConfig().
- If an error is returned, check whether it’s the “filename is missing” error.
- If so, print "Please provide a filename!".
- Otherwise, print the error message.

Step 4 — Create and extract a custom error with errors.As()
- Define a custom error type.
- Implement the Error() method so that it returns a message.
- Then, in a separate function:
  • Return a wrapped error.
  • In main(), call the function and extract the custom error type.
  • If extraction is successful, print the error code.

Step 5 — Unwrap a chained error
Demonstrate how to unwrap multiple levels of wrapped errors:
- Create a base error.
- Wrap it multiple times.
- Retrieve the original base error step by step.
- Print each level to show the unwrapping process.

*/

// global error
var ErrFilenameMissing = errors.New("filename is missing")

// Step 1
// readFile simulates reading a file and returns different errors or success messages.
func readFile(name string) (string, error) {
	if name == "" {
		return "", ErrFilenameMissing
	} else if name == "config.txt" {
		return "file loaded successfully", nil
	} else {
		return "", fmt.Errorf("file %s not found", name)
	}
}

// Step 2
// loadConfig calls readFile() and wraps any returned error with additional context.
func loadConfig(name string) error {
	if _, err := readFile(name); err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}
	return nil
}

// Step 4
// Custom error type
type NetworkError struct {
	Code int
}

// The Error() method so that it returns a message
func (nErr *NetworkError) Error() string {
	return fmt.Sprintf("Network error with code %d", nErr.Code)
}

// Function that returns a wrapped custom error
func connectServer() error {
	return fmt.Errorf("Connection failed: %w", &NetworkError{404})
}

func main() {
	var wg sync.WaitGroup

	// Step 1
	// Empty filename -> expect error.
	wg.Add(1)
	go func() {
		defer wg.Done()
		str, err := readFile("")
		fmt.Printf("string: \"%s\", error: %v\n", str, err)
	}()

	// Valid file -> expect success.
	wg.Add(1)
	go func() {
		defer wg.Done()
		str, err := readFile("config.txt")
		fmt.Printf("string: \"%s\", error: %v\n", str, err)
	}()

	// Unknown file -> expect error.
	wg.Add(1)
	go func() {
		defer wg.Done()
		str, err := readFile("foo.txt")
		fmt.Printf("string: \"%s\", error: %v\n", str, err)
	}()

	wg.Wait()

	// Step 2
	// Test 1: Empty filename → expect wrapped error
	err := loadConfig("")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Test 2: Valid file → expect success
	err = loadConfig("config.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Successfully loaded")
	}

	// Test 3: Invalid file → expect wrapped error
	err = loadConfig("name.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 3

	err = loadConfig("")
	if err != nil {
		if errors.Is(err, ErrFilenameMissing) {
			fmt.Println("Please provide a filename!")
		} else {
			fmt.Printf("Error message: %v\n", err)
		}
	}

	// Step 4
	// Call the function that may return a wrapped error
	err = connectServer()

	// Try to extract the custom error type from the error chain
	var netErr *NetworkError
	if errors.As(err, &netErr) {
		// Extraction successful — print the error code
		fmt.Printf("The extraction was successful. Error code: %d\n", netErr.Code)
	} else {
		// Extraction failed — print a fallback message
		fmt.Println("No NetworkError found.")
	}

	// Step 5
	// Create the base error
	baseErr := errors.New("base error")

	// Wrap the base error at level 1
	wrappedLevel1 := fmt.Errorf("wrapped level 1: %w", baseErr)
	fmt.Printf("Wrapped Level 1: %v\n", wrappedLevel1)

	// Wrap the error again at level 2
	wrappedLevel2 := fmt.Errorf("wrapped level 2: %w", wrappedLevel1)
	fmt.Printf("Wrapped Level 2: %v\n", wrappedLevel2)

	// Unwrap the second level
	unwrappedOnce := errors.Unwrap(wrappedLevel2)
	fmt.Printf("After first unwrapping: %v\n", unwrappedOnce)

	// Unwrap again to get the original base error
	unwrappedTwice := errors.Unwrap(unwrappedOnce)
	fmt.Printf("After second unwrapping (base error): %v\n", unwrappedTwice)
}
