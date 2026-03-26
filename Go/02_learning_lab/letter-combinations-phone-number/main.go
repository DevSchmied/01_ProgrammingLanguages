package main

/*
LeetCode 17: Letter Combinations of a Phone Number
Quelle: LeetCode (https://leetcode.com/problems/letter-combinations-of-a-phone-number)

Problem Statement:

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Digit to letter mapping:
2 -> "abc"
3 -> "def"
4 -> "ghi"
5 -> "jkl"
6 -> "mno"
7 -> "pqrs"
8 -> "tuv"
9 -> "wxyz"

Examples:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Explanation: Each digit maps to a set of letters. All combinations are formed by taking one letter from the first digit's mapping and one letter from the second digit's mapping.

Input: digits = "2"
Output: ["a","b","c"]
Explanation: Only one digit, so return all letters that digit maps to.

Constraints:
1 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9']
*/

func letterCombinations(digits string) []string {
	res := []string{""}

	letters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		temp := make([]string, 0)

		for _, prefix := range res {
			for j := 0; j < len(letters[digit]); j++ {
				temp = append(temp, prefix+string(letters[digit][j]))
			}
		}

		res = temp
	}

	return res
}
