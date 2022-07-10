package main

import (
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

func parseInput(input string) ([]*Cucumber, []*Cucumber, *map[Position]*Cucumber) {
	lines := strings.Split(input, "\n")
	rc := []*Cucumber{}
	dc := []*Cucumber{}
	positionMap := map[Position]*Cucumber{}

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

	return rc, dc, &positionMap
}

func step(rc, dc []*Cucumber, positionMap *map[Position]*Cucumber, gh, gw int) bool {

	// Check if all RightCucs can move
	canRightMove := make([]*Cucumber, 0, len(rc))
	for _, c := range rc {
		r := c.P[0]
		if r == gw-1 && (*positionMap)[Position{0, c.P[1]}] == nil {
			canRightMove = append(canRightMove, c)
		} else if _, ok := (*positionMap)[Position{r + 1, c.P[1]}]; !ok && r+1 < gw {
			canRightMove = append(canRightMove, c)
		}
	}

	// Move RightCucs
	for _, c := range canRightMove {
		newP := Position{(c.P[0] + 1) % gw, c.P[1]}
		delete((*positionMap), c.P)
		c.P[0] = newP[0]
		(*positionMap)[c.P] = c
	}

	// Check if all DownCucs can move
	canDownMove := make([]*Cucumber, 0, len(dc))
	for _, c := range dc {
		d := c.P[1]
		if _, ok := (*positionMap)[Position{c.P[0], d + 1}]; !ok && d+1 < gh {
			canDownMove = append(canDownMove, c)
		} else if d == gh-1 && (*positionMap)[Position{c.P[0], 0}] == nil {
			canDownMove = append(canDownMove, c)
		}
	}

	// Move DownCucs
	for _, c := range canDownMove {
		newP := Position{c.P[0], (c.P[1] + 1) % gh}
		delete((*positionMap), c.P)
		c.P[1] = newP[1]
		(*positionMap)[c.P] = c
	}

	return !(len(canRightMove) == 0 && len(canDownMove) == 0)
}

func printGrid(gh, gw int, positionMap *map[Position]*Cucumber) string {
	grid := make([][]string, gh)

	for i := range grid {
		grid[i] = make([]string, gw)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	for k, v := range *positionMap {
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
	rightCucs, downCucs, pMap := parseInput(input)
	gw := strings.Index(input, "\n")
	gh := len(strings.Split(input, "\n"))

	n := 1
	for step(rightCucs, downCucs, pMap, gh, gw) {
		n++
		// log.Printf("After %d iteration(s)\n\n%s", n, printGrid())
	}

	return n
}
