package main

import (
	"errors"
	"fmt"
	logger "github.com/LegendsAndreas/logger"
	"io"
	"math"
	"strconv"
)

const pyramidRows = 4
const path = "C:\\Users\\andre\\Desktop\\Pog grammering\\Golang\\addition_pyramid"

type Brick struct {
	openingBracket rune
	number         int
	closingBracket rune
}

func main() {
	var pyramid = makePyramid(pyramidRows)
	printPyramid(pyramid)
	simplePyramidGame(pyramid)

}

// makePyramid constructs a pyramid structure of Bricks with the specified number of rows.
func makePyramid(rows int) [][]Brick {
	pyramid := make([][]Brick, 0)
	// Rows added for each iteration:  +1   +2   +3   +4    +5...
	// Total rows: 				     0 -> 1 -> 3 -> 6 -> 10 -> 15...
	for i := 0; i < rows; i++ {
		var bricks []Brick
		for j := 0; j < i+1; j++ {
			var tempBrick Brick
			tempBrick.openingBracket = '['
			tempBrick.closingBracket = ']'
			bricks = append(bricks, tempBrick)
		}
		pyramid = append(pyramid, bricks)

	}

	return pyramid
}

// printPyramid prints a pyramid of Bricks, formatting each row with appropriate spaces for alignment.
func printPyramid(pyramid [][]Brick) {
	//fmt.Println(len(pyramid))
	var spaces = ""
	for i := 0; i < len(pyramid); i++ {
		//fmt.Println("i brick =", pyramid[i])
		// Spacing for formatting.
		spaces = getSpaces(len(pyramid)*2 - i*2)
		fmt.Print(spaces)
		for j := 0; j < len(pyramid[i]); j++ {
			// fmt.Println("j brick =", pyramid[i][j])
			fmt.Printf("%c%2d%c", pyramid[i][j].openingBracket, pyramid[i][j].number, pyramid[i][j].closingBracket)
		}
		// Newline for formatting.
		fmt.Println()
	}
}

// getSpaces returns a string of spaces with the length specified by the input parameter 'num'.
// The function concatenates individual space characters to form the final string.
func getSpaces(num int) string {
	var spaces = ""
	for i := 0; i < num; i++ {
		spaces = spaces + " "
	}
	return spaces
}

// simplePyramidGame orchestrates the steps to play a pyramid game where user inputs values into a pyramid structure.
func simplePyramidGame(pyramid [][]Brick) {
	var minimunValue = getMinimumValue(pyramid)
	var topValue = getTopValue(minimunValue)
	setTopValue(topValue, pyramid)
	beginSimpleGame(pyramid)
}

// getMinimunValue calculates the minimum valid number for the top brick in a pyramid based on its depth (number of levels).
// The minimum value is computed as 2 raised to the power of (depth - 1).
func getMinimumValue(pyramid [][]Brick) int {
	var minimumValueFloat64 = math.Pow(2, float64(len(pyramid)-1))
	var minimumValueInt = int(minimumValueFloat64)
	return minimumValueInt
}

// getTopValue prompts the user to input a top value for
// the pyramid that is greater than or equal to the provided minimum value.
func getTopValue(minimunValue int) int {
	var topValue = 0
	for {
		fmt.Print("Enter top value>")
		_, err := fmt.Scan(&topValue)
		if err != nil {
			return 0
		}

		if topValue >= minimunValue {
			fmt.Println("The top value is:", topValue)
			break
		}

		fmt.Printf("Top value cant be lower than the minimum value (%d).\n", minimunValue)
	}

	return topValue
}

// setTopValue sets the number of the top brick (0, 0) in the pyramid to topValue.
func setTopValue(topValue int, pyramid [][]Brick) {
	pyramid[0][0].number = topValue
}

// beginSimpleGame initiates a simple game where users input values into a pyramid structure.
//
// Parameters:
//   - pyramid ([][]Brick): A 2D slice representing the pyramid's grid of bricks, where each brick holds a numerical value.
//
// The function continually prompts the user to input a row, brick, and value to populate a specific position
// in the pyramid. It validates the inputs and checks if the entered values meet the game criteria. The game
// loop continues until the user wins by filling the pyramid correctly.
func beginSimpleGame(pyramid [][]Brick) {
	var inputRow int
	var inputBrick int
	var inputValue int
	for {
		// We print the pyramid at the start of the loop, so that the user can always see their change.
		printPyramid(pyramid)

		// Asks for which row to add to.
		fmt.Print("Enter row>")
		_, err := fmt.Scan(&inputRow)
		if err != nil {
			handleInputError(err)
			continue
		}

		// Asks for which brick to add to.
		fmt.Print("Enter brick>")
		_, err = fmt.Scan(&inputBrick)
		if err != nil {
			handleInputError(err)
			continue
		}

		// Checks that the entered row and/or brick is valid.
		if validety := validateRowNBrick(pyramid, inputRow, inputBrick); !validety {
			fmt.Println("Invalid row or brick.")
			continue
		}

		// Asks for what number to add.
		fmt.Print("Enter number>")
		_, err = fmt.Scan(&inputValue)
		if err != nil {
			handleInputError(err)
			continue
		}

		// Checks that the entered number does not lose the player the game.

		// Adds the entered number to the entered row and brick.
		// The user will not take into account that we start at 0 and not one, so we minus one.
		pyramid[(inputRow - 1)][(inputBrick - 1)].number = inputValue

		// If every block adds up, we return true and the player has one.
		if validateValues(pyramid) {
			fmt.Println("You won!")
			break
		} else {
			fmt.Println("Not there yet...")
		}
	}
}

func handleInputError(err error) {
	var numErr *strconv.NumError
	if err == io.EOF {
		fmt.Println("Input was terminated unexpectedly:", err)
	} else if errors.As(err, &numErr) && errors.Is(numErr.Err, strconv.ErrSyntax) {
		fmt.Println("Invalid input:", err)
	} else {
		fmt.Println("Unknown error:", err)
	}

	// Finally just logs the error.
	logger.Error(err, "log", path)
}

func validateRowNBrick(pyramid [][]Brick, row int, brick int) bool {
	//fmt.Println(len(pyramid))
	// Checks that the entered row actually exist.
	if row < 0 || row > len(pyramid) {
		fmt.Println("Invalid row.")
		return false
	}

	//fmt.Println(len(pyramid[row-1]))
	// Checks that the entered brick actually exist.
	if brick < 0 || brick > len(pyramid[row-1]) {
		fmt.Println("Invalid brick.")
		return false
	}

	return true
}

// validateValues checks whether each brick's number is equal to the sum of the two bricks directly below it in the pyramid.
//
// Parameters:
//
//	pyramid ([][]Brick): The pyramid that you want to check.
func validateValues(pyramid [][]Brick) bool {
	// For the length of the entire pyramid.
	for i := 0; i < len(pyramid)-1; i++ {
		// For the length of the entire row of the pyramid.
		for j := 0; j < len(pyramid[i]); j++ {
			if pyramid[i][j].number != pyramid[i+1][j].number+pyramid[i+1][j+1].number {
				return false
			}
		}
	}

	return true
}
