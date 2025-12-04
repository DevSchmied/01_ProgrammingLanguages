package main

import (
	"fmt"
	"strings"
)

/*
TASK 3 — Delivery Route Analysis

You work as a logistician in a transport company. You need to determine:

- Which routes (city pairs “from → to”) were used most frequently
- How many unique routes there were
- How many routes occurred exactly once

Input:
A list of shipments, each represented by a pair of cities (origin → destination).

Output:
- A list of routes with the maximum number of shipments, each printed in the format:
	<from> <to>: <count>
	(If several routes share the same maximum count — print all of them in order of first appearance.)
- Total number of unique routes
- Number of routes used exactly once
- If the input contains any invalid entries, print:
	invalid data

---

Input Format
First line: one integer N, where 2 < N < 100 — the number of shipments
Each of the next N lines contains two city names (strings without spaces), separated by a single space
Each city name is between 1 and 20 characters

---

Output Format
First print all routes with the highest shipment count:
<from> <to>: <count>

(one per line)

Then print:
Unique routes: <number>
Routes with one shipment: <number>

If invalid input is detected:
invalid data

---

Example 1

Input:

7
Moscow SPb
SPb Moscow
Moscow SPb
Kazan Moscow
Moscow SPb
Kazan Moscow
Kazan Moscow


Output:

Moscow SPb: 3
Unique routes: 3
Routes with one shipment: 1

---

Example 2

Input:

5
Moscow SPb
SPb Moscow
Kazan Moscow
Moscow SPb
Kazan SPb


Output:

Moscow SPb: 2
SPb Moscow: 1
Kazan Moscow: 1
Kazan SPb: 1
Unique routes: 4
Routes with one shipment: 4

---

Example 3

Input:

2
Moscow
SPb Kazan


Output:

invalid data

*/

type Route struct {
	From string
	To   string
}

type RouteStat struct {
	MaxRoutes      []Route
	MaxRoutesCount int
	UniqueCount    int
	Once           int
}

// AnalyzeRoutes analyzes the list of routes and returns the route statistics.
func AnalyzeRoutes(routes []Route) RouteStat {

	// Validate number of routes
	if len(routes) <= 2 || len(routes) >= 100 {
		return RouteStat{}
	}

	// result maps "From To" -> how many times the route occurs
	result := make(map[string]int)

	// keys stores the order of the first appearance of each unique route
	keys := make([]string, 0, len(routes))

	// Count occurrences and track order
	for i := 0; i < len(routes); i++ {

		fromTrimmed := strings.TrimSpace(routes[i].From)
		toTrimmed := strings.TrimSpace(routes[i].To)

		// Validate city names
		if len(fromTrimmed) < 1 || len(fromTrimmed) > 20 ||
			len(toTrimmed) < 1 || len(toTrimmed) > 20 {

			fmt.Println("invalid data")
			return RouteStat{}
		}

		key := fromTrimmed + " " + toTrimmed

		// If this route appears for the first time — remember it
		if _, exists := result[key]; !exists {
			keys = append(keys, key)
		}

		// Count the route
		result[key]++
	}

	// If somehow no valid routes exist
	if len(result) == 0 {
		return RouteStat{}
	}

	routeStat := RouteStat{
		UniqueCount: len(result),
	}

	// Determine:
	// 1) the maximum shipment count
	// 2) how many routes occurred once
	for _, count := range result {
		if count > routeStat.MaxRoutesCount {
			routeStat.MaxRoutesCount = count
		}
		if count == 1 {
			routeStat.Once++
		}
	}

	// Build the list of routes with maximum shipments (in order of appearance)
	for _, key := range keys {
		if result[key] == routeStat.MaxRoutesCount {
			parts := strings.SplitN(key, " ", 2)
			routeStat.MaxRoutes = append(routeStat.MaxRoutes, Route{
				From: parts[0],
				To:   parts[1],
			})
		}
	}

	return routeStat
}
