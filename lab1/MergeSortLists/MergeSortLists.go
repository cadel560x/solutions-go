package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Create two lists
	// aList := new(list.List)
	// bList := new(list.List)
	var aList []int8
	var bList []int8

	// Populate them with a random amount of ordered random numbers
	populateList(&aList)
	populateList(&bList)

	// Print the lists
	fmt.Println("aList: ", aList)
	// printList(aList)
	fmt.Print("\n")
	fmt.Print("bList: ", bList)
	// printList(bList)

} // main

func populateList(lista *[]int8) {
	// Set the seed for the random source
	// s1 := rand.NewSource(time.Now().UnixNano())

	// Debug
	randomLength := uint8(rand.Intn(9) + 1)
	// randomLength := uint8(rand.New(s1).Intn(9) + 1)
	fmt.Println("Debug: populateList: randomLength: ", randomLength)

	*lista = make([]int8, randomLength, randomLength)

	for index, value := range *lista {
		// Generate negative and positive integer numbers between -128 to 127
		// Taken from StackOverflow: https://stackoverflow.com/questions/3938992/how-to-generate-random-positive-and-negative-numbers-in-java

		// Debug
		randomNumber := int8(rand.Intn(255) - 128)
		// randomNumber := int8(rand.New(s1).Intn(255) - 128)
		fmt.Println("Debug: populateList: for loop: randomNumber ", randomNumber)

		for randomNumber < value {
			randomNumber = int8(rand.Intn(255) - 128)
		}

		lista(index) = randomNumber
		// lista.PushBack(randomNumber)
	}
} // populateList

func printList(lista []int8) {
	for _, value := range lista {
		fmt.Print(value, " -> ")
	}
	fmt.Println("END")
} // printList
