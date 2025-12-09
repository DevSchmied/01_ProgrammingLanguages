package main

/*
Task. Trip Planning

You are given a list of cities and a list of hotels, where each hotel corresponds to a city by index.
You need to output all cities that have the maximum number of hotels.
If multiple cities share the maximum, output them in alphabetical order.
The city name must be printed with the first letter uppercase and the remaining letters lowercase.

================================

Input format

Two lines:
- A list of cities — strings (Latin alphabet), separated by commas without spaces.
- A list of hotels — strings, separated by commas without spaces.

The number of cities and hotels is the same.
City comparison must be case-insensitive.

================================

Output format

Each line:
- CityName NumberOfHotels

================================

Example 1

Input:
- Paris,London,London,Paris,Lyon,London
- HotelA,HotelB,HotelC,HotelD,HotelE,HotelF


Output:

London 3

(“London” appears 3 times — the maximum.)

================================

Example 2

Input:

- Nice,Nice,Nice,Nice,Nice,Bordeaux,Bordeaux,Bordeaux,Bordeaux,Bordeaux
- Negresco1,Negresco2,Negresco3,Negresco4,Negresco5,Regent1,Regent2,Regent3,Regent4,Regent5


Output:

Bordeaux 5
Nice 5
*/

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func planTrip(cityList, hotelList string) []string {

	cities := strings.Split(strings.ToLower(strings.TrimSpace(cityList)), ",")
	hotels := strings.Split(strings.ToLower(strings.TrimSpace(hotelList)), ",")
	if len(cities) != len(hotels) || len(cityList) == 0 || len(hotelList) == 0 {
		log.Printf("invalid input")
		return nil
	}

	// Counting number of hotels for each city
	count := make(map[string]int)
	for _, c := range cities {
		count[c]++
	}

	// Finding the maximum number of hotels
	var valMax int
	for _, v := range count {
		if v > valMax {
			valMax = v
		}
	}

	// Building the result (capitalize city name + count)
	result := []string{}
	for k, v := range count {
		if v == valMax {
			keyRune := []rune(k)
			city := strings.ToUpper(string(keyRune[0])) + string(keyRune[1:])
			result = append(result, fmt.Sprintf("%s %d", city, v))
		}
	}

	// Alphabetical sorting
	sort.Strings(result)

	return result
}
