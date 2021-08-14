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
// EXERCISE: Constant Length
//
//  Calculate how many characters inside the `home`
//  constant and print it.
//
//
// EXPECTED OUTPUT
//  There are 16 characters inside "Milky Way Galaxy"
// ---------------------------------------------------------

func main() {
	const (
		home = "Milky Way Galaxy"
		lenght = len(home)
	)

	fmt.Printf("There are %d characters inside %q\n", lenght,home)
}
