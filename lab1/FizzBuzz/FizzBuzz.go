package main

import (
	"fmt"
	"strconv"
)

func main() {
	var output string

	// Iterate from '1' to '100'
	for i := 1; i < 101; i++ {
		// Multiple of 3 is 'Fizz'
		if i%3 == 0 {
			output = "Fizz"
		}

		// Multiple of 5 is 'Buzz'
		if i%5 == 0 {
			// Notice the '+='. This allows 'FizzBuzz' for multiples of '3' and '5'
			output += "Buzz"
		}

		// No multiple of '3' or '5'? Number is converted to 'string'
		if len(output) == 0 {
			output = strconv.Itoa(i)
		}

		fmt.Println(output)

		output = ""

	} // for
} // main
