package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	Dead_space  = "-"
	Alive_space = "X"
	Rows        = 10
	Columns     = 10
)

func print_board(board [][]string) {
	for row := 0; row < Rows; row++ {
		fmt.Println(strings.Join(board[row], " "))
	}
}

func calculate_state(board [][]string, row int, column int) string {
	neighbours := [8][2]int{
		{row - 1, column - 1},
		{row - 1, column},
		{row - 1, column + 1},
		{row, column - 1},
		{row, column + 1},
		{row + 1, column - 1},
		{row + 1, column},
		{row + 1, column + 1},
	}
	var valid_neighbours [][2]int
	for _, neighbour := range neighbours {
		if -1 < neighbour[0] && neighbour[0] < Rows && -1 < neighbour[1] && neighbour[1] < Columns {
			valid_neighbours = append(valid_neighbours, neighbour)
		}
	}

	alive_neighbour_counter := 0
	for _, neighbour := range valid_neighbours {
		neighbour_state := board[neighbour[0]][neighbour[1]]
		if neighbour_state == Alive_space {
			alive_neighbour_counter++
		}
	}

	current_cell := board[row][column]

	if alive_neighbour_counter == 3 {
		return Alive_space
	} else if current_cell == Alive_space && alive_neighbour_counter == 2 {
		return Alive_space
	} else {
		return Dead_space
	}
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
