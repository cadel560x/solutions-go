package main

import (
	"fmt"
	"math/big"
)

func main() {
	// '100!' is a huge number, we are using 'big.Int' type
	facto := new(big.Int).SetUint64(0)

	var digiSum uint16

	facto = factorial(100)
	fmt.Println("Factorial of 100: ", facto)

	digiSum = digitSum(facto)
	fmt.Println("The digit sum of ", facto, "is :", digiSum)

} // main

func factorial(n uint8) *big.Int {
	// 'big.Int' factorial variable with value of '1'
	factorial := new(big.Int).SetUint64(1)

	// 'n' paramenter as a 'big.Int'
	i := new(big.Int).SetUint64(uint64(n))

	// 'for i > 0' in 'big.Int'
	for i.Sign() > 0 {
		// 'factorial *= i' in 'big.Int'
		factorial.Mul(factorial, i)

		// 'i--' in 'big.Int'
		i.Sub(i, new(big.Int).SetUint64(1))

		// Debug
		// fmt.Println("Debug: func factorial: for loop: i", i)
		// fmt.Println("Debug: func factorial: for loop: factorial", factorial)
	}

	return factorial
} // factorial

func digitSum(m *big.Int) uint16 {
	n := new(big.Int)
	n.Set(m)

	var sum uint16
	var digit uint8

	// 'for i > 0' in 'big.Int'
	for n.Sign() > 0 {
		// 'digit = n % 10' in 'big.Int'
		digit = uint8(new(big.Int).Mod(n, new(big.Int).SetUint64(10)).Uint64())

		sum += uint16(digit)

		// 'n /= 10' in 'big.Int'
		n.Div(n, new(big.Int).SetUint64(10))

		// Debug
		// fmt.Println("Debug: func digitSum: for loop: n", n)
		// fmt.Println("Debug: func digitSum: for loop: digit", digit)
		// fmt.Println("Debug: func digitSum: for loop: sum", sum)
	}

	return sum
} // digitSum
