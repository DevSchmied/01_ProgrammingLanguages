package main

import "fmt"

func main() {
	var x int
	var p *int

	x = 126
	p = &x
	fmt.Println("Address:", p)
	fmt.Println("Value over pointer:", *p)

	*p = 621
	fmt.Println("Value over variable:", x)
}
