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
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
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

	type export struct {
		ID    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Price int    `json:"price,omitempty"`
		Genre string `json:"genre,omitempty"`
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
			if len(t) != 2 {
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
		case "save":
			var data []export
			for k,v := range games{
				data = append(data, export{ID: k, Name: v.name,Price: v.price, Genre: v.genre})
			}
			json,err := json.MarshalIndent(data,"","\t")
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
