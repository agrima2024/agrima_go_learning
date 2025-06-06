package main

import (
	"fmt"
)

type Movie struct {
	Name string
	Rating float32
}

func main() {
	var m Movie 
	fmt.Printf("%+v\n", m)
	m.Name = "Metropolis"
	m.Rating = 0.9918;
	fmt.Printf("%+v\n", m)
}