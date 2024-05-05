package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var value float64 = 0
	var operator string
	var number float64

	for operator != "x" {
		fmt.Print("Enter operation>")
		_, err := fmt.Scan(&operator, &number)
		if err != nil {
			log.Fatal("Input error:", err)
		}

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
	case "set":
		setFunc(num, &*val)
	case "sqrt":
		sqrtFunc(num, &*val)
	default:
		fmt.Println("Invalid use of operator")

	}
}

func plusFunc(num float64, val *float64)  { *val += num }
func minusFunc(num float64, val *float64) { *val -= num }
func timesFunc(num float64, val *float64) { *val *= num }
func powerFunc(num float64, val *float64) { *val = math.Pow(*val, num) }
func setFunc(num float64, val *float64)   { *val = num }
func sqrtFunc(num float64, val *float64)  { *val = math.Sqrt(*val) }
func divideFunc(num float64, val *float64) {
	if num == 0 {
		fmt.Println("Cannot divide by zero")
	} else {
		*val /= num
	}
}
