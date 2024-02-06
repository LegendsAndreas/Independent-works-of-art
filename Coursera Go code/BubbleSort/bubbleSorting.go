package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	intSli := make([]int, 0)
	var x string
	var i int

	// Asks user to enter numbers to sort, until 10 numbers has been entered, or the user write 'x'. 'x' is always converted to lower, in case the user enters and uppercase 'x'
	for i < 10 {
		fmt.Printf("Enter up to 10 number(x to stop)> ")
		fmt.Scan(&x)

		if strings.ToLower(x) == "x" {
			break
		}

		x, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}

		intSli = append(intSli, x)
		i++
	}

	// Sorts the slice
	bubbleSort(intSli)
}

// Sorts the slice via. bubble sort of all things.
func bubbleSort(sli []int) {
	for i := 0; i < len(sli); i++ {

		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				swap(sli, j)
			}
		}
	}
	fmt.Println(sli)
}

// In case 'i' is bigger than 'i + 1', the 2 will switch.
func swap(sli []int, i int) {
	tempt := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = tempt
}
