package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var input string

	// Asks user for a binary sequence that they wish to convert into an actual number.
	fmt.Print("Enter binary to convert to number (you are punching in a binary sequence, please for the love of God dont enter anything but 1's and 0's)> ")
	_, err := fmt.Scan(&input)

	// In case some goofy ahh mf don't actually enters a string. Should be pretty hard to do, considering it is a string.
	if err != nil {
		log.Fatal(err)
	}

	output := binaryToDecimal(input)

	// The output value is changes to int, so that we don't get a bunch of zeros at the end.
	fmt.Printf("The binary %s in numbers are: %d", input, int(output))
}

/**
"biStr" is the binary string sequence that the user wishes to convert into a decimal number. Method for conversion was found from the fine folk over at GeeksForGeeks: https://www.geeksforgeeks.org/binary-to-decimal/
Returns the number equivalent to the binary sequence.
 */
func binaryToDecimal(biStr string) float64 {
	decimal := 0.0
	sLength := len(biStr)
	var biNum int

	// It starts at the end of the string and makes it way towards the start of it.
	for i := 0; i < sLength; i++ {
		if biStr[sLength-1-i] == '0' {
			biNum = 0
		} else {
			biNum = 1
		}

		// Formula is: 2^i * biNum
		tempt := math.Pow(2, float64(i)) * float64(biNum)
		decimal += tempt
	}

	return decimal
}
