package main

import "fmt"
import "strings"

func print_board(board [][]string, rows int, columns int) {
	for row := 0; row < rows; row++ {
		fmt.Println(strings.Join(board[row], " "))
	}
}

func main() {
	fmt.Println("Game of Life!")

	const rows = 50
	const columns = 50
	const dead_space = "-"
	const alive_space = "X"

	board := make([][]string, rows)

	for row := 0; row < rows; row++ {
		board[row] = make([]string, columns)
		for column := 0; column < columns; column++ {
			board[row][column] = dead_space
		}
	}
	print_board(board, rows, columns)
}
