package main

import "fmt"

/*
LeetCode 73: Set Matrix Zeroes
Quelle: LeetCode (https://leetcode.com)

Problem Statement:

Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.

Examples:
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
Explanation: The element at position (1,1) is 0, so row 1 and column 1 are set to 0.

Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
Explanation: The zeros at positions (0,0) and (0,3) cause row 0, column 0, and column 3 to be set to 0. Row 1 and row 2 have zeros in affected columns.

Constraints:
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-2^31 <= matrix[i][j] <= 2^31 - 1
*/

func setZeroes(matrix [][]int) {
	rowIdxSeen := make(map[int]struct{})
	colIdxSeen := make(map[int]struct{})
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				rowIdxSeen[row] = struct{}{}
				colIdxSeen[col] = struct{}{}
			}
		}
	}

	for row := 0; row < len(matrix); row++ {
		if _, exists := rowIdxSeen[row]; exists {
			for col := 0; col < len(matrix[0]); col++ {
				matrix[row][col] = 0
			}
		}
	}

	for col := 0; col < len(matrix[0]); col++ {
		if _, exists := colIdxSeen[col]; exists {
			for row := 0; row < len(matrix); row++ {
				matrix[row][col] = 0
			}
		}
	}
	fmt.Println(matrix)
}

func main() {
	// Test case 1: 3x3 matrix with single zero at center
	matrix1 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix1)

	// Test case 2: 3x4 matrix with zeros at (0,0) and (0,3)
	matrix2 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix2)

}
