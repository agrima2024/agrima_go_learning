package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile("questions_123.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var parsedData []map[string]interface{}
	err = json.Unmarshal(content, &parsedData)
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
	name_input := scanner.Text()

	var AnswerCounter int
	for _, questionData := range parsedData {
		question := questionData["Question"].(string)
		answer := questionData["Answer"].(string)
		fmt.Println("Question: ", question)

		fmt.Print("Enter your answer: ")
		if !scanner.Scan() {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}
		inputAnswer := scanner.Text()

		if inputAnswer == answer {
			fmt.Println("Correct Answer!")
			AnswerCounter = AnswerCounter + 1
		} else {
			fmt.Println("Wrong, the right answer is", answer)
		}
	}

	fmt.Println("You got", AnswerCounter, "questions right out of", len(parsedData), "questions.")

	type Person struct {
		Name       string  `json:"name"`
		Questions  int     `json:"questions"`
		Correct    int     `json:"correct"`
		Percentage float32 `json:"percentage"`
	}

	data := map[string]interface{}{
		"name":       name_input,
		"questions":  int(len(parsedData)),
		"correct":    AnswerCounter,
		"percentage": (float64(AnswerCounter) / float64(len(parsedData))) * 100,
	}
	fmt.Println(int(len(parsedData)))
	fmt.Println(AnswerCounter)
	fmt.Println((float64(AnswerCounter)) / float64(len(parsedData)))

	fmt.Printf("%T\n", len(parsedData))
	fmt.Printf("%T\n", AnswerCounter)
	fmt.Printf("%T\n", float64(AnswerCounter)/float64(len(parsedData)))

	fmt.Println(data)
	jsonData, err := json.Marshal(data)
	file, err := os.Create(name_input + ".json")
	defer file.Close()
	_, err = file.Write(jsonData)
}
