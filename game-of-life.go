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

func calculate_state (board [][]string, row int, column int) string {
	neighbours := [8][2]int{
		{ row - 1, column - 1 },
		{ row - 1, column },
		{ row - 1, column + 1},
		{ row, column - 1},
		{ row, column + 1},
		{ row + 1, column - 1 },
		{ row + 1, column },
		{ row + 1, column + 1},
	}
	fmt.Println(neighbours)
	return Dead_space
}

func main() {
	fmt.Println("Game of Life!")
	rand.Seed(time.Now().Unix())

	choices := []string{Dead_space, Alive_space}

	board := make([][]string, Rows)

	for row := 0; row < Rows; row++ {
		board[row] = make([]string, Columns)
		for column := 0; column < Columns; column++ {
			random_index := rand.Intn(len(choices))
			random_choice := choices[random_index]
			board[row][column] = random_choice
		}
	}

	for turn := 0; ; turn++ {
		fmt.Printf("Turn %v\n", turn)
		print_board(board)
		time.Sleep(1 * time.Second)
		for row := 0; row < Rows; row++ {
			for column := 0; column < Columns; column++ {
				new_state := calculate_state(board, row, column)
				board[row][column] = new_state
			}
		}
	}
}
