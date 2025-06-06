package main

import (
	"io/ioutil"
	"log"
)

func main() {
	s := "Hello World"
	err := ioutil.WriteFile("example02.txt", []byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}
}