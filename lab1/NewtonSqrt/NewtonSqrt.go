package main

import (
	"fmt"
	"math"
)

func newtonSqrt(x float64) float64 {
	// Square root of 0 is 0
	if x == 0 {
		return 0
	}

	// Starting point
	z := 1.0

	for i := 0; i < int(x); i++ {
		z = z - ((z*z - x) / (2 * z))
	}
	return z
} // newtonSqrt

func main() {
	// Debug
	// fmt.Println(sqrt(2))
	// fmt.Println(math.Sqrt(2))

	numbers := 20
	for i := 0; i < numbers; i++ {
		mathSqrt := math.Sqrt(float64(i))
		newton := newtonSqrt(float64(i))
		fmt.Println(i, "square root:")
		fmt.Println("  math.Sqrt:", mathSqrt)
		fmt.Println("  Newton's method:", newton)
		fmt.Println("  Difference:", math.Abs(mathSqrt-newton))
	}
} // main
