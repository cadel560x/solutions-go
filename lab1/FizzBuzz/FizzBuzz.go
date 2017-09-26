package main

import (
	"fmt"
	"strconv"
)

func main() {
	var output string

	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			output = "Fizz"
		}

		if i%5 == 0 {
			output += "Buzz"
		}

		if len(output) == 0 {
			output = strconv.Itoa(i)
		}

		fmt.Println(output)

		output = ""

	} // for
} // main
