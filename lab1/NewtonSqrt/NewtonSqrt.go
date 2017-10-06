package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}
	z := 1.0
	for i := 0; i < int(x); i++ {
		z = z - ((z*z - x) / (2 * z))
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(math.Sqrt(2))
}
