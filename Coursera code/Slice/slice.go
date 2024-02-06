package main

import (
	"fmt"
	"strconv"
)

func main() {
	sli := make([]int, 0, 3)
	var input string = ""

	for {
		// Asks user for integer, which will later be converted
		fmt.Printf("Enter your integers> ")
		fmt.Scan(&input)

		// The loop will end
		if input == "x" {
			break
		}

		// In case of error
		pass, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		// Adds elements to the slice sli and sorts it.
		sli = addInt(sli, pass)
		sli = sortArr(sli)

		fmt.Printf("%d ", sli)
		fmt.Printf("Len=%d, cap=%d\n", len(sli), cap(sli))
	}
}

// Add integer to slice.
func addInt(x []int, i int) []int {
	x = append(x, i)
	return x
}

// Sorts the slice via. Insertion Sort.
func sortArr(x []int) []int {
	for i := 1; i < len(x); i++ {
		var curr int = x[i]
		var j int = i - 1

		for j > -1 && curr < x[j] {
			x[j+1] = x[j]
			j--
		}
		x[j+1] = curr
	}
	return x
}
