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
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Please enter min and max argumeters")
		return
	}

	min , err1 := strconv.Atoi(args[0])
	max , err2:= strconv.Atoi(args[1])
	if err1 != nil || err2 != nil || min > max {
		fmt.Println("Please enter valid number")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		if i % 2 != 0 {
			continue
		}
		sum+=i
		fmt.Printf("%d", i)
		if i != max {
			fmt.Print(" + ")
		}
	}

	fmt.Printf(" = %d\n", sum)
}
