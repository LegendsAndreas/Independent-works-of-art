package main

import "fmt"

func main() {
	var number int = 240
	fmt.Println(numToBi(number))
}

// numToBi converts a decimal number to its binary representation as a string.
// If the number is greater than 255 or less than 0, it returns an error message.
func numToBi(num int) string {
	if num > 255 || num < 0 {
		return "It is a regular 8 byte string buddy, dont go over 255 or below 0."
	}

	binaryStr := "" // The binary number that will be returned, as a string.

	digitEight := 128
	digitSeven := 64
	digitSix := 32
	digitFive := 16
	digitFour := 8
	digitThree := 4
	digitTwo := 2
	digitOne := 1

	// If 'num' is bigger than the appropriate digit, the appropriate digit value will be subtracted from 'num' and
	// continue to the next digit and a '1' will be added to 'binaryStr'.
	// Else, a '0' will be added.
	if num >= digitEight { // 128
		binaryStr += "1"
		num -= digitEight
	} else {
		binaryStr += "0"
	}

	if num >= digitSeven { // 64
		binaryStr += "1"
		num -= digitSeven
	} else {
		binaryStr += "0"
	}

	if num >= digitSix { // 32
		binaryStr += "1"
		num -= digitSix
	} else {
		binaryStr += "0"
	}

	if num >= digitFive { // 16
		binaryStr += "1"
		num -= digitFive
	} else {
		binaryStr += "0"
	}

	if num >= digitFour { // 8
		binaryStr += "1"
		num -= digitFour
	} else {
		binaryStr += "0"
	}

	if num >= digitThree { // 4
		binaryStr += "1"
		num -= digitThree
	} else {
		binaryStr += "0"
	}

	if num >= digitTwo { // 2
		binaryStr += "1"
		num -= digitTwo
	} else {
		binaryStr += "0"
	}

	if num >= digitOne { // 1
		binaryStr += "1"
		num -= digitOne
	} else {
		binaryStr += "0"
	}

	return binaryStr
}
