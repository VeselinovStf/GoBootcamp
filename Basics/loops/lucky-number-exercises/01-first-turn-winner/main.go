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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
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
		fmt.Println("You win")
		return
	}

	fmt.Println("Try again")
}
