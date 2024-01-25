package main

import (
	"fmt"
	"log"
)

func main() {
	var input string

	fmt.Print("Enter a string you wish to translate to binary> ")
	_, err := fmt.Scan(&input)

	if err != nil {
		log.Fatal(err)
	}

}
