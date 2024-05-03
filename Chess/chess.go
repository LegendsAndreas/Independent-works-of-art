// To find our information: https://www.chess.com/article/view/how-to-set-up-a-chessboard#chess-queen
package main

import (
	"fmt"
	"log"
	"math"
	"sync"
)

// Chess pieces. "b" = black, "w" = white
var wPAWN = "♟"
var bPAWN = "♙"
var wROOK = "♜"
var bROOK = "♖"
var wKNIGHT = "♞"
var bKNIGHT = "♘"
var wBISHOP = "♝"
var bBISHOP = "♗"
var wQUEEN = "♛"
var bQUEEN = "♕"
var wKING = "♚"
var bKING = "♔"

// `square` represents a chess square on the chessboard. It has the following properties:
// - `letter`: the letter identifier of the square (e.g., "a1", "b2").
// - `color`: the color of the square (e.g., "white", "black").
// - `piece`: a pointer to the chess piece on the square. Set to `nil` if the square is empty.
type square struct {
	letter         string
	color          string
	gridCoordinate []int
	piece          *string
}

// Used for the goroutine, to check if either black or white King is still alive.
var wg sync.WaitGroup

// After careful consideration and thought, I have come to the realization that I should have used a map of some kind.
func main() {
	// Our chess board, as an array of squares.
	var chessBoard []square
	// Creates the chess board with the appropriate letters and colors.
	chessBoard = createBoard(chessBoard)
	// Set the chess pieces to be at their appropriate starting squares.
	initializePieces(chessBoard)

	// 'startingSquareLetter' represents the square with the chess piece that the user wants to move.
	// 'endingSquareLetter' represents the square that the user wants to move their piece to.
	var startingSquareLetter, endingSquareLetter string
	// These two variables look if the white and black queen is still alive, respectively.
	var wKingStatus = true
	var bKingtatus = true

	// The loop for our chess game.
	for {
		// Prints both the initiated board and updated board.
		printBoard(chessBoard)

		// Asks the user for which piece he/she wants to move and where he/she want to move that piece to.
		// This is done by the user typing 2 squares, the starting square and the ending square.
		fmt.Print("Enter move (eg. b2 b4)> ")
		_, err := fmt.Scan(&startingSquareLetter, &endingSquareLetter)
		if err != nil {
			log.Fatal("Input error:", err)
		}

		// In case the user wants to stop playing, he has to enter x, in which case the loop breaks
		if startingSquareLetter == "x" || endingSquareLetter == "x" {
			fmt.Println("You have chosen to end the game! :(")
			break
		}

		// In case a move is not valid, the program prints a message and starts the loop from the beginning.
		if !moveCheck(startingSquareLetter, endingSquareLetter, chessBoard) {
			fmt.Println("Invalid move. Please try again.")
			continue
		}

		// Moves the piece, assuming that the move is valid. This functions job is to ONLY move a piece.
		chessBoard = move(startingSquareLetter, endingSquareLetter, chessBoard)

		// It looks if either the white or black Queen is alive at the same time, by executing them as Goroutines.
		wg.Add(2)
		go isKingAlive(&wKingStatus, wKING, chessBoard)
		go isKingAlive(&bKingtatus, bKING, chessBoard)
		wg.Wait()
		if wKingStatus == true && bKingtatus == false { // In case, the white King is alive, while the black is dead.
			fmt.Println("The black King is dead! White has won!")
			printBoard(chessBoard) // Prints the final board
			break
		} else if wKingStatus == false && bKingtatus == true { // In case, the black King is alive, while the white is dead.
			printBoard(chessBoard) // Prints the final board
			fmt.Println("The white King is dead! Black has won!")
			break
		}
	}

}

