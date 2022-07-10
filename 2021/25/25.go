package main

import (
	"log"
	"strings"
)

type Direction rune

const (
	RIGHT Direction = '>'
	DOWN  Direction = 'v'
)

type Position [2]int

type Cucumber struct {
	D Direction
	P Position
}

var positionMap = map[Position]*Cucumber{}

func parseInput(input string) ([]*Cucumber, []*Cucumber) {
	lines := strings.Split(input, "\n")
	rc := []*Cucumber{}
	dc := []*Cucumber{}

	for y, line := range lines {
		for x, char := range line {
			if char != '.' {

				p := Position{x, y}
				var c Cucumber

				switch char {
				case '>':
					c = Cucumber{
						Direction(char), p,
					}
					rc = append(rc, &c)
				case 'v':
					c = Cucumber{
						Direction(char), p,
					}
					dc = append(dc, &c)
				}

				positionMap[p] = &c
			}
		}
	}

	return rc, dc
}

var GRID_WIDTH int
var GRID_HEIGHT int

func step(rc []*Cucumber, dc []*Cucumber) bool {

	// Check if all RightCucs can move
	canRightMove := make([]*Cucumber, 0, len(rc))
	for _, c := range rc {
		r := c.P[0]
		if r == GRID_WIDTH-1 && positionMap[Position{0, c.P[1]}] == nil {
			canRightMove = append(canRightMove, c)
		} else if _, ok := positionMap[Position{r + 1, c.P[1]}]; !ok && r+1 < GRID_WIDTH {
			canRightMove = append(canRightMove, c)
		}
	}

	// Move RightCucs
	for _, c := range canRightMove {
		newP := Position{(c.P[0] + 1) % GRID_WIDTH, c.P[1]}
		delete(positionMap, c.P)
		c.P[0] = newP[0]
		positionMap[c.P] = c
	}

	// Check if all DownCucs can move
	canDownMove := make([]*Cucumber, 0, len(dc))
	for _, c := range dc {
		d := c.P[1]
		if _, ok := positionMap[Position{c.P[0], d + 1}]; !ok && d+1 < GRID_HEIGHT {
			canDownMove = append(canDownMove, c)
		} else if d == GRID_HEIGHT-1 && positionMap[Position{c.P[0], 0}] == nil {
			canDownMove = append(canDownMove, c)
		}
	}

	// Move DownCucs
	for _, c := range canDownMove {
		newP := Position{c.P[0], (c.P[1] + 1) % GRID_HEIGHT}
		delete(positionMap, c.P)
		c.P[1] = newP[1]
		positionMap[c.P] = c
	}

	return !(len(canRightMove) == 0 && len(canDownMove) == 0)
}

func printGrid() string {
	grid := make([][]string, GRID_HEIGHT)

	for i := range grid {
		grid[i] = make([]string, GRID_WIDTH)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	for k, v := range positionMap {
		grid[k[1]][k[0]] = string(v.D)
	}

	var b strings.Builder
	for y, row := range grid {
		b.WriteString(strings.Join(row, ""))
		if len(grid)-1 != y {
			b.WriteString("\n")
		}
	}

	return b.String()
}

func run(input string) int {
	rightCucs, downCucs := parseInput(input)
	GRID_WIDTH = strings.Index(input, "\n")
	GRID_HEIGHT = len(strings.Split(input, "\n"))

	n := 1
	for step(rightCucs, downCucs) {
		n++
		log.Printf("After %d iteration(s)\n\n%s", n, printGrid())
	}

	return n
}
