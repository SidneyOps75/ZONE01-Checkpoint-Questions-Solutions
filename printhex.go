// printhex
// Instructions

// Write a program that takes a positive (or zero) number expressed in base 10, and that displays it in base 16 (with lowercase letters) followed by a newline ('\n').

//     If the number of arguments is different from 1, the program displays nothing.
//     Error cases have to be handled as shown in the example below.

// Usage

// $ go run . 10
// a
// $ go run . 255
// ff
// $ go run . 5156454
// 4eae66
// $ go run .
// $ go run . "123 132 1" | cat -e
// ERROR$
// $

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	// Check the number of arguments
	if len(os.Args) != 2 {
		return // Exit if the number of arguments is not 2
	}

	// Parse the input argument
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR") // Print "ERROR" if argument parsing fails
		return
	}

	// Convert to hexadecimal manually
	hexStr := ""
	for input > 0 {
		digit := input % 16
		if digit < 10 {
			hexStr = string('0'+digit) + hexStr // Add digit to hexStr
		} else {
			hexStr = string('a'+digit-10) + hexStr // Add letter to hexStr
		}
		input /= 16 // Move to the next digit
	}

	// Print the hexadecimal representation
	for _, c := range hexStr {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n') // Print a newline character
}

