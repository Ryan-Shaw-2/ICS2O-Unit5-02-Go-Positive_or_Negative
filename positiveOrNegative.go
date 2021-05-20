// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program checks if the user’s number is positive or negative

package main

import "fmt"

func main() {
	// This function checks if the user’s number is positive or negative
	var userNumber int

	// input
	fmt.Println("This program checks if the user’s number is positive or negative")
	fmt.Println()
	fmt.Print("Enter your number: ")
	fmt.Scanln(&userNumber)
	fmt.Println()

	// process
	if userNumber < 0 {
		// output
		fmt.Println("That is a negative number")
	} else {
		// output
		fmt.Println("That is a positve number")
	}
}
