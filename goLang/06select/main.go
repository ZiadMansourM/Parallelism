package main

import (
	"fmt"
	"time"
)

/*
	Problem: Recieve From Multiple Channels
*/

func senderOne(c chan<- string){
	for {
		c <- "Confirmed in 500ms"
		time.Sleep(500*time.Millisecond)
	}
}
func senderTwo(c chan<- string){
	for {
		c <- "Confirmed in Two Seconds"
		time.Sleep(2*time.Second)
	}
}
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go senderOne(c1)
	go senderTwo(c2)
	for {
		select {
			case msg := <- c1:
				fmt.Println(msg)
			case msg := <- c2:
				fmt.Println(msg)
		}
	}
}