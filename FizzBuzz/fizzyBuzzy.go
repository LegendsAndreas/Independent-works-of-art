package main

import (
	"fmt"
	"log"
	"time"
)

var arrowUp = "\u2191"
var arrowDown = "\u2193"

const MAX = 100

func main() {
	var arr [MAX]int
	for i := 0; i < MAX; i++ {
		arr[i] = i + 1
	}

	//autoFizzBuzz(arr)

	var input string
	var points int
	start := time.Now()
	for _, val := range arr {
		fmt.Printf("Fizz? Buzz? Both? Or next? %d", val)
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}

		if input == "fizz" {
			if isFizz(val) {
				fmt.Println("Correct!", arrowUp, arrowUp, arrowUp)
				points++
			} else {
				fmt.Println("Wrong!", arrowDown, arrowDown, arrowDown)
				points--
			}

		} else if input == "buzz" {
			if isBuzz(val) {
				fmt.Println("Correct!", arrowUp, arrowUp, arrowUp)
				points++
			} else {
				fmt.Println("Wrong!", arrowDown, arrowDown, arrowDown)
				points--
			}

		} else if input == "both" {
			if isBoth(val) {
				fmt.Println("Correct!", arrowUp, arrowUp, arrowUp)
				points++
			} else {
				fmt.Println("Wrong!", arrowDown, arrowDown, arrowDown)
				points--
			}

		} else if input == "next" {
			if next(val) {
				fmt.Println("Correct!", arrowUp, arrowUp, arrowUp)
				points++
			} else {
				fmt.Println("Wrong!", arrowDown, arrowDown, arrowDown)
				points--
			}
		} else {
			fmt.Println("Invalid input, MINUS POINTS!", arrowDown, arrowDown, arrowDown)
			points--
		}
	}

	timeElapsed := time.Since(start).Seconds()
	fmt.Println("You got: ", points, " in ", timeElapsed, " seconds")
}

// Plays Fizz Buzz automatically.
func autoFizzBuzz(arr [100]int) {
	for i := range arr {
		if arr[i]%3 == 0 && arr[i]%5 == 0 {
			fmt.Printf("%d:", arr[i])
			fmt.Println("Fizz Buzz!")
		} else if arr[i]%3 == 0 {
			fmt.Printf("%d:", arr[i])
			fmt.Println("Fizz!")
		} else if arr[i]%5 == 0 {
			fmt.Printf("%d:", arr[i])
			fmt.Println("Buzz!")
		}
	}
}

func isFizz(x int) bool {
	if x%3 == 0 && x%5 != 0 {
		return true
	} else {
		return false
	}
}

func isBuzz(x int) bool {
	if x%5 == 0 && x%3 != 0 {
		return true
	} else {
		return false
	}
}

func isBoth(x int) bool {
	if x%3 == 0 && x%5 == 0 {
		return true
	} else {
		return false
	}
}

func next(x int) bool {
	if x%3 != 0 && x%5 != 0 {
		return true
	} else {
		return false
	}
}
