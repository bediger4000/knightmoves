package main

/*
 * Print out probability of 1, randomly-chosen, knight's move starting from
 * each of the 64 squares on the board, staying on the board.
 */

import "fmt"

type board [8][8]int

func main() {
	var b board
	b1 := onemove(b)
	printBoard(b1, 1)
}

func printBoard(bd board, k int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			n := bd[i][j]
			fmt.Printf("%2d ", n)
			if j == 7 {
				// end of row, print newline
				fmt.Println()
			}
		}
	}
	fmt.Println()
}

var moveIncr = [][2]int{
	{-1, 2}, {1, 2},
	{-1, -2}, {1, -2},
	{2, -1}, {2, 1},
	{-2, -1}, {-2, 1},
}

func onemove(bd board) board {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			for _, increment := range moveIncr {
				xNext := i + increment[0]
				yNext := j + increment[1]
				if xNext >= 0 && xNext <= 7 {
					if yNext >= 0 && yNext <= 7 {
						// knight at <i,j> still on the board
						bd[i][j]++
					}
				}
			}
		}
	}
	return bd
}
