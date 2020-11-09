package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	moves := 1.0
	for i := 0; i < k; i++ {
		moves *= 8.0
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			sum := onemove(i, j, k)
			p := float64(sum) / moves
			fmt.Printf("%.04f ", p)
			if j == 7 {
				fmt.Println()
			}
		}
	}
}

var onBoardCnt = [8][8]int{
	{2, 3, 4, 4, 4, 4, 3, 2},
	{3, 4, 6, 6, 6, 6, 4, 3},
	{4, 6, 8, 8, 8, 8, 6, 4},
	{4, 6, 8, 8, 8, 8, 6, 4},
	{4, 6, 8, 8, 8, 8, 6, 4},
	{4, 6, 8, 8, 8, 8, 6, 4},
	{3, 4, 6, 6, 6, 6, 4, 3},
	{2, 3, 4, 4, 4, 4, 3, 2},
}

var moveIncr = [][2]int{
	{-1, 2}, {1, 2},
	{-1, -2}, {1, -2},
	{2, -1}, {2, 1},
	{-2, -1}, {-2, 1},
}

func onemove(i, j, k int) int {
	if k == 1 {
		return onBoardCnt[i][j]
	}

	sum := 0
	for _, increment := range moveIncr {
		xNext := i + increment[0]
		yNext := j + increment[1]
		if 0 <= xNext && xNext <= 7 {
			if 0 <= yNext && yNext <= 7 {
				sum += onemove(xNext, yNext, k-1)
			}
		}
	}

	return sum
}
