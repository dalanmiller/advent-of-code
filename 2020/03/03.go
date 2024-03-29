package main

import (
	"fmt"

	. "../helpers"
)

func part1() {
	rows, _ := ReadInput("./input")

	x, y, trees := 0, 0, 0

	for y < len(rows) {
		x = (x + 3) % 31
		y++

		if y < len(rows) && rows[y][x] == '#' {
			trees++
		}
	}

	fmt.Println("Trees hit: ", trees)
}

type path struct {
	xInc int
	yInc int
}

func part2() {
	rows, _ := ReadInput("./input")

	paths := []path{
		path{xInc: 1, yInc: 1},
		path{xInc: 3, yInc: 1},
		path{xInc: 5, yInc: 1},
		path{xInc: 7, yInc: 1},
		path{xInc: 1, yInc: 2},
	}

	prod := 1

	for _, p := range paths {
		x, y, trees := 0, 0, 0
		for y < len(rows) {
			x = (x + p.xInc) % 31
			y = y + p.yInc

			if y < len(rows) && rows[y][x] == '#' {
				trees++
			}
		}

		prod *= trees
	}

	fmt.Print("Trees outcome product: ", prod)
}

func main() {
	part1()
	part2()
}
