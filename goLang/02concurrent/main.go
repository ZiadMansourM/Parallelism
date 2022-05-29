package main

import (
	"fmt"
	"time"
)

/*
	Notes:
		- Print 100x
		- Spin Go routine for each letter
		- Threads are 1000 to 2000 kB :: Go routines are 2 KB
		- "Go Schedular" - "Single/MultiCore" - "Realtion"
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
		go print(letter)
	}
	time.Sleep(3*time.Second)
}