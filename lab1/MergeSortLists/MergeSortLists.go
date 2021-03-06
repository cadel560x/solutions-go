package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate lists
	slice1 := generateSlice(5)
	slice2 := generateSlice(5)

	// Sort the lists
	slice1 = mergeSort(slice1)
	slice2 = mergeSort(slice2)

	// Print the lists
	fmt.Println("\nLists before merge:")
	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)

	fmt.Println("\nLists after merge:")
	fmt.Println(merge(slice1, slice2))

} // main

// Generates a list of 'size', filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return slice
} // generateSlice

// Runs MergeSort algorithm on a single list
func mergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return merge(mergeSort(slice[:mid]), mergeSort(slice[mid:]))
} // mergeSort

// Merges left and right lists into newly created list
func merge(left, right []int) []int {

	// Set the 'size' of the resulting 'slice' after merging
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		// If there are remaining elements of the 'right list' to put in the resulting 'slice', put them in.
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
			// If there are remaining elements of the 'left list' to put in the resulting 'slice', put them in.
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
			// If 'left list' element is smaller than 'right list' element, goes into the resulting 'slice'
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			// If 'right list' element is smaller than 'left list' element, goes into the resulting 'slice'
			slice[k] = right[j]
			j++
		} // if - else if - else if - else

	} // for

	return slice
} // merge
