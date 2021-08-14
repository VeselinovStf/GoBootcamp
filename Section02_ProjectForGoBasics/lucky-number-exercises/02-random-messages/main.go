// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Enter a single number argument!")
		return
	}

	playerN, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Enter a number")
		return
	}

	const gameCheat = 2

	rand.Seed(time.Now().UnixNano())

	if rand.Intn(gameCheat) == playerN {
		switch rand.Intn(3) {
		case 0:
			fmt.Println("🎉  YOU WIN!")
		case 1:
			fmt.Println("🎉  YOU'RE AWESOME!")
		case 2:
			fmt.Println("🎉  PERFECT!")
		}
		return

	}

	msg := "%s Try again?\n"

	switch rand.Intn(2) {
	case 0:
		fmt.Printf(msg, "☠️  YOU LOST...")
	case 1:
		fmt.Printf(msg, "☠️  JUST A BAD LUCK...")
	}
}
