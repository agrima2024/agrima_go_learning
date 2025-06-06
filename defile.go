package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("example01.txt")
	if err != nil {
		log.Fatal(err)
	}
}