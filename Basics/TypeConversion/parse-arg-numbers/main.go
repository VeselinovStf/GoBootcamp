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
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Parse Arg Numbers
//
//  Use strconv.ParseInt function to get int8, int16, and
//  int32, and int64 values from command-line.
//
// HINT
//  The third argument to ParseInt function represents
//  the bitsize.
//
//  So, giving it 8 returns an int8 convertable value;
//  whereas 16 returns an int16 convertable value.
//
//  Please explore the documentation of ParseInt function
//  and learn how it works.
//
// EXPECTED OUTPUT
//   When runned like this:
//     go run main.go 50 25000 2000000 50000000000000000 00000100
//
//   It should return this:
//     int8 value is : 50
//     int16 value is: 25000
//     int32 value is: 2000000
//     int64 value is: 50000000000000000
//     00000100 is: 4
// ---------------------------------------------------------

func main() {

	val, _ := strconv.ParseInt(os.Args[1], 10, 8)
	fmt.Println("int8 value is:", int8(val))

	// 1. Get an int16 value using ParseInt and convert it and print it
	val, _ = strconv.ParseInt(os.Args[2], 10, 16)
	fmt.Println("int16 value is:", int16(val))

	// 2. Get an int32 value using ParseInt and convert it and print it
	val, _ = strconv.ParseInt(os.Args[3], 10, 32)
	fmt.Println("int32 value is:", int32(val))

	// 3. Get an int64 value using ParseInt and convert it and print it
	val, _ = strconv.ParseInt(os.Args[4], 10, 64)
	fmt.Println("int64 value is:", int64(val))

	// 4. Get an int8 value using ParseInt and convert it and print it
	val, _ = strconv.ParseInt(os.Args[5], 2, 64)
	fmt.Printf("%s is: %d\n", os.Args[5], val)

}