func moveCheck(startSquare string, endSquare string, board []square) bool {
	// If both squares are equal to each other, we return false.
	if startSquare == endSquare {
		fmt.Println("The starting square is the same as the ending square!")
		return false
	}

	// Checks if the starting square exists.
	var sSquare = false
	var sPiece square // Assuming that the starting square exists, sPiece will be assigned the value of the board where the square exists.
	for i, sq := range board {
		if sq.letter == startSquare {
			sSquare = true
			sPiece = board[i] // Will be used later on.
		}
	}

	// Checks if the ending square exists.
	var eSquare = false
	var ePiece square
	for i, eq := range board {
		if eq.letter == endSquare {
			eSquare = true
			ePiece = board[i] // Will be used later on.
		}
	}

	// If either of the entered squares does not exist, we return false.
	if !sSquare || !eSquare {
		fmt.Println("One or both of the entered squares does not exist!")
		return false
	}

	// If the starting square (startSquare) does not have a piece in it, we return false.
	if sPiece.piece == nil {
		fmt.Println("The starting square does not have a piece!")
		return false
	}

	// If a piece is attacking a piece of the same color, we return false.
	// If you want to access a specific element, you have to put parenthesizing around the pointer part.

	if isEnemy(sPiece, ePiece) != true {
		fmt.Println("The piece is attacking a piece of the same color!")
		return false
	}
	/*
		if sPiece.piece != nil && ePiece.piece != nil {
			if string((*sPiece.piece)[0]) == string((*ePiece.piece)[0]) {
				fmt.Println(string((*sPiece.piece)[0]))
				fmt.Println(string((*ePiece.piece)[0]))
				fmt.Println("The piece is attacking a piece of the same color!")
				return false
			}
		}*/

	// If the user enters a pattern not supported by the appropriate piece, we return false.
	if !validMovement(sPiece, ePiece, board) {
		fmt.Println("Wrong movement")
		return false
	}

	// Assuming that we get no return false statements anywhere, we can return true.
	return true
}

// isEnemy returns true if the starting piece and the ending piece are enemies, and false otherwise.
// The startingPiece and endingPiece parameters are both of type square.
func isEnemy(startingPiece square, endingPiece square) bool {
	// If the starting piece is white.
	if startingPiece.piece != nil && endingPiece.piece != nil {
		if *startingPiece.piece == "♟" || *startingPiece.piece == "♜" ||
			*startingPiece.piece == "♞" || *startingPiece.piece == "♝" ||
			*startingPiece.piece == "♛" || *startingPiece.piece == "♚" {
			// And the ending piece is also white, we return false.
			if *endingPiece.piece == "♟" || *endingPiece.piece == "♜" ||
				*endingPiece.piece == "♞" || *endingPiece.piece == "♝" ||
				*endingPiece.piece == "♛" || *endingPiece.piece == "♚" {
				return false
			}
			// If the starting piece is black.
		} else if *startingPiece.piece == "♙" || *startingPiece.piece == "♖" ||
			*startingPiece.piece == "♘" || *startingPiece.piece == "♗" ||
			*startingPiece.piece == "♕" || *startingPiece.piece == "♔" {
			// And if the ending piece is also black, we return false.
			if *endingPiece.piece == "♙" || *endingPiece.piece == "♖" ||
				*endingPiece.piece == "♘" || *endingPiece.piece == "♗" ||
				*endingPiece.piece == "♕" || *endingPiece.piece == "♔" {
				return false
			}

		}

	}
	return true
}

