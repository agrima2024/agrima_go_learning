package main

import (
    "fmt"
    "regexp"
)
 
func calculater(a int, b int, operation string) int {
    sub := regexp.MustCompile(`^[-]`)
    add := regexp.MustCompile(`^[+]`)
    div := regexp.MustCompile(`^[/]`)
    mult := regexp.MustCompile(`^[*]`)
	var c int = 0

    if add.MatchString(operation) {
    	fmt.Println(a+b)
		c = a+b
    } 
	if sub.MatchString(operation) {
    	fmt.Println(a-b)
		c =  a-b
    }
    if div.MatchString(operation) {
    	fmt.Println(a/b)
		c =  a/b
    }
    if mult.MatchString(operation) {
    	fmt.Println(a*b)
		c = a*b
    } 
	return c
}

func main() {
   calculater(12, 3, "*")

}