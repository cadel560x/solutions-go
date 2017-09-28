package main

import (
	"fmt"
	"math/big"
)

func main() {
	var output big.Int

	output = factorial(100)

	fmt.Println("Factorial of 100: ", output)
	// fmt.Println("The digit sum of ", output, "is :", digitSum(output))

} // main

func factorial(n uint8) big.Int {
	var factorial, m big.Int
	factorial.SetUint64(1)

	fmt.Println("Debug: func factorial: n", n)

	for n > 0 {
		m.SetUint64(uint64(n))
		factorial.Mul(factorial, m)
		n--
		fmt.Println("Debug: func factorial: for loop: factorial", factorial)
		fmt.Println("Debug: func factorial: for loop: n", n)
	}

	return factorial
} // factorial

func digitSum(n float64) float64 {
	var sum float64
	var digit uint8
	var pow10 uint64 = 1

	for n > 0 {
		pow10 *= 10
		digit = uint8(uint64(n) % pow10)
		sum += float64(digit)

		n /= float64(pow10)

		fmt.Println("Debug: func digitSum: for loop: n", n)
		fmt.Println("Debug: func digitSum: for loop: pow10", pow10)
		fmt.Println("Debug: func digitSum: for loop: digit", digit)
		fmt.Println("Debug: func digitSum: for loop: sum", sum)
	}

	return sum
} // digitSum
