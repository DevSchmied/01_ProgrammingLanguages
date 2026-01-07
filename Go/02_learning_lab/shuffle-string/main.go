package main

/* Leetcodetask: 1528. Shuffle String

You are given a string s and an integer array indices of the same length.

The string s will be shuffled such that the character at the i-th position
moves to indices[i] in the shuffled string.

Return the shuffled string.

Example 1:
Input:
    s = "codeleet"
    indices = [4,5,6,7,0,2,1,3]
Output:
    "leetcode"
Explanation:
    "codeleet" becomes "leetcode" after shuffling.

Example 2:
Input:
    s = "abc"
    indices = [0,1,2]
Output:
    "abc"
Explanation:
    Each character remains in its original position.

Constraints:
    - s.length == indices.length == n
    - 1 <= n <= 100
    - s consists only of lowercase English letters
    - 0 <= indices[i] < n
    - All values in indices are unique
*/
import "fmt"

func restoreStringBySearch(s string, indices []int) string {
	if len(s) != len(indices) {
		return ""
	}

	resRunes := make([]rune, 0, len(s))

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if indices[j] == i {
				resRunes = append(resRunes, rune(s[j]))
				break
			}
		}
	}

	return string(resRunes)
}

func restoreStringDirect(s string, indices []int) string {
	if len(s) != len(indices) {
		return ""
	}

	sRunes := []rune(s)
	resRunes := make([]rune, len(s))

	for i := 0; i < len(sRunes); i++ {
		resRunes[indices[i]] = sRunes[i]
	}

	return string(resRunes)
}

func main() {
	fmt.Println("=======================1. Solution=======================")
	fmt.Println(restoreStringBySearch("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
	fmt.Println(restoreStringBySearch("abc", []int{0, 1, 2}))

	fmt.Println("\n=======================2. Solution=======================")
	fmt.Println(restoreStringDirect("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
	fmt.Println(restoreStringDirect("abc", []int{0, 1, 2}))
}
