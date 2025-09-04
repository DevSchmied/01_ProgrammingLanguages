package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ------------------------------------------------------------
// This file contains a set of Go practice questions
// with examples where each task is implemented in one file.
// Variable names are unique across tasks to avoid conflicts.
// Topics covered: maps, strings, runes, conversions,
// math operations, rounding, regex, and user input.
// ------------------------------------------------------------

func printHeader(n int) {
	fmt.Printf("--------------------%2d. exercise--------------------\n", n)
}

func main() {

	// 1. Declare a map named brMap where both keys and values are of type rune.
	//    Add pairs of brackets ( → ), [ → ], { → } in three different ways.
	//    Use unique variable names: brMap1, brMap2, brMap3.

	printHeader(1)
	// 1st way
	brMap1 := make(map[rune]rune)
	brMap1['('] = ')'
	brMap1['['] = ']'
	brMap1['{'] = '}'
	fmt.Println("brMap1:", brMap1)

	// 2nd way
	brMap2 := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	fmt.Println("brMap2:", brMap2)

	// 3rd way
	brMap3 := map[rune]rune{}
	brMap3['('] = ')'
	brMap3['['] = ']'
	brMap3['{'] = '}'
	fmt.Println("brMap3:", brMap3)

	// 2. Declare an integer variable b equal to 42 and convert it
	//    to a string in two different ways. Use variables b1/str1 and b2/str2.

	printHeader(2)

	// 1st way
	b1 := 42
	str1 := fmt.Sprintf("%d", b1)
	fmt.Printf("str1 has type %T and value %s\n", str1, str1)

	// 2nd way
	b2 := 42
	str2 := strconv.Itoa(b2)
	fmt.Printf("str2 has type %T and value %s\n", str2, str2)

	// 3. Split a string s1 containing several words
	//    (e.g., "Go is very cool") into an array of words sArray1 by spaces.

	printHeader(3)
	s1 := "Go is very cool"
	sArray1 := strings.Split(s1, " ")
	fmt.Println("The string s1 split into words:", sArray1)

	// 4. Remove all spaces from the string s2 = "G o p r a c t i c e"
	//    and reassign the result back to s2.

	printHeader(4)
	s2 := "G o p r a c t i c e"
	s2 = strings.ReplaceAll(s2, " ", "")
	fmt.Println("s2 without spaces:", s2)

	// 5. Convert the string s3 = "GoLang IS AwEsoMe"
	//    to lowercase and reassign the result back to s3.

	printHeader(5)
	s3 := "GoLang IS AwEsoMe"
	s3 = strings.ToLower(s3)
	fmt.Println("s3 in lowercase:", s3)

	// 6. Create a regular expression that matches any character
	//    except lowercase Latin letters (a-z) and remove such characters
	//    from the string s4 = "Go!@#123lang" in two different ways.

	printHeader(6)
	s4 := "Go!@#123lang"
	// 1st way: regex
	re := regexp.MustCompile(`[^a-z]`)
	s4Clean1 := re.ReplaceAllString(s4, "")
	fmt.Println("s4 cleaned with regex:", s4Clean1)

	// 2nd way: manual filtering
	var builder strings.Builder
	for _, char := range s4 {
		if char >= 'a' && char <= 'z' {
			builder.WriteRune(char)
		}
	}
	s4Clean2 := builder.String()
	fmt.Println("s4 cleaned manually:", s4Clean2)

	// 7. Remove the character at index j from the string mag1 = "practice".
	//    For example, remove the character at index 2.

	printHeader(7)
	j := 2
	mag1 := "practice"
	magR := []rune(mag1)
	magR = append(magR[:j], magR[j+1:]...)
	mag1 = string(magR)
	fmt.Println("mag1 without character at index 2:", mag1)

	// 8. Convert the string s5 = "golang" to a rune slice sRune1.

	printHeader(8)
	s5 := "golang"
	sRune1 := []rune(s5)
	fmt.Println("s5 as rune slice:", sRune1)
	for idx, r := range sRune1 {
		fmt.Printf("index: %d, rune: %c\n", idx, r)
	}

	// 9. Create a map named intMap1 where both keys and values are of type int.

	printHeader(9)
	intMap1 := make(map[int]int)
	intMap2 := map[int]int{}
	fmt.Printf("Empty map intMap1: %v, type: %T\n", intMap1, intMap1)
	fmt.Printf("Empty map intMap2: %v, type: %T\n", intMap2, intMap2)

	// 10. Assign value 7 to the element with key 5 in intMap1.

	printHeader(10)
	intMap1[5] = 7
	fmt.Println("intMap1 after assignment:", intMap1)

	// 11. Check if intMap1 contains a value for key 5,
	//     and if it does, print it.

	printHeader(11)
	if val, exists := intMap1[5]; exists {
		fmt.Println("Value at key 5 in intMap1:", val)
	} else {
		fmt.Println("Key 5 not found in intMap1")
	}

	// 12. Assign the value of nums1[0] to a string variable tempAnf1,
	//     where nums1 is a slice []int{0, 1, 2, 4, 5, 7}.

	printHeader(12)
	var tempAnf1 string
	nums1 := []int{0, 1, 2, 4, 5, 7}
	tempAnf1 = strconv.Itoa(nums1[0])
	fmt.Println("tempAnf1 contains nums1[0] as string:", tempAnf1)

	// 13. Declare a variable a1 equal to 144 and calculate its square root.
	//     What should you pay attention to?

	printHeader(13)
	var a1 float64 // must be float64 for math.Sqrt
	a1 = 144
	sqrt := math.Sqrt(a1)
	fmt.Printf("Square root of %.0f is %.1f (type: %T)\n", a1, sqrt, sqrt)

	// 14. Round the variable a2 = 7.65 in Go:
	//     - up (ceil),
	//     - down (floor),
	//     - to the nearest integer (round).

	printHeader(14)
	a2 := 7.65
	fmt.Println("Rounded up (ceil):", math.Ceil(a2))
	fmt.Println("Rounded down (floor):", math.Floor(a2))
	fmt.Println("Rounded (nearest):", math.Round(a2))

	// 15. Write a Go program that asks the user for a string input
	//     and saves it to the variable action1.

	printHeader(15)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter some text: ")
	action1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading input:", err)
	}
	action1 = strings.TrimSpace(action1)
	fmt.Println("You entered:", action1)
}
