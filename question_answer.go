package main

import (
    "fmt"
    "os"
    "bufio"
)

func question() {
    file, abc := os.Open("/home/cssdesk/agrima_go/question.txt")
    if abc != nil {
        fmt.Println("File reading error", abc)
        return
    }
    defer file.Close()

    answer := make(map[string]string)
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        currentLineText := scanner.Text()
        if scanner.Scan() {
            answer[currentLineText] = scanner.Text()
        } else {
            fmt.Println("Warning: No answer found for question:", currentLineText)
        }
    }

    var correctAnswersCount int
    for question, correctAnswer := range answer {
        fmt.Println("Question:", question)
        fmt.Print("Enter your answer: ")

        scanner := bufio.NewScanner(os.Stdin) 
        scanner.Scan()
        userAnswer := scanner.Text()

        if userAnswer == correctAnswer {
            fmt.Println("Correct!")
            correctAnswersCount++
        } else {
            fmt.Println("Incorrect. The correct answer is:", correctAnswer)
        }
    }

    fmt.Println("You answered", correctAnswersCount, "out of", len(answer), "questions correctly.")
}

func main() {
    question()
}

// A text file with Q and A
// Need to write a function to read all the Q and corresponding A
// Need to write a function to ask one Question at a time and based on answer provided match with corresponding A and return correct or wrong
// keep a count to give final score at the end


// 	{
// 		question: "what is agrima age? "
// 		answer: 14
// 	},
// 	{
// 		question: "what is agrima age? "
// 		answer: 14
// 	},
// 	{
// 		question: "what is agrima age? "
// 		answer: 14
// 	},
// 	{
// 		question: "what is agrima age? "
// 		answer: 14
// 	},

// }