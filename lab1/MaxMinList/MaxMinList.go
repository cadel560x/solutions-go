package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create a list
	aList := new(list.List)

	var minList, maxList uint8

	// Set the seed for the random source
	s1 := rand.NewSource(time.Now().UnixNano())

	// Populate the list with 10 positve integers numbers greater than zero and lower than 256
	for i := 0; i < 10; i++ {
		// Debug
		// randomNumber := uint8(rand.Intn(255) + 1)
		randomNumber := uint8(rand.New(s1).Intn(255) + 1)
		aList.PushBack(randomNumber)
	}

	// Print the list
	fmt.Println("The list:")
	for e := aList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(uint8))
	}

	// Find out the smallest and largest in the list
	minList, maxList = minMaxList(aList)

	// Print the output
	fmt.Println("The smallest element is: ", minList)
	fmt.Println("The largest element is: ", maxList)
} // main

func minMaxList(lista *list.List) (uint8, uint8) {
	var min, max uint8

	// Set the 'min' and 'max' local variables to the first element of the list
	min = lista.Front().Value.(uint8)
	max = min

	// Debug
	// fmt.Println("min: ", min, " max: ", max)

	// Set the iterator to the second element in the list
	for e := lista.Front().Next(); e != nil; e = e.Next() {
		if e.Value.(uint8) < min {
			min = e.Value.(uint8)
		}

		if e.Value.(uint8) > max {
			max = e.Value.(uint8)
		}
	} // for

	return min, max
} // minMaxList
