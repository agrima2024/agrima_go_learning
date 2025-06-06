package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	i := 2
	time.Sleep(time.Second * time.Duration(i))
	fmt.Printf("Your %+v second timer is up!\n", i)
}

func main() {
	slowFunc()
}