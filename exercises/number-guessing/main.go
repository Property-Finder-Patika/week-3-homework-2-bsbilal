package main

import (
	"fmt"
	"math/rand"
)

var targetNumber int = 0

func main() {
	CreateNewGame()
}

func CreateNewGame() {
	targetNumber = rand.Intn(9999-1000) + 1000
	fmt.Println("The new game has begun.")
	NumberGuess()
}

func NumberGuess() {
	var guessedNumbers []map[string]string
	for true {
		fmt.Println("please enter new number")
		var number string
		fmt.Scanln(&number)
		if true {
			guessed := make(map[string]string)
			guessed["entered"] = number
			guessed["result"] = "-1+2"
			guessedNumbers = append(guessedNumbers, guessed)

		} else {

			fmt.Println("Please provide 1 values.")

		}

	}

}