func validMovement(startPiece square, endingPiece square, board []square) bool {
	// We check the movement of the piece, by comparing the starting position to the ending position.
	// startingGridX/Y is equal to the coordinates for where the piece started.
	// endingGridX/Y is equal to the coordinates for the piece moves to.
	var startingGridX = startPiece.gridCoordinate[1]
	var startingGridY = startPiece.gridCoordinate[0]
	var endingGridX = endingPiece.gridCoordinate[1]
	var endingGridY = endingPiece.gridCoordinate[0]

	if *startPiece.piece == "♟" { // Code for White Pawn
		// If a pawn moves diagonally and the square it moves to does not have a piece, it returns false.
		if endingPiece.piece == nil && (startingGridX > endingGridX || startingGridX < endingGridX) {
			fmt.Println("There is no piece for the Pawn to capture!")
			return false
		}

		if startingGridY == 2 && endingGridY == 4 && startingGridX == endingGridX && endingPiece.piece == nil { // The pawn moves 2 squares up from the start, assuming that the ending square is empty.
			if hasPiece(startingGridX, startingGridY+1, board) {
				fmt.Println("The path you chose for the White Pawn, has a piece in it's way!")
				return false
			}
			return true
		} else if startingGridY+1 == endingGridY && math.Abs(float64(startingGridX)-float64(endingGridX)) == 1 && endingPiece.piece != nil { // The Pawn moves diagonally, either up-left or up-right, assuming that the ending square is not empty.
			// This piece of code checks if the pawn is at the end of the board. In which case, it will become a Queen.
			if endingGridY == 8 {
				for i := range board {
					if board[i].letter == startPiece.letter {
						board[i].piece = &wQUEEN
						return true
					}
				}
			}
			return true
		} else if startingGridY+1 == endingGridY && startingGridX == endingGridX && endingPiece.piece == nil { // The pawn moves up once, assuming that the square ahead of it is empty.
			// This piece of code checks if the pawn is at the end of the board. In which case, it will become a Queen.
			if endingGridY == 8 {
				for i := range board {
					if board[i].letter == startPiece.letter {
						board[i].piece = &wQUEEN
						return true
					}
				}
			}
			return true
		}

	} else if *startPiece.piece == "♙" { // Code for Black Pawn
		// If a pawn moves diagonally and the square it moves to does not have a piece, it returns false.
		if endingPiece.piece == nil && (startingGridX > endingGridX || startingGridX < endingGridX) {
			fmt.Println("There is no piece for the Pawn to capture!")
			return false
		}

		if startingGridY == 7 && endingGridY == 5 && startingGridX == endingGridX && endingPiece.piece == nil { // The Pawn moves 2 squares down from the start, assuming that the ending square is empty.
			if hasPiece(startingGridX, startingGridY-1, board) { // We look at the square +1 ahead of the pawn and of there is a piece in the way, we truen false.
				fmt.Println("The path you chose for the White Pawn, has a piece in it's way!")
				return false
			}
			return true
		} else if startingGridY-1 == endingGridY && math.Abs(float64(startingGridX)-float64(endingGridX)) == 1 && endingPiece.piece != nil { // The Pawn moves diagonally, either down-left or down-right, assuming that the ending square is not empty.

			// This piece of code checks if the black pawn is at the end of the board. In which case, it will become a black Queen.
			if endingGridY == 1 {
				for i := range board {
					if board[i].letter == startPiece.letter {
						board[i].piece = &bQUEEN
						return true // The loop ends prematurely when the appropriate square is found, so we don't have to go through more than we need to.
					}
				}
			}

			return true
		} else if startingGridY-1 == endingGridY && startingGridX == endingGridX && endingPiece.piece == nil { // The Pawn moves down once, assuming that the square ahead of it is empty.

			// This piece of code checks if the black pawn is at the end of the board. In which case, it will become a black Queen.
			if endingGridY == 1 {
				for i := range board {
					if board[i].letter == startPiece.letter {
						board[i].piece = &bQUEEN
						return true // The loop ends prematurely when the appropriate square is found, so we don't have to go through more than we need to.
					}
				}
			}
			return true
		}

	} else if *startPiece.piece == "♜" || *startPiece.piece == "♖" { // Code for Rook
		if startingGridY < endingGridY && startingGridX == endingGridX { // The Rook moves up.
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridY--

				if startingGridY != endingGridY {
					if hasPiece(endingGridX, endingGridY, board) == true {
						fmt.Println("The path you chose for the Rook, has a piece in it's way!")
						return false
					}
				}
			}

		} else if startingGridY == endingGridY && startingGridX < endingGridX { // The Rook moves right
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridX--

				if startingGridX != endingGridX {
					if hasPiece(endingGridX, endingGridY, board) == true {
						fmt.Println("The path you chose for the Rook, has a piece in it's way!")
						return false
					}
				}
			}
		} else if startingGridY > endingGridY && startingGridX == endingGridX { // The Rook moves down.
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridY++

				if startingGridY != endingGridY {
					if hasPiece(endingGridX, endingGridY, board) == true {
						fmt.Println("The path you chose for the Rook, has a piece in it's way!")
						return false
					}
				}
			}
		} else if startingGridY == endingGridY && startingGridX > endingGridX { // The Rook moves left.
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridX++

				if startingGridX != endingGridX {
					if hasPiece(endingGridX, endingGridY, board) == true {
						fmt.Println("The path you chose for the Rook, has a piece in it's way!")
						return false
					}
				}
			}
		}

		// Code for Knight
	} else if *startPiece.piece == "♞" || *startPiece.piece == "♘" {
		// For the Knight, we can return true if the absolute difference between the starting and ending grid coordinates is 2 in one direction and 1 in the other direction. The math.Abs function only accepts floats, so we convert them.
		if (math.Abs(float64(startingGridX)-float64(endingGridX)) == 2 && math.Abs(float64(startingGridY)-float64(endingGridY)) == 1) ||
			(math.Abs(float64(startingGridX)-float64(endingGridX)) == 1 && math.Abs(float64(startingGridY)-float64(endingGridY)) == 2) {
			return true
		}
	} else if *startPiece.piece == "♝" || *startPiece.piece == "♗" { // Code for Bishop

		if startingGridY < endingGridY && startingGridX > endingGridX { // The Bishop moves up-left.
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridX++
				endingGridY--

				// If the square we are checking has another piece on it, we return false, since the Bishop cant jump over pieces.
				// We also have to make sure that we are not checking the square of the piece that we want to move, since that of course has a piece.
				if startingGridX != endingGridX && startingGridY != endingGridY {
					if hasPiece(endingGridX, endingGridY, board) {
						fmt.Println("The path you chose for the Bishop, has a piece in it's way!")
						return false
					}
				}
			}

		} else if startingGridY < endingGridY && startingGridX < endingGridX { // The Bishop moves up-right.
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridX--
				endingGridY--

				// If the square we are checking has another piece on it, we return false, since the Bishop cant jump over pieces.
				if startingGridX != endingGridX && startingGridY != endingGridY {
					if hasPiece(endingGridX, endingGridY, board) {
						fmt.Println("The path you chose for the Bishop, has a piece in it's way!")
						return false
					}
				}
			}

		} else if startingGridY > endingGridY && startingGridX < endingGridX { // The Bishop moves down-right.
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridX--
				endingGridY++

				// If the square we are checking has another piece on it, we return false, since the Bishop cant jump over pieces.
				if startingGridX != endingGridX && startingGridY != endingGridY {
					if hasPiece(endingGridX, endingGridY, board) {
						fmt.Println("The path you chose for the Bishop, has a piece in it's way!")
						return false
					}
				}
			}

		} else if startingGridY > endingGridY && startingGridX > endingGridX { // The Bishop moves down-left
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridX++
				endingGridY++

				// If the square we are checking has another piece on it, we return false, since the Bishop cant jump over pieces.
				if startingGridX != endingGridX && startingGridY != endingGridY {
					if hasPiece(endingGridX, endingGridY, board) {
						fmt.Println("The path you chose for the Bishop, has a piece in it's way!")
						return false
					}
				}
			}
		}

	} else if *startPiece.piece == "♛" || *startPiece.piece == "♕" { // Code for Queen
		// For the Queen, we can check if the movement is valid by combining the movement of the Rook and the Bishop.
		if startingGridY < endingGridY && startingGridX == endingGridX { // The Rook moves up.
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridY--

				if startingGridY != endingGridY {
					if hasPiece(endingGridX, endingGridY, board) == true {
						fmt.Println("The path you chose for the Queen, has a piece in it's way!")
						return false
					}
				}
			}

		} else if startingGridY == endingGridY && startingGridX < endingGridX { // The Rook moves right
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridX--

				if startingGridX != endingGridX {
					if hasPiece(endingGridX, endingGridY, board) == true {
						fmt.Println("The path you chose for the Queen, has a piece in it's way!")
						return false
					}
				}
			}

		} else if startingGridY > endingGridY && startingGridX == endingGridX { // The Rook moves down.
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridY++

				if startingGridY != endingGridY {
					if hasPiece(endingGridX, endingGridY, board) == true {
						fmt.Println("The path you chose for the Queen, has a piece in it's way!")
						return false
					}
				}
			}

		} else if startingGridY == endingGridY && startingGridX > endingGridX { // The Rook moves left.
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridX++

				if startingGridX != endingGridX {
					if hasPiece(endingGridX, endingGridY, board) == true {
						fmt.Println("The path you chose for the Queen, has a piece in it's way!")
						return false
					}
				}
			}
		} else if startingGridY < endingGridY && startingGridX > endingGridX { // The Queen moves up-left.
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridX++
				endingGridY--

				// If the square we are checking has another piece on it, we return false, since the Queen cant jump over pieces.
				if startingGridX != endingGridX && startingGridY != endingGridY {
					if hasPiece(endingGridX, endingGridY, board) {
						fmt.Println("The path you chose for the Queen, has a piece in it's way!")
						return false
					}
				}
			}

		} else if startingGridY < endingGridY && startingGridX < endingGridX { // The Queen moves up-right.
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridX--
				endingGridY--

				// If the square we are checking has another piece on it, we return false, since the Queen cant jump over pieces.
				if startingGridX != endingGridX && startingGridY != endingGridY {
					if hasPiece(endingGridX, endingGridY, board) {
						fmt.Println("The path you chose for the Queen, has a piece in it's way!")
						return false
					}
				}
			}

		} else if startingGridY > endingGridY && startingGridX < endingGridX { // The Queen moves down-right.
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridX--
				endingGridY++

				// If the square we are checking has another piece on it, we return false, since the Queen cant jump over pieces.
				if startingGridX != endingGridX && startingGridY != endingGridY {
					if hasPiece(endingGridX, endingGridY, board) {
						fmt.Println("The path you chose for the Queen, has a piece in it's way!")
						return false
					}
				}
			}

		} else if startingGridY > endingGridY && startingGridX > endingGridX { // The Queen moves down-left
			for startingGridX > 0 && startingGridX < 9 && startingGridY > 0 && startingGridY < 9 {
				if startingGridX == endingGridX && startingGridY == endingGridY {
					return true
				}
				endingGridX++
				endingGridY++

				// If the square we are checking has another piece on it, we return false, since the Queen cant jump over pieces.
				if startingGridX != endingGridX && startingGridY != endingGridY {
					if hasPiece(endingGridX, endingGridY, board) {
						fmt.Println("The path you chose for the Queen, has a piece in it's way!")
						return false
					}
				}
			}
		}

	} else if *startPiece.piece == "♚" || *startPiece.piece == "♔" { // Code for King
		// For the King, we can return true if the absolute difference between the starting and ending grid coordinates is less than or equal to 1 in both the x and y directions.
		// We also have to turn the startingGrid values into floats, because the math.Absolute function only accepts floats.
		if math.Abs(float64(startingGridX)-float64(endingGridX)) <= 1 && math.Abs(float64(startingGridY)-float64(endingGridY)) <= 1 {
			return true
		}
	}

	return false
}

