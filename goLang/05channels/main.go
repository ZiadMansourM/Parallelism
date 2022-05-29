package main

import (
	"fmt"
	"time"
)

/*
	Problem: slow receiver

	Channels: Don't communicate by sharing memory; share memory by communicating.
	> Don't overengineer ITC by "shared memory, complicated, error-prone synchronisation primitives"
	but instead use message-passing between goroutines (green threads).
	> Queue. It is a simple message passing pattern.
	> int || complicated data structure like a map. Give away ownership by sending the value or a pointer
	to a different goroutine via a channel
	>>> there is no shared space, each goroutine only sees the portion of memory it owns.

	Note:
		# To send there has to be a reciever on the other side || buffer space
		- reciver Block waiting
		> always the sender is the one to close the channel
*/

func senderRoutine(c chan<- string, letters []string) {
	for _, letter := range letters {
		fmt.Println(letter, "sending ....")
		c <- letter
	}
	close(c)
}

func main() {
	letters  := []string {"A", "B", "C", "D"}
	c := make(chan string, 1)

	go senderRoutine(c, letters)

	time.Sleep(5*time.Second) // no one on the recieve side untill t = 5
	fmt.Println("Reciever Starts @t=5")


	for res := range c {
		fmt.Println("Reciever:", res)
	}
}