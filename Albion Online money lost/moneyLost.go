package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	var input int
	// Asks for the amount of money you lost.
	fmt.Printf("Enter how much money you lost> ")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}

	// Reads the number in the albionMoneyLost text file and assigns it to "data" as byte
	data, err2 := ioutil.ReadFile("albionMoneyLost.txt")
	if err2 != nil {
		log.Fatal(err2)
	}

	// We convert "data" as Byte to String, after which it is converted further to an Integer and stored in "tempt".
	tempt, err3 := strconv.Atoi(string(data))
	if err3 != nil {
		log.Fatal(err3)
	}

	// The value of "input" is then added to tempt, after which it is converted to a string, so that it can be converted to byte-type.
	tempt = tempt + input
	t := strconv.Itoa(tempt) // Since "tempt" is an integer-type, we have to create a new variable that is either already a string or can become a string.
	byteData := []byte(t)

	// The converted numbers are then added to the albionMoneyLost text file.
	err4 := ioutil.WriteFile("albionMoneyLost.txt", byteData, 0777)
	if err4 != nil {
		log.Fatal(err4)
	}
}
