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
// EXERCISE: Print the runes
//
//  1. Loop over the "console" word and print its runes one by one,
//     in decimals, hexadecimals and binary.
//
//  2. Manually put the runes of the "console" word to a byte slice, one by one.
//
//     As the elements of the byte slice use only the rune literals.
//
//     Print the byte slice.
//
//  3. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only decimal numbers.
//
//  4. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only hexadecimal numbers.
//
//
// EXPECTED OUTPUT
//   Run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
	const word = "КОНзОла"
	var bs []byte
	var ds []int
	var hs []string
	for _,r := range word{
		fmt.Printf("%c\n",r)
		fmt.Printf("\t%-01s %d\n","decimal:",r)
		fmt.Printf("\thex: %-10x\n",r)
		fmt.Printf("\tbinary: %#b\n",r)

		bs = append(bs, byte(r))
		ds = append(ds, int(r))
		hs = append(hs, fmt.Sprintf("%10x",r))
	}

	fmt.Printf("BYTES: %v\n",bs)
	fmt.Printf("DECIMALS: %v\n",ds)
	fmt.Printf("HEXES: %v\n",hs)
}
