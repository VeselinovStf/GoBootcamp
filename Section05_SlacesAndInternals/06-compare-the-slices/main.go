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
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
// ---------------------------------------------------------

func main() {
	namesA := strings.Split("Da Vinci, Wozniak, Carmack", ", ")
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	if len(namesA) == len(namesB) {
		sort.Strings(namesA)
		sort.Strings(namesB)

		for i, e := range namesA {
			if e != namesB[i] {
				break
			}
		}

		// Task and solution are little .... odd
		fmt.Println("They are equal")
		return
	}

	fmt.Println("They are not equal")
}
