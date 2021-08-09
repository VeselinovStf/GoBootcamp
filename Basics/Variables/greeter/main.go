// Package main implements
// Greeter main starting point
package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	argumentsLength := len(arguments)
	fmt.Printf("Args Lenght: %d\n", len(arguments))
	fmt.Printf("Path: %s\n", arguments[0])
	for i := 1; i < argumentsLength; i++ {
		fmt.Printf("Argumeter - %d : %s\n", i, arguments[i])
	}
}
