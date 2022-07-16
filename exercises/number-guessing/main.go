package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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

		var sNumber string
		fmt.Scanln(&sNumber)

		numberConverted, err := strconv.Atoi(sNumber)
		if err != nil {
			continue
		}

		if len(sNumber) != 4 {
			fmt.Println("Please Enter number which has 4 digits")
			continue
		} else if numberConverted < 0 {
			fmt.Println("Please enter positive number")
			continue
		}
		sTargetNumber := strconv.Itoa(targetNumber)

		trueGuesses := 0
		falseGuesses := 0
		for i := 0; i < 4; i++ {
			foundIndex := strings.Index(sTargetNumber, string(sNumber[i]))
			if foundIndex != -1 {
				if foundIndex == i {
					trueGuesses++
				} else {
					falseGuesses++
				}
			}
		}
		guessed := make(map[string]string)
		guessed["entered"] = sNumber
		guessed["result"] = ""
		if trueGuesses > 0 {
			guessed["result"] += "+"
			guessed["result"] += strconv.Itoa(trueGuesses)

		}
		if falseGuesses > 0 {
			guessed["result"] += "-"
			guessed["result"] += strconv.Itoa(falseGuesses)

		}
		guessedNumbers = append(guessedNumbers, guessed)

		for _, item := range guessedNumbers {
			fmt.Printf("Entered Number : %s , Guess Result is :%s \n", item["entered"], item["result"])
		}
		if trueGuesses == 4 {
			fmt.Println("You have won!")
			return
		} else {

			fmt.Println("try again!")
		}

	}

}
