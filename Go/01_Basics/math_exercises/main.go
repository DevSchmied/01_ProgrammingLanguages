package main

import (
	"fmt"
	"math"
)

func printHeader(n int) {
	fmt.Printf("\n------------------------------%d. exercise------------------------------\n", n)
}

func main() {

	fmt.Println("------------------------------math-exercises------------------------------")

	/* 1. Declare two variables: a equal to 3.5 and b equal to 4.2.
	   Find the MINIMUM of these two values and store it in the variable minValue. */

	printHeader(1)

	a := 3.5
	b := 4.2
	minValue := math.Min(a, b)

	fmt.Printf("The minimum of %.1f and %.1f is %.1f\n", a, b, minValue)

	/* 2. Declare two variables: c equal to 3.5 and d equal to 4.2.
	   Find the MAXIMUM of these two values and store it in the variable maxValue. */

	printHeader(2)

	c := 3.5
	d := 4.2

	maxValue := math.Max(c, d)

	fmt.Printf("The maximum of %.1f and %.1f is %.1f\n", c, d, maxValue)

	/* 3. Declare a variable x equal to 16.0.
	   Find the SQUARE ROOT of this number and store it in the variable sqrtValue. */

	printHeader(3)

	x := 16.0
	sqrtValue := math.Sqrt(x)

	fmt.Printf("The square root of %.2f is %.2f\n", x, sqrtValue)

	/* 4. Declare a variable y equal to -7.8.
	   Find the ABSOLUTE VALUE of this number and store it in the variable absValue. */

	printHeader(4)

	y := -7.8

	absValue := math.Abs(y)

	fmt.Printf("The absolute value of %.2f is %.2f\n", y, absValue)

	/* 5. Declare a variable base equal to 2.0 and a variable exponent equal to 3.0.
	   Find the value of base raised to the POWER of exponent and store it in the variable valueExp. */

	printHeader(5)

	base := 2.0
	exponent := 3.0

	valueExp := math.Pow(base, exponent)

	fmt.Printf("%.1f raised to the power of %.1f is %.1f\n", base, exponent, valueExp)

	/* 6. Declare a variable angle equal to π/2.
	   Find the SINE of this angle and store it in the variable valueSine. */

	printHeader(6)

	angle := math.Pi / 2
	valueSine := math.Sin(angle)

	fmt.Printf("The sine of %.2f is %.2f\n", angle, valueSine)

	/* 7. Declare a variable anglePi equal to π.
	   Find the COSINE of this angle and store it in the variable valueCosine. */

	printHeader(7)

	anglePi := math.Pi
	valueCosine := math.Cos(anglePi)

	fmt.Printf("The cosine of %.2f is %.2f\n", anglePi, valueCosine)

	/* 8. Declare a variable min and store in it the MINIMUM VALUE for the type INT using the constant. */

	printHeader(8)

	min := math.MinInt64

	fmt.Printf("The minimum value of int is %d\n", min)

	/* 9. Declare a variable max and store in it the MAXIMUM VALUE for the type INT using the constant. */

	printHeader(9)

	max := math.MaxInt64

	fmt.Printf("The maximum value of int is %d\n", max)

}
