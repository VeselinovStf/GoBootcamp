// Package main implements Retro Clock
package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	// noColor := placeholder{
	// 	"   ",
	// 	"   ",
	// 	"   ",
	// 	"   ",
	// 	"   ",
	// }

	nums := []placeholder{
		zero, one, two, three, four,
		five, six, seven, eight, nine,
	}

	screen.Clear()

	for {

		screen.MoveTopLeft()

		h, m, s := time.Now().Clock()

		timeElements := [...]placeholder{
			nums[h/10], nums[h%10],
			colon,
			nums[m/10], nums[m%10],
			colon,
			nums[s/10], nums[s%10],
		}

		for line := range timeElements[0] {
			for index, digit := range timeElements {
				next := timeElements[index][line]
				if digit == colon && s%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
	}
}
