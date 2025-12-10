package main

/*
Task. Employee Ages

You are given data about company employees.
You must determine:
1- the minimum age of an employee;
2- the median age of all employees;
3- the maximum age.

The median is the central number in a sorted list.
If the number of employees is even, the median is calculated as the average of the two middle values, rounded down (integer division).

==============

Input format

A single line containing employee data separated by ;.
Each employee is described as:
Name,Age,Department

Age is an integer from 20 to 100.

Output format

Three integers separated by spaces:
MinAge MedianAge MaxAge

==============

Example 1

Input:

John,28,Engineering;Alex,34,HR;Dennis,45,Marketing;Anna,30,Engineering;Bob,24,HR


Output:

24 30 45

==============

Example 2

Input:

Paul,28,Engineering;Elena,34,Marketing


Output:

28 31 34


(Median = (28 + 34) / 2 = 31)
*/

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

// calculateAgeStatistics returns: min, median, max ages.
func calculateAgeStatistics(input string) ([]int, error) {

	// Empty input guard
	if strings.TrimSpace(input) == "" {
		return nil, errors.New("Input is empty")
	}

	// Split employees by ';'
	records := strings.Split(strings.TrimSpace(input), ";")
	ages := make([]int, 0, len(records))

	for _, record := range records {
		fields := strings.Split(record, ",")
		if len(fields) < 2 {
			return nil, errors.New("Invalid employee format")
		}

		// Parse age
		age, err := strconv.Atoi(fields[1])
		if err != nil || age < 20 || age > 100 {
			return nil, errors.New("Invalid age value")
		}

		ages = append(ages, age)
	}

	sort.Ints(ages)

	if len(ages) == 0 {
		return nil, errors.New("No age data provided")
	}

	if len(ages) == 1 {
		// min = median = max
		return []int{ages[0], ages[0], ages[0]}, nil
	}

	// Median calculation
	var median int
	if len(ages)%2 == 0 {

		left := ages[len(ages)/2-1]
		right := ages[len(ages)/2]
		median = (left + right) / 2
	} else {

		median = ages[len(ages)/2]
	}

	// min, median, max
	return []int{ages[0], median, ages[len(ages)-1]}, nil
}
