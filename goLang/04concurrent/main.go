package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func print(letter string) {
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Nanosecond)
		fmt.Print(letter)
	}
	wg.Done()
}

func main() {
	letters  := []string {"A", "B", "C", "D", "E"}
	for _, letter := range letters {
		wg.Add(1)
		go print(letter)
	}
	wg.Wait()
}