// hasPiece checks if a specific square on the board contains a chess piece.
// It takes in the x and y coordinates of the square and the chess board as inputs.
// It returns true if the square has a piece and false otherwise.
func hasPiece(xCoordinate int, yCoordinate int, board []square) bool {
	for i := range board {
		if board[i].gridCoordinate[0] == yCoordinate && board[i].gridCoordinate[1] == xCoordinate {
			if board[i].piece == nil {
				// fmt.Printf("grid x: %d and grid y: %d does NOT have a piece.\n", xCoordinate, yCoordinate) // Tool for bug-finding
				return false
			} else {
				fmt.Println(*board[i].piece)
				// fmt.Printf("grid x: %d and grid y: %d DOES have a piece.\n", xCoordinate, yCoordinate) // Tool for bug-finding
				return true
			}
		}
	}

	// SHOULD, be impossible, but the program gets mad if this return statement does not exist.
	fmt.Println("That coordinate combination does not even exist!")
	return true
}

// initializePieces sets the chess pieces to their appropriate starting squares on the chess board.
func initializePieces(board []square) {
	// Sets black and white pawns to their appropriate squares.
	for i := 8; i < 16; i++ {
		board[i].piece = &bPAWN
	}
	for i := 48; i < 56; i++ {
		board[i].piece = &wPAWN
	}

	// Sets the black rooks, knights, bishops, queen and king to their appropriate squares
	board[0].piece = &bROOK
	board[7].piece = &bROOK
	board[1].piece = &bKNIGHT
	board[6].piece = &bKNIGHT
	board[2].piece = &bBISHOP
	board[5].piece = &bBISHOP
	board[3].piece = &bQUEEN
	board[4].piece = &bKING

	// Sets the white rooks, knights, bishops, queen and king to their appropriate squares
	board[56].piece = &wROOK
	board[63].piece = &wROOK
	board[57].piece = &wKNIGHT
	board[62].piece = &wKNIGHT
	board[58].piece = &wBISHOP
	board[61].piece = &wBISHOP
	board[59].piece = &wQUEEN
	board[60].piece = &wKING
}

