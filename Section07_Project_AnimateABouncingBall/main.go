package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	var (
		x, y   = 20, 20
		lx, ly = 0,6
		ball   = '⚽'
		space  = ' '
		vx, vy = 1, 1
	)
	board := make([][]bool, x)

	for row := range board {
		board[row] = make([]bool, y)
	}

	for{
		lx += vx
		ly += vy

		board[lx][ly] = true

		if lx <= 0 || lx >= y -1{
			vx*= -1
		}

		if ly <= 0 || ly >= x -1{
			vy*= -1
		}

		buffer := make([]rune, x*y)

		buffer = buffer[:0]

		screen.Clear()
		
		for r := range board[0] {
			for c := range board {
				cell := space

				if board[r][c] {
					cell = ball
				}

				buffer = append(buffer, cell, space)
			}

			buffer = append(buffer, '\n')
		}

		screen.MoveTopLeft()

		fmt.Print(string(buffer))

		board[lx][ly] = false

		time.Sleep(time.Second / 20)

	}

}
