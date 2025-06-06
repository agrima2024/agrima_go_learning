package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Question struct {
    Question string `json:"Question"`
    Answer   string `json:"Answer"`
}

func main() {
    fmt.Println("Welcome to the Trivia Quiz!")

  
    content, err := ioutil.ReadFile("states_capitals.json")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    var questions []Question
    err = json.Unmarshal(content, &questions)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Printf("Enter your name: ")
    if !scanner.Scan() {
        fmt.Println("Error reading input:", scanner.Err())
        return
    }
    name := scanner.Text()
    
    
    	err = json.Unmarshal(content, &questions)
if err != nil {
    fmt.Println("Error parsing JSON:", err)
    return
}
fmt.Println("Number of questions:", len(questions))




    var score int
    for _, question := range questions {
        fmt.Println("Question:", question.Question)

        fmt.Print("Enter your answer: ")
        if !scanner.Scan() {
            fmt.Println("Error reading input:", scanner.Err())
            return
        }
        answer := scanner.Text()

        if answer == question.Answer {
            fmt.Println("Correct Answer!")
            score++
        } else {
            fmt.Println("Wrong, the right answer is", question.Answer)
        }
    }

    percentage := float32(score) / float32(len(questions)) * 100
    fmt.Printf("%s, you got %d questions right out of %d (%.2f%%)\n", name, score, len(questions), percentage)

    data := map[string]interface{}{
        "name":       name,
        "questions":  len(questions),
        "correct":    score,
        "percentage": percentage,
    }

    jsonData, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error creating results JSON:", err)
        return
    }

    file, err := os.Create(name + ".json")
    defer file.Close()
    if err != nil {
        fmt.Println("Error creating results file:", err)
        return
    }

    _, err = file.Write(jsonData)
    if err != nil {
        fmt.Println("Error writing results to file:", err)
        return
    }

    fmt.Println("Results saved to", name+".json")
    fmt.Println("Thanks for playing!")
}

