package main

import (
	"fmt"
	"time"

)

func slowFunc() {
	time.Sleep(time.Second * 2)
	c <- "slowFunc() finished"
}

func main() {
	c := make(chan string)
	go slowFunc()

	msg := <-c
	fmt.Println(msg)
}