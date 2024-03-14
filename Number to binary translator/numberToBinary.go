package main

import "fmt"

func main() {
	var number int = 69
	fmt.Println(numToBi(number))

}

func numToBi(num int) string {
	binaryStr := ""
	hun := 128
	sixtyfour := 64
	thirtytwo := 32
	sixteen := 16
	eight := 8
	four := 4
	two := 2
	one := 1

	if num >= hun {
		binaryStr += "1"
		num -= hun
	} else {
		binaryStr += "0"
	}

	if num >= sixtyfour {
		binaryStr += "1"
		num -= sixtyfour
	} else {
		binaryStr += "0"
	}

	if num >= thirtytwo {
		binaryStr += "1"
		num -= thirtytwo
	} else {
		binaryStr += "0"
	}

	if num >= sixteen {
		binaryStr += "1"
		num -= sixteen
	} else {
		binaryStr += "0"
	}

	if num >= eight {
		binaryStr += "1"
		num -= eight
	} else {
		binaryStr += "0"
	}

	if num >= four {
		binaryStr += "1"
		num -= four
	} else {
		binaryStr += "0"
	}

	if num >= two {
		binaryStr += "1"
		num -= two
	} else {
		binaryStr += "0"
	}

	if num >= one {
		binaryStr += "1"
		num -= one
	} else {
		binaryStr += "0"
	}

	return binaryStr
}
