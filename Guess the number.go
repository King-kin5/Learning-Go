package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	min, max := 1, 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(max-min) + min

	fmt.Println("Guess a number between 1 and 100")
	fmt.Println("Please input your guess")

	attempts := 0

	for {
		attempts++

		var userGuess int
		_, err := fmt.Scan(&userGuess)

		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		fmt.Println("Your guess is", userGuess)

		if userGuess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Try again")
		} else if userGuess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Try again")
		} else {
			fmt.Println("Correct, you Legend! You guessed right after", attempts, "attempts")
			break
		}
	}
}
