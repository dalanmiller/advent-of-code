package main

import (
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type Line struct {
	start       Coordinate
	end         Coordinate
	coordinates []Coordinate
}

func (l *Line) calculate_coordinates(diagonal bool) {
	l.coordinates = []Coordinate{
		l.start,
	}

	// We're going vertical
	if l.start.x == l.end.x {
		pointer := l.start.y

		for pointer != l.end.y {
			if l.start.y < l.end.y {
				pointer++
			} else {
				pointer--
			}

			l.coordinates = append(l.coordinates, Coordinate{
				x: l.start.x, y: pointer,
			})
		}

		// We're going horizontal
	} else if l.start.y == l.end.y {
		pointer := l.start.x
		for pointer != l.end.x {
			if l.start.x < l.end.x {
				pointer++
			} else {
				pointer--
			}

			l.coordinates = append(l.coordinates, Coordinate{
				x: pointer, y: l.start.y,
			})
		}

		// We're going diagonal
	} else if diagonal {

		pointer_x := l.start.x
		pointer_y := l.start.y

		for pointer_x != l.end.x && pointer_y != l.end.y {
			if l.start.x < l.end.x {
				pointer_x++
			} else {
				pointer_x--
			}

			if l.start.y < l.end.y {
				pointer_y++
			} else {
				pointer_y--
			}

			l.coordinates = append(l.coordinates, Coordinate{
				x: pointer_x, y: pointer_y,
			})
		}
	}
}

func parse_lines(input string, diagonal bool) []Line {
	lines := strings.Split(input, "\n")

	grid_lines := make([]Line, len(lines))
	for i, line := range lines {

		split := strings.Split(line, " -> ")
		start := strings.Split(split[0], ",")
		start_x, _ := strconv.Atoi(start[0])
		start_y, _ := strconv.Atoi(start[1])
		end := strings.Split(split[1], ",")
		end_x, _ := strconv.Atoi(end[0])
		end_y, _ := strconv.Atoi(end[1])

		new_line := Line{
			start: Coordinate{
				x: start_x,
				y: start_y,
			},
			end: Coordinate{
				x: end_x,
				y: end_y,
			},
		}
		new_line.calculate_coordinates(diagonal)

		grid_lines[i] = new_line
	}

	return grid_lines
}

func run(input string) int {
	lines := parse_lines(input, false)

	grid := make([][]int, 999)
	for i := range grid {
		grid[i] = make([]int, 999)
	}

	for _, line := range lines {
		if len(line.coordinates) == 1 {
			continue
		}
		for _, coord := range line.coordinates {
			grid[coord.x][coord.y]++
		}
	}

	count := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] > 1 {
				count++
			}
		}
	}

	return count
}

func run_two(input string) int {
	lines := parse_lines(input, true)

	grid := make([][]int, 999)
	for i := range grid {
		grid[i] = make([]int, 999)
	}

	for _, line := range lines {
		for _, coord := range line.coordinates {
			grid[coord.x][coord.y]++
		}
	}

	count := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] > 1 {
				count++
			}
		}
	}

	return count
}
