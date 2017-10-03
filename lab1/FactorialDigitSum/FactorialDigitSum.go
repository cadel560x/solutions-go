package main

import (
	"fmt"
	"math/big"
)

func main() {
	// var facto uint64
	facto := new(big.Int).SetUint64(0)
	// var digiSum uint16

	facto = factorial(100)
	fmt.Println("Factorial of 100: ", facto)

	digiSum = digitSum(facto)
	// fmt.Println("The digit sum of ", facto, "is :", digiSum)

} // main

func factorial(n uint8) *big.Int {
	// factorial := new(big.Int).SetUint64(1)
	// factorial := new(big.Int).SetUint64(uint64(n))
	// i := new(big.Int).SetUint64(uint64(n))

	// for i.Sign() > 0 {
	// 	factorial *= i
	// }

	return new(big.Int).MulRange(1, int64(n))
	// 93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000
} // factorial

func digitSum(n big.Int) uint16 {
	// Converting parameter 'n' from float64 type to big.Int type
	// nBigInt, _ := big.NewFloat(n).Int(nil)

	var sum uint16
	var digit uint8

	// for nBigInt.Sign() > 0 {
	// 	digit = uint8(new(big.Int).Mod(nBigInt, new(big.Int).SetUint64(10)).Uint64())
	// 	sum += uint16(digit)

	// 	nBigInt.Div(nBigInt, new(big.Int).SetUint64(10))

	for n.Sign() > 0 {
		digit = uint8(new(big.Int).Mod(nBigInt, new(big.Int).SetUint64(10)).Uint64())
		sum += uint16(digit)

		nBigInt.Div(nBigInt, new(big.Int).SetUint64(10))

		fmt.Println("Debug: func digitSum: for loop: n", nBigInt)
		fmt.Println("Debug: func digitSum: for loop: digit", digit)
		fmt.Println("Debug: func digitSum: for loop: sum", sum)
	}

	return sum
} // digitSum
