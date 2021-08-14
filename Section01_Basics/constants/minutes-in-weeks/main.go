// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Minutes in Weeks
//
//  Calculate how many minutes in two weeks.
//
// EXPECTED OUTPUT
//  There are 20160 minutes in 2 weeks.
// ---------------------------------------------------------

func main() {

	const (
		minPerDay = 24 * 60
		weekDays = 7
	)

	fmt.Printf("There are %d minutes in %d weeks.\n", 2 * (minPerDay * 7),2)
}
