package main

import (
	"fmt"
	"log"
	"math"
)

type Input struct {
	operator string
	number   float64
}

func main() {
	var inputs Input
	var value float64 = 0

	for inputs.operator != "x" {
		fmt.Print("Enter operation>")
		_, err := fmt.Scan(&inputs.operator, &inputs.number)
		if err != nil {
			log.Fatal(err)
		}

		inputs.Calculator(&value)
		fmt.Printf("Current value = %f\n", value)
	}

	fmt.Println("The final product was:", value)
}

func (inputs Input) Calculator(val *float64) {
	switch inputs.operator {
	case "+":
		inputs.Plus(&*val)
	case "-":
		inputs.Minus(&*val)
	case "*":
		inputs.Times(&*val)
	case "set":
		inputs.Set(&*val)
	case "^":
		inputs.Power(&*val)
	case "sqrt":
		inputs.Sqrt(&*val)
	case "/":
		inputs.Divide(&*val)
	default:
		fmt.Println("Invalid operator")
	}
}

// In Golang, you have to pass a custom type to a method.

func (inputs Input) Plus(val *float64)  { *val += inputs.number }
func (inputs Input) Minus(val *float64) { *val -= inputs.number }
func (inputs Input) Times(val *float64) { *val *= inputs.number }
func (inputs Input) Power(val *float64) { *val = math.Pow(*val, inputs.number) }
func (inputs Input) Sqrt(val *float64)  { *val += math.Sqrt(*val) }
func (inputs Input) Set(val *float64)   { *val = inputs.number }
func (inputs Input) Divide(val *float64) {
	if *val == 0 {
		fmt.Println("Cannot divide by zero")
	} else {
		*val /= inputs.number
	}
}
