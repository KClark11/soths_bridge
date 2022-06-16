package main

// Can you guess how much gold it take's to cross the knight's bridge? Guess 1 - 100 and see!

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

	seconds := time.Now().Unix() // Get current time and date as an integer
	rand.Seed(seconds)           //Seed the random number generator
	target := rand.Intn(100) + 1
	fmt.Println("Before you enter the castle, you must cross the bridge. There is a knight guarding it!")
	fmt.Println(" 'To cross my bridge you must pay the toll in gold!' ")
	fmt.Println("Can you guess how much gold it costs?")
	//fmt.Println(target)

	reader := bufio.NewReader(os.Stdin) // Let's program read input from keyboard

	success := false // set up to print failure by default.

	for guesses := 0; guesses <= 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left!")

		fmt.Println("Make your guess: ")
		input, err := reader.ReadString('\n') // reads typed input from keyboard until ENTER is pressed
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("NO! You're guess is too LOW!")
		} else if guess > target {
			fmt.Println("NO! You're guess is too HIGH!")
		} else {
			success = true
			fmt.Println("Splendid! You guessed corectly!")
			fmt.Println("YOU WIN!")
			break
		}

	}
	if !success {
		fmt.Println("Sorry, you didn't guess correctly. It was:", target, "You can't cross my bridge!")
		fmt.Println("GAME OVER")
	}

}
