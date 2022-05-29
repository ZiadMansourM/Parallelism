package main

import (
	"fmt"
	"time"
)

/*
	Notes:
		- Print 100x
		- What is a routine
*/

func print(letter string) {
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Nanosecond)
		fmt.Print(letter)
	}
}

func main() {
	letters  := []string {"A", "B", "C", "D", "E"}
	for _, letter := range letters {
		print(letter)
	}
}