package main

import (
	"bytes"
	"container/list"
	"fmt"
)

func main() {
	input := ""

	fmt.Print("Please enter a word: ")
	fmt.Scanln(&input)

	fmt.Printf("Is %s palindrome: %t\n\n", input, palindrome(input))

} // main

func palindrome(testString string) bool {
	// I didn't find a native, out of the box implementation of a stack in Go
	// So, let's use a linked list a stack
	aStack := list.New()

	var stringBuilder bytes.Buffer

	// Push 'testString' into 'aStack'
	for i := 0; i < len(testString); i++ {
		// Convert 'string' characters into 'runes' because Go is so cool. Wow!
		aStack.PushBack(rune(testString[i]))
	}

	// Debug
	// fmt.Print("Debug: palindrome(): ")
	// for e := aStack.Front(); e != nil; e = e.Next() {
	// 	fmt.Printf("%c", e.Value.(rune))
	// }
	// fmt.Print("\n")

	// Pops out the stack into a 'bytes.Buffer' (like Java's StringBuilder)
	for e := aStack.Back(); e != nil; e = e.Prev() {
		// Saw this in StackOverflow https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
		stringBuilder.WriteRune(e.Value.(rune))
	}

	// Debug
	// fmt.Println("Debug: palindrome(): stringBuilder: ", string(stringBuilder.String()))

	if testString == stringBuilder.String() {
		return true
	}

	return false
}
