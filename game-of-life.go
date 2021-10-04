package main

import "fmt"
import "strings"
import "time"

func print_board(board [][]string, rows int, columns int) {
	for row := 0; row < rows; row++ {
		fmt.Println(strings.Join(board[row], " "))
	}
}

func main() {
	fmt.Println("Game of Life!")

	const rows = 40
	const columns = 40
	const dead_space = "-"
	const alive_space = "X"

	board := make([][]string, rows)

	for row := 0; row < rows; row++ {
		board[row] = make([]string, columns)
		for column := 0; column < columns; column++ {
			board[row][column] = dead_space
		}
	}

	for turn := 0; ; turn++ {
		fmt.Printf("Turn %v\n", turn)
		print_board(board, rows, columns)
		time.Sleep(1 * time.Second)
	}
}
