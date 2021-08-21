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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

type jsonGame struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"`
	Genre string `json:"genre,omitempty"`
}

func main() {

		// load the initial data from json
	var decoded []jsonGame
	if err := json.Unmarshal([]byte(data), &decoded); err != nil {
		fmt.Println("Sorry, there is a problem:", err)
		return
	}

	// load the data into usual game values
	var games []game
	for _, dg := range decoded {
		games = append(games, game{item{dg.ID, dg.Name, dg.Price}, dg.Genre})
	}
	// ----------------------------------------------------------

	// index the games by id
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

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
			if len(t) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(t[1])
			if err != nil {
				fmt.Println(" wrong id")
				continue
			}
			if g, ok := byID[id]; ok {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
				continue
			}

			fmt.Println("sorry. I don't have the game")
		case "save":
			var data []jsonGame
			for k, v := range games {
				data = append(data, jsonGame{ID: k, Name: v.name, Price: v.price, Genre: v.genre})
			}
			json, err := json.MarshalIndent(data, "", "\t")
			if err != nil {
				fmt.Println("Sorry:", err)
				continue
			}

			fmt.Println(string(json))
			return

		default:
			fmt.Println("Unknown Command")
		}

	}
}
