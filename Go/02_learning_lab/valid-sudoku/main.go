package main

import "fmt"

/*
LeetCode 36: Valid Sudoku
Quelle: LeetCode (https://leetcode.com)

Problem Statement:

Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
1. Each row must contain the digits 1-9 without repetition.
2. Each column must contain the digits 1-9 without repetition.
3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:
- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.

Examples:

Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
Explanation: This is a valid partially filled Sudoku board.

Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

Constraints:
board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'
*/

func isValidSudoku(board [][]byte) bool {
	rowSeen := make(map[byte]struct{})
	colSeen := make(map[byte]struct{})

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if board[i][j] != '.' {
				if _, exists := rowSeen[board[i][j]]; exists {
					return false
				}
				rowSeen[board[i][j]] = struct{}{}
			}

			if board[j][i] != '.' {
				if _, exists := colSeen[board[j][i]]; exists {
					return false
				}
				colSeen[board[j][i]] = struct{}{}
			}
		}
		rowSeen = map[byte]struct{}{}
		colSeen = map[byte]struct{}{}
	}

	for rowBlock := 0; rowBlock < 9; rowBlock += 3 {
		for colBlock := 0; colBlock < 9; colBlock += 3 {
			blockSeen := make(map[byte]struct{})
			for row := 0; row < 3; row++ {
				for col := 0; col < 3; col++ {
					if board[rowBlock+row][colBlock+col] != '.' {
						if _, exists := blockSeen[board[rowBlock+row][colBlock+col]]; exists {
							return false
						}
						blockSeen[board[rowBlock+row][colBlock+col]] = struct{}{}
					}
				}
			}
		}
	}

	return true
}

func main() {
	// Test case 1: Valid Sudoku board (from Example 1)
	board1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	result1 := isValidSudoku(board1)
	fmt.Printf("Test 1 - Valid Sudoku board -> Result: %t (Expected: true)\n", result1)

	// Test case 2: Invalid Sudoku board with duplicate 8 in top-left 3x3 sub-box (from Example 2)
	board2 := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	result2 := isValidSudoku(board2)
	fmt.Printf("Test 2 - Invalid Sudoku board with duplicate 8 in sub-box -> Result: %t (Expected: false)\n", result2)
}
