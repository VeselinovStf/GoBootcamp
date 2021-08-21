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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		headerSpred                          = strings.Split(header, separator)
		dataSpred                            = strings.Split(data, "\n")
		locations                            []string
		sizes, beds, baths, prices           []int
		headerSeparator                      = "="
		sumSize, sumBeds, sumBaths, sumPrice float64
	)

	for _, e := range dataSpred {
		line := strings.Split(e, separator)

		locations = append(locations, line[0])

		val, _ := strconv.Atoi(line[1])
		sizes = append(sizes, val)
		sumSize += float64(val)

		val, _ = strconv.Atoi(line[2])
		beds = append(beds, val)
		sumBeds += float64(val)

		val, _ = strconv.Atoi(line[3])
		baths = append(baths, val)
		sumBaths += float64(val)

		val, _ = strconv.Atoi(line[4])
		prices = append(prices, val)
		sumPrice += float64(val)
	}

	for _, e := range headerSpred {
		fmt.Printf("%-15s", e)
	}

	fmt.Println()
	fmt.Println(strings.Repeat(headerSeparator, len(headerSpred)*7))

	for i := range dataSpred {
		fmt.Printf("%-15s", locations[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}

	fmt.Println()
	fmt.Println(strings.Repeat(headerSeparator, len(headerSpred)*7))
	fmt.Printf("%-15s", " ")
	fmt.Printf("%-15.2f", sumSize/float64(len(sizes)))
	fmt.Printf("%-15.2f", sumBeds/float64(len(beds)))
	fmt.Printf("%-15.2f", sumBaths/float64(len(baths)))
	fmt.Printf("%-15.2f", sumPrice/float64(len(prices)))

}
