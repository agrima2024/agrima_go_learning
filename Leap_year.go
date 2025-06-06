package main

import (
    "fmt"
)

func leap_year(year int) string {

    if year%4 == 0 {
        fmt.Println(year, "given is a leap year!")
    } else {
        fmt.Println( year, "is not a leap year.")
    }

    return ""
}

func main() {
    leap_year(2024) 
    leap_year(2023) 
}