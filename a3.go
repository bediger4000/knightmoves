package main

/*
 * After k moves, how many knights stay on the board when starting at a given
 * position?
 *
 * Actually do the work of trying paths no more than k-moves long and seeing
 * which ones stay on the board.
 */

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("%s k n\nRun k-length knight's paths n times per square\n", os.Args[0])
		return
	}
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	cnt, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Looking %d moves deep, %d times\n", k, cnt)

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			onBoardCount := 0
			for n := 0; n < cnt; n++ {
				if onemove(i, j, k) {
					onBoardCount++
				}
			}
			percent := float64(onBoardCount) / float64(cnt)
			fmt.Printf("%.04f ", percent)
			if j == 7 {
				fmt.Println()
			}
		}
	}
}

var moveIncr = [][2]int{
	{-1, 2}, {1, 2},
	{-1, -2}, {1, -2},
	{2, -1}, {2, 1},
	{-2, -1}, {-2, 1},
}

func onemove(i, j, k int) bool {
	k--
	increment := moveIncr[rand.Intn(8)]
	xNext := i + increment[0]
	yNext := j + increment[1]
	if xNext >= 0 && xNext <= 7 {
		if yNext >= 0 && yNext <= 7 {
			// knight at <i,j> still on the board
			if k == 0 {
				// did k moves, still on the board
				return true
			}
			return onemove(xNext, yNext, k)
		}
	}
	return false
}
