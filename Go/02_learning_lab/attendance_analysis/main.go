package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
TASK — Attendance Statistics Analysis

You are developing a program to analyze student attendance.
Each input line contains a student’s unique ID and the number of classes they attended.
Your task is to output the following statistics:

1. Total number of students
2. Total number of attended classes
3. Average number of attended classes per student
(rounded to one decimal place)
4. Number of students with low attendance
(less than 5 classes)
5. Number of students with medium attendance
(from 5 to 9 classes, inclusive)
6. Number of students with high attendance
(10 or more classes)

If the input contains any invalid records, output the line:
"no valid data"

Input Format

1. An integer 1 < N < 100 — the number of students.
2. Then follow N lines, each with two integers:

<student_id> <classes_attended>

Conditions:
- student_id > 0
- classes_attended ≥ 0

Output Format
Each item is printed on a separate line:

Total students: <number>
Total classes: <number>
Average classes: <number with 1 decimal>
Low attendance (<5): <number> students
Medium attendance (5-9): <number> students
High attendance (>=10): <number> students


If the data is invalid → print:
no valid data


Examples:

Example 1

Input:

5
1 3
2 7
3 12
4 0
5 10


Output:

Total students: 5
Total classes: 32
Average classes: 6.4
Low attendance (<5): 2 students
Medium attendance (5-9): 1 students
High attendance (>=10): 2 students

Example 2

Input:

3
2 -5
3 10
5 -10


Output:

no valid data

Example 3

Input:

2
1 0
2 0


Output:

Total students: 2
Total classes: 0
Average classes: 0.0
Low attendance (<5): 2 students
Medium attendance (5-9): 0 students
High attendance (>=10): 0 students
*/

type AnalyzeAttendanceResult struct {
	TotalStudents, TotalLessons int
	Low, Medium, High           int
	Valid                       int
}

// AnalyzeAttendance analyzes the attendance data.
func AnalyzeAttendance(data []string) *AnalyzeAttendanceResult {

	if len(data) < 2 {
		return &AnalyzeAttendanceResult{Valid: -1}
	}

	n, err := strconv.Atoi(strings.TrimSpace(data[0]))
	if err != nil || n <= 1 || n >= 100 {
		return &AnalyzeAttendanceResult{Valid: -1}
	}

	if len(data)-1 < n {
		return &AnalyzeAttendanceResult{Valid: -1}
	}

	var totalStd int
	var totalLes int
	var low int
	var medium int
	var high int

	for i := 1; i <= n; i++ {

		fields := strings.Fields(data[i])
		if len(fields) != 2 {
			return &AnalyzeAttendanceResult{Valid: -1}
		}

		firstEl, err1 := strconv.Atoi(fields[0])
		secEl, err2 := strconv.Atoi(fields[1])

		if err1 != nil || err2 != nil {
			return &AnalyzeAttendanceResult{Valid: -1}
		}

		if firstEl <= 0 || secEl < 0 {
			return &AnalyzeAttendanceResult{Valid: -1}
		}

		totalStd++
		totalLes += secEl

		switch {
		case secEl >= 0 && secEl < 5:
			low++
		case secEl >= 5 && secEl < 10:
			medium++
		case secEl >= 10:
			high++
		}
	}
	return &AnalyzeAttendanceResult{
		TotalStudents: totalStd,
		TotalLessons:  totalLes,
		Low:           low,
		Medium:        medium,
		High:          high,
		Valid:         1,
	}
}

// PrintResults prints the analysis results.
func PrintResults(analyzeAttendanceResult *AnalyzeAttendanceResult) {
	if analyzeAttendanceResult.Valid < 0 {
		fmt.Println("no valid data")
		return
	}

	avg := 0.0
	if analyzeAttendanceResult.TotalStudents > 0 {
		avg = float64(analyzeAttendanceResult.TotalLessons) / float64(analyzeAttendanceResult.TotalStudents)
	}

	fmt.Printf("Total students: %d\n", analyzeAttendanceResult.TotalStudents)
	fmt.Printf("Total classes: %d\n", analyzeAttendanceResult.TotalLessons)
	fmt.Printf("Average classes: %.1f\n", avg)
	fmt.Printf("Low attendance (<5): %d students\n", analyzeAttendanceResult.Low)
	fmt.Printf("Medium attendance (5-9): %d students\n", analyzeAttendanceResult.Medium)
	fmt.Printf("High attendance (>=10): %d students\n", analyzeAttendanceResult.High)

}

func main() {

	fmt.Println("=== Example 1 ===")
	ex1 := []string{"5", "1 3", "2 7", "3 12", "4 0", "5 10"}
	r1 := AnalyzeAttendance(ex1)
	PrintResults(r1)

	fmt.Println()

	fmt.Println("=== Example 2 ===")
	ex2 := []string{"3", "2 -5", "3 10", "5 -10"}
	r2 := AnalyzeAttendance(ex2)
	PrintResults(r2)

	fmt.Println()

	fmt.Println("=== Example 3 ===")
	ex3 := []string{"2", "1 0", "2 0"}
	r3 := AnalyzeAttendance(ex3)
	PrintResults(r3)
}