// printBoard prints the chess board, including the pieces if a square has a chess piece.
// The board parameter is an array of squares representing the chess board.
func printBoard(board []square) {
	// There HAS to be a comma when we print the new line and a space when we print, otherwise the squares that we print will not be formatted correctly.
	numRow := 7
	fmt.Print("8") // We simply print out 8 at the beginning to make sure the code works later on.
	for i := 0; i < len(board); i++ {
		// Once a row has been created, which is 8 in length, we print out a comma with a newline to format it correctly, and then print out what which number row it is, starting from the top, down to 1.
		if i%8 == 0 && i > 0 {
			fmt.Println(",")
			fmt.Print(numRow)
			numRow--
		}
		// In case a square has a chess piece, it will then print the piece, rather than the square.
		if board[i].piece == nil {
			fmt.Printf("%s ", board[i].color)
		} else {
			fmt.Printf("%s ", *board[i].piece)
		}
	}
	fmt.Println(",")
	fmt.Println(" a b c d e f g h") // The letter rows gets printed out.
}

// createBoard generates a chess board with the appropriate letters and colors.
// It takes an empty array of squares as input and returns a populated chess board.
func createBoard(board []square) []square {
	// If you look at a chessboard, we start in the top left corner, to the right, then downwards.
	// This is done, so that it gets print our as: Bottom left corner, to the right, then upwards, later on in the printBoard function.
	// Unicode: "\u2B1C" creates a white square, "\u2B1B" creates a black square.

	// Row 8
	board = append(board, square{letter: "a8", color: "⬜"})
	board = append(board, square{letter: "b8", color: "⬛"})
	board = append(board, square{letter: "c8", color: "⬜"})
	board = append(board, square{letter: "d8", color: "⬛"})
	board = append(board, square{letter: "e8", color: "⬜"})
	board = append(board, square{letter: "f8", color: "⬛"})
	board = append(board, square{letter: "g8", color: "⬜"})
	board = append(board, square{letter: "h8", color: "⬛"})

	// Row 7
	board = append(board, square{letter: "a7", color: "⬛"})
	board = append(board, square{letter: "b7", color: "⬜"})
	board = append(board, square{letter: "c7", color: "⬛"})
	board = append(board, square{letter: "d7", color: "⬜"})
	board = append(board, square{letter: "e7", color: "⬛"})
	board = append(board, square{letter: "f7", color: "⬜"})
	board = append(board, square{letter: "g7", color: "⬛"})
	board = append(board, square{letter: "h7", color: "⬜"})

	// Row 6
	board = append(board, square{letter: "a6", color: "⬜"})
	board = append(board, square{letter: "b6", color: "⬛"})
	board = append(board, square{letter: "c6", color: "⬜"})
	board = append(board, square{letter: "d6", color: "⬛"})
	board = append(board, square{letter: "e6", color: "⬜"})
	board = append(board, square{letter: "f6", color: "⬛"})
	board = append(board, square{letter: "g6", color: "⬜"})
	board = append(board, square{letter: "h6", color: "⬛"})

	// Row 5
	board = append(board, square{letter: "a5", color: "⬛"})
	board = append(board, square{letter: "b5", color: "⬜"})
	board = append(board, square{letter: "c5", color: "⬛"})
	board = append(board, square{letter: "d5", color: "⬜"})
	board = append(board, square{letter: "e5", color: "⬛"})
	board = append(board, square{letter: "f5", color: "⬜"})
	board = append(board, square{letter: "g5", color: "⬛"})
	board = append(board, square{letter: "h5", color: "⬜"})

	// Row 4
	board = append(board, square{letter: "a4", color: "⬜"})
	board = append(board, square{letter: "b4", color: "⬛"})
	board = append(board, square{letter: "c4", color: "⬜"})
	board = append(board, square{letter: "d4", color: "⬛"})
	board = append(board, square{letter: "e4", color: "⬜"})
	board = append(board, square{letter: "f4", color: "⬛"})
	board = append(board, square{letter: "g4", color: "⬜"})
	board = append(board, square{letter: "h4", color: "⬛"})

	// Row 3
	board = append(board, square{letter: "a3", color: "⬛"})
	board = append(board, square{letter: "b3", color: "⬜"})
	board = append(board, square{letter: "c3", color: "⬛"})
	board = append(board, square{letter: "d3", color: "⬜"})
	board = append(board, square{letter: "e3", color: "⬛"})
	board = append(board, square{letter: "f3", color: "⬜"})
	board = append(board, square{letter: "g3", color: "⬛"})
	board = append(board, square{letter: "h3", color: "⬜"})

	// Row 2
	board = append(board, square{letter: "a2", color: "⬜"})
	board = append(board, square{letter: "b2", color: "⬛"})
	board = append(board, square{letter: "c2", color: "⬜"})
	board = append(board, square{letter: "d2", color: "⬛"})
	board = append(board, square{letter: "e2", color: "⬜"})
	board = append(board, square{letter: "f2", color: "⬛"})
	board = append(board, square{letter: "g2", color: "⬜"})
	board = append(board, square{letter: "h2", color: "⬛"})

	// Row 1
	board = append(board, square{letter: "a1", color: "⬛"})
	board = append(board, square{letter: "b1", color: "⬜"})
	board = append(board, square{letter: "c1", color: "⬛"})
	board = append(board, square{letter: "d1", color: "⬜"})
	board = append(board, square{letter: "e1", color: "⬛"})
	board = append(board, square{letter: "f1", color: "⬜"})
	board = append(board, square{letter: "g1", color: "⬛"})
	board = append(board, square{letter: "h1", color: "⬜"})

	// Our grid system goes from the bottom left to right and then upwards. We start from 1,1 to 1,2... 1,8 to 2,1 to 2,2...
	row := 1
	colom := 8
	for i := len(board) - 1; i >= 0; i-- {
		board[i].gridCoordinate = append(board[i].gridCoordinate, row, colom)
		colom--

		if i%8 == 0 {
			row++
			colom = 8
		}
	}

	// Prints the grid coordinates of the board.
	for i := range board {
		if i%8 == 0 && i > 0 {
			fmt.Println()
		}
		fmt.Print(board[i].gridCoordinate)
	}

	fmt.Println()

	return board
}

