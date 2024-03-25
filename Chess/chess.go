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

type square struct {
	letter string
	color  string
	piece  *string
}

func main() {
	var chessBoard []square
	chessBoard = createBoard(chessBoard)
	initializePawns(chessBoard)

	var startingSquare, endingSquare string

	for{
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
		if !moveCheck(startingSquare, endingSquare, chessBoard){
			fmt.Println("Invalid move. Please try again.")
			continue
		}

		chessBoard = move(startingSquare, endingSquare, chessBoard)
	}

}

func move(startSquare string, endSquare string, board []square) []square {
	if

	return board
}

func moveCheck(startSquare string, endSquare string, board []square) bool {
	var boolshit = true

	if startSquare != "a1" || endSquare != "a1" {
		boolshit = false
	} else if startSquare == endSquare{
		boolshit = false
	}



	return boolshit
}

func initializePawns(board []square) {
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
}

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
