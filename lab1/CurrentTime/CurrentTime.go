package main

import (
	"fmt"
	"time"
)

func main() {
	// Get's current date and time
	t := time.Now()

	// Prints the time and then the date
	fmt.Printf("%02d:%02d:%02d %d-%02d-%02d\n",
		t.Hour(), t.Minute(), t.Second(),
		t.Year(), t.Month(), t.Day())

}
