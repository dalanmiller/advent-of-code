package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) manhattan() int {
	x := c.X
	y := c.Y

	if c.X < 0 {
		x = -c.X
	}

	if c.Y < 0 {
		y = -c.Y
	}

	return x + y
}

type Direction string

const (
	UP    Direction = "U"
	DOWN            = "D"
	RIGHT           = "R"
	LEFT            = "L"
)

var direction_map = map[string]Direction{
	"U": UP,
	"D": DOWN,
	"L": LEFT,
	"R": RIGHT,
}

type Line struct {
	Begin       Coordinate
	End         Coordinate
	Direction   Direction
	Coordinates []Coordinate
}

func (l *Line) coordinates() []Coordinate {
	// if len(l.Coordinates) != 0 {
	// 	return l.Coordinates
	// }

	coordinates := []Coordinate{}
	current_position := l.Begin

	coordinates = append(coordinates, current_position)

	for current_position != l.End {
		switch l.Direction {
		case UP:
			current_position.Y++
		case RIGHT:
			current_position.X++
		case DOWN:
			current_position.Y--
		case LEFT:
			current_position.X--
		}
		coordinates = append(coordinates, current_position)
	}

	// l.Coordinates = coordinates
	return coordinates
}

type Path struct {
	Lines []Line
}

func (p *Path) full_coordinates() []Coordinate {
	full_coordinates := []Coordinate{}
	added := make(map[Coordinate]bool)

	for _, line := range p.Lines {
		coords := line.coordinates()

		for _, coord := range coords {
			_, ok_added := added[coord]
			if !ok_added {
				full_coordinates = append(full_coordinates, coord)
				added[coord] = true
			}
		}

	}

	return full_coordinates
}

func (p1 *Path) intersect(p2 *Path) map[Coordinate]int {
	seen := make(map[Coordinate]int)
	intersection := make(map[Coordinate]int)

	step := 0
	for _, coord := range p1.full_coordinates() {
		seen[coord] = step
		step++
	}

	step = 0
	for _, coord := range p2.full_coordinates() {
		p1_step_val, seen_ok := seen[coord]
		_, added_ok := intersection[coord]
		if seen_ok && !added_ok {
			intersection[coord] = step + p1_step_val
		}
		step++
	}

	delete(intersection, Coordinate{X: 0, Y: 0})

	return intersection
}

func run(lines string) (int, int) {

	wires_paths := strings.Split(string(lines), "\n")
	wires := make(map[int]*Path)

	for i := range wires_paths {
		wires[i] = &Path{}
	}

	for i, wire := range wires_paths {
		pos := Coordinate{X: 0, Y: 0}
		for _, dir := range strings.Split(wire, ",") {

			start := Coordinate{X: pos.X, Y: pos.Y}
			direction := direction_map[string(dir[0])]
			distance, err := strconv.Atoi(dir[1:])
			if err != nil {
				log.Fatalf("Could not convert to int %s", err)
			}
			var end Coordinate
			switch direction {
			case UP:
				end = Coordinate{X: pos.X, Y: pos.Y + distance}
			case RIGHT:
				end = Coordinate{X: pos.X + distance, Y: pos.Y}
			case DOWN:
				end = Coordinate{X: pos.X, Y: pos.Y - distance}
			case LEFT:
				end = Coordinate{X: pos.X - distance, Y: pos.Y}
			}
			line := Line{
				Begin:     start,
				End:       end,
				Direction: direction,
			}

			wires[i].Lines = append(wires[i].Lines, line)
			pos = end
		}
	}

	full_intersects := make(map[Coordinate]int)
	for _, path1 := range wires {
		for _, path2 := range wires {
			if path1 != path2 {
				for coord, steps := range path1.intersect(path2) {
					full_intersects[coord] = steps
				}
			}
		}
	}

	min := math.MaxInt
	min_combined_steps := math.MaxInt
	for coord, steps := range full_intersects {
		manhattan_distance := coord.manhattan()
		if manhattan_distance < min {
			min = manhattan_distance
		}

		if steps < min_combined_steps {
			min_combined_steps = steps
		}
	}

	return min, min_combined_steps
}
