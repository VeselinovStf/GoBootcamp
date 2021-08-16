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
// EXERCISE: Append #2
//
//  1. Create the following nil slices:
//     + Pizza toppings
//     + Departure times
//     + Student graduation years
//     + On/off states of lights in a room
//
//  2. Append them some elements (use your creativity!)
//
//  3. Print all the slices
//
//
// EXPECTED OUTPUT
// (Your output may change, mine is like so:)
//
//  pizza       : [pepperoni onions extra cheese]
//
//  graduations : [1998 2005 2018]
//
//  departures  : [2019-01-28 15:09:31.294594 +0300 +03 m=+0.000325020
//  2019-01-29 15:09:31.294594 +0300 +03 m=+86400.000325020
//  2019-01-30 15:09:31.294594 +0300 +03 m=+172800.000325020]
//
//  lights      : [true false true]
//
//
// ---------------------------------------------------------

func main() {
	var (
		pizzaToppings        []string
		departureTimes       []time.Time
		studentGratuateYears []int
		lightsStatus         []bool
	)

	pizzaToppings = append(pizzaToppings, "chees", "fish", "salamy")

	now := time.Now()
	departureTimes = append(departureTimes, now.Add(time.Hour*24), now.Add(time.Hour*24))
	studentGratuateYears = append(studentGratuateYears, 2021, 1998, 1999)
	lightsStatus = append(lightsStatus, false, true, false)

	fmt.Printf("pizza	: %s\n", pizzaToppings)
	fmt.Printf("gratuation	: %d\n", studentGratuateYears)
	fmt.Printf("departure	: %s\n", departureTimes)
	fmt.Printf("lights	: %t\n", lightsStatus)

}
