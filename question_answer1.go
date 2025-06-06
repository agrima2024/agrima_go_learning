package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
)

type QuestionAnswer struct {
    Question string
    Answer   string
}

func askQuestions(questions []QuestionAnswer) int {
    var correctAnswersCount int

    for _, qa := range questions {
        fmt.Println("Question:", qa.Question)
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter your answer: ")
        userAnswer, _ := reader.ReadString('\n')
        userAnswer = userAnswer[:len(userAnswer)-1] // Remove trailing newline

        if userAnswer == qa.Answer {
            fmt.Println("Correct!")
            correctAnswersCount++
        } else {
            fmt.Println("Incorrect. The correct answer is:", qa.Answer)
        }
    }

    return correctAnswersCount
}

func readQuestions(filename string) []QuestionAnswer {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return nil
    }
    defer file.Close()

    var questions []QuestionAnswer
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&questions)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        return nil
    }

    return questions
}

func main() {
    questions := readQuestions("questions_123.json")
    if questions == nil {
        fmt.Println("Error reading questions from JSON file.")
        return
    }

    correctAnswers := askQuestions(questions)
    fmt.Println("You answered", correctAnswers, "out of", len(questions), "questions correctly.")
}