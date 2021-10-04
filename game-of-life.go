package main

import (
	"fmt"
	"strings"
	"time"
	"math/rand"
)

const (
	Dead_space = "-"
	Alive_space = "X"
	Rows = 40
	Columns = 40
)

func print_board(board [][]string) {
	for row := 0; row < Rows; row++ {
		fmt.Println(strings.Join(board[row], " "))
	}
}

func main() {
	fmt.Println("Game of Life!")

	choices := []string{Dead_space, Alive_space}

	board := make([][]string, Rows)

	for row := 0; row < Rows; row++ {
		board[row] = make([]string, Columns)
		for column := 0; column < Columns; column++ {
			random_choice := choices[rand.Intn(len(choices))]
			board[row][column] = random_choice
		}
	}

	for turn := 0; ; turn++ {
		fmt.Printf("Turn %v\n", turn)
		print_board(board)
		time.Sleep(1 * time.Second)
	}
}
