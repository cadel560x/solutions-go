package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Printf("%02d:%02d:%02d %d-%02d-%02d\n",
		t.Hour(), t.Minute(), t.Second(),
		t.Year(), t.Month(), t.Day())

}
