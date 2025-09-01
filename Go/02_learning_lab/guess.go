package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// TASK: Generate a random number guessing game.

// Generate a random number between 1 and 100 and store it.
//
// Ask the player to guess the number and store their answer.
//
// If the player's guess is lower than the secret number,
// print a message (e.g., "Oops. Your guess was LOW").
//
// If the player's guess is higher than the secret number,
// print another message (e.g., "Oops. Your guess was HIGH").
//
// Allow the player up to 10 guesses. Before each guess,
// inform the player how many attempts they have left.
//
// If the player's guess equals the secret number,
// print a success message and stop asking for guesses.
//
// If the player runs out of attempts without guessing the number,
// print the message: "Sorry. You didn’t guess my number. It was: [the secret number]".

func main() {
	// Generate a random number between 1 and 100
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1

	reader := bufio.NewReader(os.Stdin)

	// Allow the player up to 10 guesses
	for attempt := 1; attempt <= 10; attempt++ {
		remaining := 10 - attempt + 1
		fmt.Printf("You have %d attempts left. Enter your guess: ", remaining)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			attempt-- // do not count invalid input as an attempt
			continue
		}

		// Compare the guess with the target
		if guess == target {
			fmt.Printf("Congratulations! You guessed the number: %d\n", guess)
			return
		} else if guess < target {
			fmt.Println("Oops. Your guess was TOO LOW.")
		} else {
			fmt.Println("Oops. Your guess was TOO HIGH.")
		}
	}

	// If the player runs out of attempts
	fmt.Printf("Sorry, you didn’t guess my number. It was %d.\n", target)
}
