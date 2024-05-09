package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Used to keep track of the time it took for a player to complete a game of Fizz Buzz and the amount of points they got
type score struct {
	time  int
	point int
}

var arrowUp = "\u2191"
var arrowDown = "\u2193"

const MAX = 15

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

		// The player guesses "fizz"
		if input == "fizz" {
			if isFizz(val) {
				fmt.Println("Correct!", arrowUp, arrowUp, arrowUp)
				points++
			} else {
				fmt.Println("Wrong!", arrowDown, arrowDown, arrowDown)
				points--
			}

			// The player guesses "buzz"
		} else if input == "buzz" {
			if isBuzz(val) {
				fmt.Println("Correct!", arrowUp, arrowUp, arrowUp)
				points++
			} else {
				fmt.Println("Wrong!", arrowDown, arrowDown, arrowDown)
				points--
			}

			// The player guesses "both"
		} else if input == "both" {
			if isBoth(val) {
				fmt.Println("Correct!", arrowUp, arrowUp, arrowUp)
				points++
			} else {
				fmt.Println("Wrong!", arrowDown, arrowDown, arrowDown)
				points--
			}

			// The player guesses "next"
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

	// Records the time it took for the player to complete a game of Fizz Buzz and converting it to a string in seconds.
	timeElapsed := strconv.FormatInt(int64(time.Since(start).Seconds()), 10)

	// Opens the text file 'FizzBuzzScoreboard.txt', used for keeping track of player scores.
	scoreboardFile, err := os.OpenFile("FizzBuzzScoreboard.txt", os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// We convert 'points' from integer to string, so that we can store it in the scoreboardFile.
	convertedPoints := strconv.FormatInt(int64(points), 10)

	// We write the time it took for the player to complete and points gained into the scoreboardFile.
	_, err2 := scoreboardFile.WriteString(timeElapsed + " " + convertedPoints + "\n")
	if err2 != nil {
		log.Fatal(err2)
	}

	// We are now done with the file, so we close it.
	err3 := scoreboardFile.Close()
	if err != nil {
		log.Fatal(err3)
	}

	// Prints the players final score.
	fmt.Println("In", timeElapsed, "seconds, you got:", points, "points!")

	// Gets the scores of all the games and stores them in 'scores'.
	var scores []score
	scores = getScores(scores)
	for i := range scores {
		fmt.Printf("Score %d is: %d and %d\n", i, scores[i].time, scores[i].point)
	}

	//
	var highScores = scores
	highScores = getHighScores(highScores)

	for i := range highScores {
		fmt.Printf("High score %d is: %d and %d\n", i+1, highScores[i].time, highScores[i].point)
	}
}

// getHighScores sorts the highScoresArr in descending order based on the point field of each score struct.
// It uses the insertion sort algorithm to accomplish this.
// The function takes in a slice of score structs as input and returns the sorted slice.
func getHighScores(highScoresArr []score) []score {
	var i int
	var j int

	for i = 1; i < len(highScoresArr); i++ {
		elem := highScoresArr[i].point // Starts at 1
		j = i - 1                      // Starts at 0

		for j >= 0 && highScoresArr[j].point < elem {
			highScoresArr[j+1].point = highScoresArr[j].point
			j = j - 1
		}
		highScoresArr[j+1].point = elem
	}

	return highScoresArr
}

// Plays Fizz Buzz automatically.
func autoFizzBuzz(arr [MAX]int) {
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

func getScores(scores []score) []score {

	// Since we just have to open a file, we can just write: os.Open([FILE-NAME]), because it is read-only by default.
	scoreFile, err := os.Open("FizzBuzzScoreboard.txt")
	if err != nil {
		log.Fatal("getScores error 1:", err)
	}

	fileScanner := bufio.NewScanner(scoreFile)
	fileScanner.Split(bufio.ScanLines) // This basically tells the scanner to read the file, until it comes to a newline.

	// As long as we have not read through the entire file, the loop will continue.
	for fileScanner.Scan() {
		// Reads the file, one line at a time and stores the content in 'line'.
		line := fileScanner.Text()
		// 'strings.Split(line, " ")' takes the variable 'line' and splits it into a slice of substrings, breaking at each space.
		// It then takes element 0 of that slice, which represents time in seconds, and stores it in 'time'.
		// Lastly, it converts the string into an integer.
		time, err2 := strconv.Atoi(strings.Split(line, " ")[0])
		if err2 != nil {
			log.Fatal("getScores error 2, conversion error:", err2)
		}
		// Same as with the 'time' variable, only it stores element 1 of the slice of substrings.
		point, err3 := strconv.Atoi(strings.Split(line, " ")[1])
		if err3 != nil {
			log.Fatal("getScores error 3, conversion error:", err3)
		}
		// Finally, appends 'time' and 'point' as a score-struct, to the slice 'scores'.
		scores = append(scores, score{time, point})
	}

	// We are now done with the file, so we close it.
	err2 := scoreFile.Close()
	if err2 != nil {
		log.Fatal("getScores error 2:", err2)
	}
	return scores
}
