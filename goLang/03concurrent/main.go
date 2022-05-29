package main

import (
	"fmt"
	"time"
)

/*
	Notes:
		- How the go schedular works
		- If the main routine dies:
			all child routines are killed :(
*/

func print(letter string) {
	for i := 0; i <= 100; i++ {
		// time.Sleep(time.Nanosecond)
		fmt.Print(letter)
	}
}

func main() {
	letters  := []string {"A", "B", "C", "D", "E"}
	for _, letter := range letters {
		go print(letter)
	}
	time.Sleep(3*time.Second)
}