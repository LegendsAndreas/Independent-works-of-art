// To find our information: https://www.chess.com/article/view/how-to-set-up-a-chessboard#chess-queen
package main

import (
	"fmt"
	"log"
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
	letter string
	color  string
	piece  *string
}

// After careful consideration and thought, i have come to the realization that i should have used map of some kind.
// I have to loop though the
func main() {
	// Our chess board, as an array of squares.
	var chessBoard []square
	// Creates the chess board with the appropriate letters and colors.
	chessBoard = createBoard(chessBoard)
	// Set the chess pieces to be at their appropriate starting squares.
	initializePieces(chessBoard)

	// 'startingSquare' represents the square with the chess piece that the user wants to move.
	// 'endingSquare' represents the square that the user wants to move their piece to.
	var startingSquare, endingSquare string
	// The loop for our chess game.
	for {
		printBoard(chessBoard)

		fmt.Print("Enter move (eg. b2 b4)> ")
		_, err := fmt.Scan(&startingSquare, &endingSquare)
		if err != nil {
			log.Fatal("Input error:", err)
		}

		// In case the user wants to stop playing, he has to enter x, in which case the loop breaks
		if startingSquare == "x" || endingSquare == "x" {
			break
		}

		// In case a move is not valid, the program prints a message and starts the loop from the beginning.
		if !moveCheck(startingSquare, endingSquare, chessBoard) {
			fmt.Println("Invalid move. Please try again.")
			continue
		}

		chessBoard = move(startingSquare, endingSquare, chessBoard)
	}

}

// move function moves a chess piece from a starting square to an ending square on the given chess board.
// It takes the starting square and ending square as string arguments, and the current board as a slice of "square" structs.
// It updates the board by moving the piece from the starting square to the ending square.
// If the piece at the starting square is not found, or the ending square is invalid, the board remains unchanged.
// The updated board is returned.
func move(startSquare string, endSquare string, board []square) []square {
	var startIdx int
	var endIdx int
	// Loops though the entire array, and when it finds the index at which the startSquare or the endSquare presides,
	// it sets startIdx and endIdx to be equal to that array index.
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

func moveCheck(startSquare string, endSquare string, board []square) bool {
	// If the entered squares exists, the bool is true.
	var sSquare bool
	var eSquare bool
	for _, sq := range board { // Checks if the starting square exists.
		if sq.letter == startSquare {
			sSquare = true
		}
	}

	for _, eq := range board { // Checks if the ending square exists.
		if eq.letter == endSquare {
			eSquare = true
		}
	}

	// If both entered squares exists, the boolshit will be true.
	// If that is not the case, we prematurely return the boolshit as false.
	var boolshit = false
	if sSquare && eSquare {
		boolshit = true
	} else {
		return boolshit
	}

	if startSquare == endSquare { // If both square are equal to each other, we set boolshit to be false.
		boolshit = false
	}

	return boolshit
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
	for i := 0; i < len(board); i++ {
		if i%8 == 0 && i > 0 {
			fmt.Println(",")
		}

		// In case a square has a chess piece, it will then print the piece, rather than the square.
		if board[i].piece == nil {
			fmt.Printf("%s ", board[i].color)
		} else {
			fmt.Printf("%s ", *board[i].piece)
		}
	}
	println(",")
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

	return board
}
