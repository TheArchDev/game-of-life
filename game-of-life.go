package main

import "fmt"


func main() {
	fmt.Println("Game of Life!")
	const x_dimension = 2
	const y_dimension = 2
	const dead_space = "O"
	const alive_space = "X"
	var board [x_dimension][y_dimension]string
	for x := 0; x < x_dimension; x ++ {
		for y := 0; y < y_dimension; y ++ {
	    	board[x][y] = dead_space
	    }
	}
	fmt.Println(board)
}
