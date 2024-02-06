package main

import "fmt"

func main() {
	//Asks user for a float to truncate.
	var floatToInt float32
	fmt.Print("Enter a floating point number> ")
	fmt.Scan(&floatToInt)

	//By converting it to an integer, we effectively truncate the floating point number.
	var inter int32 = int32(floatToInt)
	fmt.Printf("Your float point has been changed to: %d", inter)
}
