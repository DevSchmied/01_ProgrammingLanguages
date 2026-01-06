package main

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
