// To find our information: https://www.chess.com/article/view/how-to-set-up-a-chessboard#chess-queen
package main

import "fmt"

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
	number int
	color  string
	piece  *string
}

type chessPiece struct {
	name string
}

func main() {

	var chessBoard []square
	chessBoard = createBoard(chessBoard)
	//printBoard(chessBoard)
	initializePawns(chessBoard)
	printBoard(chessBoard)

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
	board = append(board, square{letter: "a", number: 8, color: "⬜"})
	board = append(board, square{letter: "b", number: 8, color: "⬛"})
	board = append(board, square{letter: "c", number: 8, color: "⬜"})
	board = append(board, square{letter: "d", number: 8, color: "⬛"})
	board = append(board, square{letter: "e", number: 8, color: "⬜"})
	board = append(board, square{letter: "f", number: 8, color: "⬛"})
	board = append(board, square{letter: "g", number: 8, color: "⬜"})
	board = append(board, square{letter: "h", number: 8, color: "⬛"})

	// Row 7
	board = append(board, square{letter: "a", number: 7, color: "⬛"})
	board = append(board, square{letter: "b", number: 7, color: "⬜"})
	board = append(board, square{letter: "c", number: 7, color: "⬛"})
	board = append(board, square{letter: "d", number: 7, color: "⬜"})
	board = append(board, square{letter: "e", number: 7, color: "⬛"})
	board = append(board, square{letter: "f", number: 7, color: "⬜"})
	board = append(board, square{letter: "g", number: 7, color: "⬛"})
	board = append(board, square{letter: "h", number: 7, color: "⬜"})

	// Row 6
	board = append(board, square{letter: "a", number: 6, color: "⬜"})
	board = append(board, square{letter: "b", number: 6, color: "⬛"})
	board = append(board, square{letter: "c", number: 6, color: "⬜"})
	board = append(board, square{letter: "d", number: 6, color: "⬛"})
	board = append(board, square{letter: "e", number: 6, color: "⬜"})
	board = append(board, square{letter: "f", number: 6, color: "⬛"})
	board = append(board, square{letter: "g", number: 6, color: "⬜"})
	board = append(board, square{letter: "h", number: 6, color: "⬛"})

	// Row 5
	board = append(board, square{letter: "a", number: 5, color: "⬛"})
	board = append(board, square{letter: "b", number: 5, color: "⬜"})
	board = append(board, square{letter: "c", number: 5, color: "⬛"})
	board = append(board, square{letter: "d", number: 5, color: "⬜"})
	board = append(board, square{letter: "e", number: 5, color: "⬛"})
	board = append(board, square{letter: "f", number: 5, color: "⬜"})
	board = append(board, square{letter: "g", number: 5, color: "⬛"})
	board = append(board, square{letter: "h", number: 5, color: "⬜"})

	// Row 4
	board = append(board, square{letter: "a", number: 4, color: "⬜"})
	board = append(board, square{letter: "b", number: 4, color: "⬛"})
	board = append(board, square{letter: "c", number: 4, color: "⬜"})
	board = append(board, square{letter: "d", number: 4, color: "⬛"})
	board = append(board, square{letter: "e", number: 4, color: "⬜"})
	board = append(board, square{letter: "f", number: 4, color: "⬛"})
	board = append(board, square{letter: "g", number: 4, color: "⬜"})
	board = append(board, square{letter: "h", number: 4, color: "⬛"})

	// Row 3
	board = append(board, square{letter: "a", number: 3, color: "⬛"})
	board = append(board, square{letter: "b", number: 3, color: "⬜"})
	board = append(board, square{letter: "c", number: 3, color: "⬛"})
	board = append(board, square{letter: "d", number: 3, color: "⬜"})
	board = append(board, square{letter: "e", number: 3, color: "⬛"})
	board = append(board, square{letter: "f", number: 3, color: "⬜"})
	board = append(board, square{letter: "g", number: 3, color: "⬛"})
	board = append(board, square{letter: "h", number: 3, color: "⬜"})

	// Row 2
	board = append(board, square{letter: "a", number: 2, color: "⬜"})
	board = append(board, square{letter: "b", number: 2, color: "⬛"})
	board = append(board, square{letter: "c", number: 2, color: "⬜"})
	board = append(board, square{letter: "d", number: 2, color: "⬛"})
	board = append(board, square{letter: "e", number: 2, color: "⬜"})
	board = append(board, square{letter: "f", number: 2, color: "⬛"})
	board = append(board, square{letter: "g", number: 2, color: "⬜"})
	board = append(board, square{letter: "h", number: 2, color: "⬛"})

	// Row 1
	board = append(board, square{letter: "a", number: 1, color: "⬛"})
	board = append(board, square{letter: "b", number: 1, color: "⬜"})
	board = append(board, square{letter: "c", number: 1, color: "⬛"})
	board = append(board, square{letter: "d", number: 1, color: "⬜"})
	board = append(board, square{letter: "e", number: 1, color: "⬛"})
	board = append(board, square{letter: "f", number: 1, color: "⬜"})
	board = append(board, square{letter: "g", number: 1, color: "⬛"})
	board = append(board, square{letter: "h", number: 1, color: "⬜"})

	return board
}
