package main

import (
	"fmt"
	"math"
)

func main() {
	var value float64 = 0
	var operator string
	var number float64

	for operator != "x" {
		fmt.Print("Enter operation>")
		fmt.Scan(&operator, &number)
		calculator(operator, number, &value)
		fmt.Printf("Current value = %f\n", value)
	}

	fmt.Println("The final product was:", value)
}

func calculator(op string, num float64, val *float64) {
	switch op {
	case "+":
		plusFunc(num, val)
	case "-":
		minusFunc(num, &*val)
	case "*":
		timesFunc(num, &*val)
	case "/":
		divideFunc(num, &*val)
	case "^":
		powerFunc(num, &*val)
	default:
		fmt.Println("Invalid use of operator")

	}
}

func plusFunc(num float64, val *float64) {
	*val += num
}

func minusFunc(num float64, val *float64) {
	*val -= num
}

func timesFunc(num float64, val *float64) {
	*val *= num
}

func divideFunc(num float64, val *float64) {
	if num == 0 {
		fmt.Println("Cannot divide by zero")
	} else {
		*val /= num
	}
}

func powerFunc(num float64, val *float64) {
	// Own power of code:
	*val = math.Pow(*val, num)
}
