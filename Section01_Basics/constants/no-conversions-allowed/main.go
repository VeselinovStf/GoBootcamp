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
	"time"
)

// ---------------------------------------------------------
// EXERCISE: No Conversions Allowed
//
//  1. Fix the program without doing any conversion.
//  2. Explain why it doesn't work.
//
// EXPECTED OUTPUT
//  10h0m0s later...
// ---------------------------------------------------------

func main() {
	const later int = 10

	hours, _ := time.ParseDuration("1h")

	// can't multiple later (int ) with hours (time.Duration) they are different types
	fmt.Printf("%s later...\n", hours * time.Duration(later))
}
