package main

import (
	"fmt"
	"time"
)

func ping1(c chan string) {
	time.Sleep(time.Second * 3)
	c <- "ping on channel1"
}

func ping2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "ping on channel2"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go ping1(channel1)
	go ping2(channel2)

	select {
	case msg1 := <-channel1:
		fmt.Println("recieved", msg1)
	
	case msg2 := <-channel2:
		fmt.Println("received", msg2)
		
	case <- time.After(500 * time.Millisecond):
		fmt.Println("no messages recieved. giving up")
	}
}