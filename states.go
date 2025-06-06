package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type US struct {
	State   string `json:"State"`
	Capital string `json:"Capital"`
}

func directions() {
	fmt.Println("Welcome to the states game!")
	fmt.Println("To play this game, name all the states of the US.")
	fmt.Println("Press enter after naming a state to see if you are right or not.")
	fmt.Println("This game will stop when you have named all the states correctly.")
}

func states() {
	content, err := ioutil.ReadFile("states_capitals.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var states []US
	err = json.Unmarshal(content, &states)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	correctGuesses := make(map[string]bool)

	fmt.Println("\nStart guessing the states:")

	for {
		fmt.Printf("Your guess: ")
		if !scanner.Scan() {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}
		guess := strings.TrimSpace(scanner.Text())

		found := false
		for _, state := range states {
			if strings.EqualFold(state.State, guess) {
				if correctGuesses[state.State] {
					fmt.Println("You already guessed that state!")
				} else {
					fmt.Println("Correct!")
					correctGuesses[state.State] = true
				}
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Incorrect. Try again!")
		}

		counter := 50 - len(correctGuesses)
		fmt.Println("You have", counter, "states left to guess.")

		if counter == 0 {
			fmt.Println("Congratulations! You've guessed all the states!")
			break
		}
	}
}

func main() {
	directions()
	states()
}
