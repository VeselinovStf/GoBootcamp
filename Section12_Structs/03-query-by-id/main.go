// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. I don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
// ---------------------------------------------------------

func main() {
	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	games := map[int]game{
		1: {
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		2: {item: item{id: 2, name: "x-com 2", price: 40}, genre: "strategy"},
		3: {item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}

		t := strings.Fields(scanner.Text())

		switch t[0] {
		case "quit":
			return
		case "list":
			fmt.Printf("Inanc's game store has %d games.\n\n", len(games))

			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}
		case "id":
			if len(t) !=  2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(t[1])
			if err != nil {
				fmt.Println(" wrong id")
				continue
			}
			if g, ok := games[id]; ok {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
				continue
			}

			fmt.Println("sorry. I don't have the game")
		default:
			fmt.Println("Unknown Command")
		}

	}
}
