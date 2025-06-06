package main

import (
    "fmt"
)

func triangle_checker(angle1 int, angle2 int) string  {
    var angle3 int 
    angle3 = 180 - (angle1+angle2)
    
    if angle1 + angle2 > 180 {
        fmt.Println("This is not a valid triangle! Try again with different angles.")
    }
    
    if angle1 > angle2 && angle1 > angle3{
        if angle1 == 90 {
            fmt.Println("This is a right triangle!")
        }
    
        if angle1 < 90 {
            fmt.Println("This is an acute triangle!")
        }
    
        if angle1 > 90 {
            fmt.Println("This is an obtuse triangle!")
        }
    }

    if angle2 > angle1 && angle2 > angle3{
        if angle2 == 90 {
            fmt.Println("This is a right triangle!")
        }
    
        if angle2 < 90 {
            fmt.Println("This is an acute triangle!")
        }
    
        if angle2 > 90 {
            fmt.Println("This is an obtuse triangle!")
        }
    }

    if angle3 > angle1 && angle3 > angle2{
        if angle3 == 90 {
            fmt.Println("This is a right triangle!")
        }
    
        if angle3 < 90 {
            fmt.Println("This is an acute triangle!")
        }
    
        if angle3 > 90 {
            fmt.Println("This is an obtuse triangle!")
        }
    }

    return ""
}

func main() {
    triangle_checker(50, 60)
}