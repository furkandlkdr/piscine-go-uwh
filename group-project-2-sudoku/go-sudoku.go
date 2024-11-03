package main

import (
	"fmt"
	"os"
)

const size = 9

// Change the arguments to the desired integer sudoku board
func parseBoard(args []string) [][]int {
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
		for j, char := range args[i] {
			if char == '.' {
				board[i][j] = 0
			} else {
				board[i][j] = int(char - '0')
			}
		}
	}
	return board
}

// Print the board
func printBoard(board [][]int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", board[i][j])
			}
		}
		fmt.Println()
	}
}

func isValid(board [][]int, row, col, num int) bool {
	for i := 0; i < size; i++ {
		if board[row][i] == num || board[i][col] == num ||
			board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}

func solveSudoku(board [][]int) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(board, i, j, num) {
						board[i][j] = num
						if solveSudoku(board) {
							return true
						}
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Usage: go run . <row1> <row2> ... <row9>")
		return
	}

	board := parseBoard(os.Args[1:])

	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("Çözüm bulunamadı.")
	}
}
