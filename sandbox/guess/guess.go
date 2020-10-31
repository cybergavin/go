// Guessing game - Challenge a user to guess a number from 1 to 100 --

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("LET'S PLAY A GUESSING GAME! YOU WILL HAVE 10 GUESSES to GUESS A RANDOM NUMBER FROM 1 TO 100!")
	fmt.Println()
	playAgain := true
	for playAgain == true {
		rand.Seed(time.Now().UnixNano())
		targetNum := rand.Intn(100) + 1
		success := false
		reader := bufio.NewReader(os.Stdin)
		for guesses := 1; guesses <= 10; guesses++ {
			fmt.Printf("Enter Guess #%d : ", guesses)
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			input = strings.TrimSpace(input)
			guess, err := strconv.Atoi(input)
			if err != nil {
				log.Fatal(err)
			}
			if guess < targetNum {
				fmt.Println("OOPS! Your guess", guess, "is lesser than the random number!")
			} else if guess > targetNum {
				fmt.Println("OOPS! Your guess", guess, "is greater than the random number!")
			} else {
				success = true
				fmt.Println("CORRECT! You guessed the number in ", guesses, "attempt(s)!")
				break
			}
		}
		if !success {
			fmt.Println("GAME OVER! Your did not guess the random number. It was", targetNum)
		}
		fmt.Println()
		fmt.Print("Do you want to play again (Y/N) ? ")
		repeatGame, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		repeatGame = strings.TrimSpace(repeatGame)
		if repeatGame == "Y" {
			playAgain = true
			fmt.Println("Ok, here we go again!")
			fmt.Println()
		} else if repeatGame == "N" {
			playAgain = false
			fmt.Println("Ok, thank you for playing. Goodbye!")
		} else {
			log.Fatal("Invalid choice! Expecting Y or N ! Goodbye!")
		}
	}
}
