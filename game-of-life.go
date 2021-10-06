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
	Rows        = 20
	Columns     = 20
)

func print_board(board [][]string) {
	colorReset := "\033[0m"
	for row := 0; row < Rows; row++ {
		fmt.Println(string(colorReset), strings.Join(board[row], " "))
	}
}

func print_board_int(board [][]int) {
	// for row := 0; row < Rows; row++ {
	// 	fmt.Println(strings.Join(board[row], " "))
	// }
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	for row := 0; row < Rows; row++ {
		for column := 0; column < Columns; column++ {
			if board[row][column] == 3 {
				fmt.Print(string(colorRed), board[row][column])
			} else {
				fmt.Print(string(colorGreen), board[row][column])
			}
		}
		fmt.Println()
	}
}

// func print_board_int_new(board [][]int) {
// 	for row := 0; row < Rows; row++ {
// 		new_row := []string
// 		for item
// 		fmt.Println(strings.Join(board[row], " "))
// 	}
// }

func count_neighbors(board [][]string) [][]int {
	counts := make([][]int, Rows)
	for row := 0; row < Rows; row++ {
		counts[row] = make([]int, Columns)
	}

	for row := range board {
		for column := range board[row] {
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

			counts[row][column] = alive_neighbour_counter
		}
	}

	return counts
}

func next_cell(row, col int, board [][]string, neighbor_counts [][]int) string {
	current_cell := board[row][col]

	if neighbor_counts[row][col] == 3 {
		return Alive_space
	} else if current_cell == Alive_space && neighbor_counts[row][col] == 2 {
		return Alive_space
	} else {
		return Dead_space
	}
}

func clear_screen() {
	fmt.Print("\033[H\033[2J")
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

		time.Sleep(100 * time.Millisecond)
		clear_screen()

		counts := count_neighbors(board)

		print_board_int(counts)

		for row := range board {
			for column := range board[row] {
				board[row][column] = next_cell(row, column, board, counts)
			}
		}
	}
}
