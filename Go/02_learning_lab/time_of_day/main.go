package main

/*
TASK — Determining the Time of Day


You are working as a backend developer. To generate text that will be shown to a user living in an arbitrary timezone, the system receives the user’s local time accurate to the hour and must return the corresponding time-of-day category.

You need to implement a function that takes an input hour (an integer from 0 to 23) and returns a string describing the time of day:

- Night — if the hour is in the range 0–5
- Morning — if the hour is in the range 6–11
- Day — if the hour is in the range 12–17
- Evening — if the hour is in the range 18–23

Input Format:
One integer N (0 ≤ N ≤ 23)

Output Format:
One string — the time-of-day category (Latin letters, without quotes)

Examples:

Example 1
Input: 3
Output: Night

Example 2
Input: 7
Output: Morning

Example 3
Input: 15
Output: Day

Example 4
Input: 20
Output: Evening
*/

import "fmt"

func GetTimeOfDay(hour int) string {

	// Validate input
	if hour < 0 || hour > 23 {
		return "Error: Invalid input"
	}

	if hour >= 0 && hour <= 5 {
		return "Night"
	}

	if hour >= 6 && hour <= 11 {
		return "Morning"
	}

	if hour >= 12 && hour <= 17 {
		return "Day"
	}

	if hour >= 18 && hour <= 23 {
		return "Evening"
	}

	return "" // technically unreachable
}

func GetTimeOfDay2(hour int) string {

	switch {
	case hour < 0 || hour > 23:
		return "Error: Invalid input"
	case hour >= 0 && hour <= 5:
		return "Night"
	case hour >= 6 && hour <= 11:
		return "Morning"
	case hour >= 12 && hour <= 17:
		return "Day"
	case hour >= 18 && hour <= 23:
		return "Evening"
	default:
		return ""
	}
}

func main() {
	fmt.Println("----------- if else version -----------")
	fmt.Println(GetTimeOfDay(3))
	fmt.Println(GetTimeOfDay(7))
	fmt.Println(GetTimeOfDay(15))
	fmt.Println(GetTimeOfDay(20))

	fmt.Println("----------- switch case version -----------")
	fmt.Println(GetTimeOfDay2(3))
	fmt.Println(GetTimeOfDay2(7))
	fmt.Println(GetTimeOfDay2(15))
	fmt.Println(GetTimeOfDay2(20))
}
