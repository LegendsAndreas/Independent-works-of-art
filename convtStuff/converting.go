package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter conversion form> ") // Example: "Int to float"
	x := bufio.NewScanner(os.Stdin)
	var input string
	if x.Scan() {
		input = x.Text()
	} else {
		log.Fatal("An error occurred during input.")
	}

	input = strings.ToLower(input)

	var convrtFrom string
	var convrtTo string
	insertConversionTypes(input, &convrtFrom, &convrtTo)

	var enteredValue int = 2
	var convertedValue float64

	if convrtFrom == "int" && convrtTo == "float" {
		convertedValue = intToFloat(enteredValue)
	} else if convrtFrom == "float" && convrtTo == "int" {
		convertedValue = floatToInt(enteredValue) // I believe if i make an interface, it can work.
	}

	fmt.Println(convertedValue)

	// Hey guuuuuuuurl!

	fmt.Println(convertedValue)

}

/*
*
This is one of the dumbest ways I could have done this, but it was pretty fun and quite the learning experience.
*/
func insertConversionTypes(inp string, convert *string, converted *string) {
	var inpLength int = len(inp)

	for i := 0; inp[i] != ' '; i++ {
		if inp[i] != ' ' {
			*convert = *convert + string(inp[i])
		}
	}

	for i := inpLength - 1; inp[i] != ' '; i-- {
		if inp[i] != ' ' {
			*converted = *converted + string(inp[i])
		}
	}

	*converted = backwards(*converted)

}

func backwards(s string) string {
	var tempt string
	for i := len(s) - 1; i >= 0; i-- {
		tempt = tempt + string(s[i])
	}

	return tempt
}

func intToFloat(x int) float64 {
	// x = float64(x) This does not work, because "x" is technically an int and so assigning it a float value fucks it.
	a := float64(x)
	return a
}

func floatToInt(x float64) int {
	a := int(x)
	return a
}