// isKingAlive checks if the given queen is still alive on the chess board.
// It updates the queenStatus variable to reflect the status of the queen.
// If the queen is found on the board, queenStatus is set to true.
// The function uses a pointer to the queenStatus variable to update its value.
// The function iterates through the board squares to find the queen in the board array.
// It stops searching when the queen is found or when the end of the board is reached.
// After it finishes searching, the function calls wg.Done() to decrement the wait group counter.
func isKingAlive(kingStatus *bool, king string, board []square) {
	// Generally risky to dereference without checking of the pointer is nil.
	if kingStatus != nil {
		*kingStatus = false
	}

	// If the Queen is found, the 'kingStatus' will be true.
	for i := range board {
		if board[i].piece != nil && *board[i].piece == king {
			*kingStatus = true
			break
		}
	}
	// 1 gets subtracted from the wg waitGroup.
	wg.Done()
}

// move function moves a chess piece from a starting square to an ending square on the given chess board.
// It takes the starting square and ending square as string arguments, and the current board as a slice of "square" structs.
// It updates the board by moving the piece from the starting square to the ending square.
// If the piece at the starting square is not found, or the ending square is invalid, the board remains unchanged.
// The updated board is returned.
func move(startSquare string, endSquare string, board []square) []square {
	var startIdx int
	var endIdx int
	// Loops though the entire array, and when it finds the index at which the startSquare or the endSquare presides, it sets startIdx and endIdx to be equal to that array index.
	// Would have been a lot easier if I had just used a map.
	for i, sq := range board {
		if sq.letter == startSquare {
			startIdx = i
		}
		if sq.letter == endSquare {
			endIdx = i
		}
	}

	// Replaces the endSquare.piece with the startSquare.piece and sets the startSquare.piece to be empty, since it no longer has a piece.
	board[endIdx].piece = board[startIdx].piece
	board[startIdx].piece = nil

	return board
}
