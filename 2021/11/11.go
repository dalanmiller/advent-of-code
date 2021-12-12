package main

import (
	"strconv"
	"strings"
)

type Octo struct {
	N       int
	Adj     []*Octo
	Flashed bool
}

type Coord struct {
	X int
	Y int
}

func adjacentCoords(i, j, width, height int) []Coord {
	var fullCoords []Coord

	// top
	fullCoords = append(fullCoords, []Coord{
		{X: j - 1, Y: i - 1},
		{X: j, Y: i - 1},
		{X: j + 1, Y: i - 1},
	}...)

	// middle

	fullCoords = append(fullCoords, []Coord{
		{X: j - 1, Y: i},
		{X: j + 1, Y: i},
	}...)

	// bottom

	fullCoords = append(fullCoords, []Coord{
		{X: j - 1, Y: i + 1},
		{X: j, Y: i + 1},
		{X: j + 1, Y: i + 1},
	}...)

	var validCoords []Coord
	for _, coord := range fullCoords {
		if coord.X >= 0 && coord.X < width && coord.Y >= 0 && coord.Y < height {
			validCoords = append(validCoords, coord)
		}
	}

	return validCoords
}

func parseInput(input string) [][]*Octo {
	rows := strings.Split(input, "\n")

	grid := make([][]*Octo, len(rows))
	for i := range grid {
		grid[i] = make([]*Octo, len(rows[0]))
	}

	for i, row := range rows {
		for j, v := range row {
			n, _ := strconv.Atoi(string(v))

			grid[i][j] = &Octo{
				N:       n,
				Flashed: false,
			}
		}
	}

	width, height := len(grid[0]), len(grid)
	for i, row := range grid {
		for j, octo := range row {
			coords := adjacentCoords(i, j, width, height)
			for _, coord := range coords {
				octo.Adj = append(octo.Adj, grid[coord.Y][coord.X])
			}
		}
	}

	return grid
}

func flash(o *Octo) {
	o.Flashed = true
	for _, adjOcto := range o.Adj {
		adjOcto.N++
	}
}

func run(input string, days int) int {
	grid := parseInput(input)

	var flashes int

	for d := 0; d < days; d++ {
		// iterate each octo
		for _, row := range grid {
			for _, octo := range row {
				octo.N++
			}
		}

		// flash if > 9 and not already flashed
		flashed := true
		for flashed {
			flashed = false
			for _, row := range grid {
				for _, octo := range row {
					if octo.N > 9 && !octo.Flashed {
						flash(octo)
						flashed = true
					}
				}
			}
		}

		// Reset
		allFlashed := true
		for _, row := range grid {
			for _, octo := range row {
				if octo.Flashed {
					flashes++
					octo.N = 0
					octo.Flashed = false
				} else {
					allFlashed = false
				}
			}
		}

		if allFlashed {
			return d + 1
		}
		// fmt.Println()
		// for i := 0; i < len(grid); i++ {
		// 	for j := 0; j < len(grid[0]); j++ {
		// 		fmt.Printf("%d", grid[i][j].N)
		// 	}
		// 	fmt.Printf("\n")
		// }
	}

	return flashes
}
