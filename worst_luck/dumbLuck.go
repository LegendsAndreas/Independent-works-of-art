package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var one int
	var two int
	var three int
	var four int
	var five int
	var six int
	var seven int
	var eight int
	var nine int
	var ten int

	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	fmt.Println(arr)

	// For 1,000 iterations, we iterate through "arr", until we randomly get a number from 0-2 (exclusive), that is equal to 1.
	// Then, one integer is the appropriate variable.
	var throw int
	for i := 0; i < 1_000; i++ {
		for _, val := range arr {
			throw = rand.Intn(2)
			if throw == 1 {
				if val == 1 {
					one++
				} else if val == 2 {
					two++
				} else if val == 3 {
					three++
				} else if val == 4 {
					four++
				} else if val == 5 {
					five++
				} else if val == 6 {
					six++
				} else if val == 7 {
					seven++
				} else if val == 8 {
					eight++
				} else if val == 9 {
					nine++
				} else if val == 10 {
					ten++
				}
				break
			}
		}
	}

	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)
	fmt.Println(five)
	fmt.Println(six)
	fmt.Println(seven)
	fmt.Println(eight)
	fmt.Println(nine)
	fmt.Println(ten)

	if ten > 0 {
		fmt.Println("DING-DING-DING-DING- WE-GOT-A-WINNER!")
	}
}
