package main

import (
	"fmt"
	"strings"
	"time"
	"math/rand"
)

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
	choices := []string{dead_space, alive_space}

	board := make([][]string, rows)

	for row := 0; row < rows; row++ {
		board[row] = make([]string, columns)
		for column := 0; column < columns; column++ {
			random_choice := choices[rand.Intn(len(choices))]
			board[row][column] = random_choice
		}
	}

	for turn := 0; ; turn++ {
		fmt.Printf("Turn %v\n", turn)
		print_board(board, rows, columns)
		time.Sleep(1 * time.Second)
	}
}
