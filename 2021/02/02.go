package main

import (
	"math"
	"strconv"
	"strings"
)

type Direction string

const (
	FORWARD Direction = "forward"
	UP      Direction = "up"
	DOWN    Direction = "down"
)

type Movement struct {
	direction Direction
	value     int
}

func parse_input(input string) *[]Movement {
	lines := strings.Split(input, "\n")

	movements := make([]Movement, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")

		direction := split[0]
		n, _ := strconv.Atoi(split[1])

		movements[i] = Movement{
			direction: Direction(direction),
			value:     n,
		}
	}

	return &movements
}

func run(input string) int {

	movements := parse_input(input)
	horiz, depth := 0, 0
	for _, movement := range *movements {
		switch movement.direction {
		case FORWARD:
			horiz += movement.value
		case DOWN:
			depth -= movement.value
		case UP:
			depth += movement.value
		}
	}

	return horiz * int(math.Abs(float64(depth)))
}

func run_two(input string) int {
	movements := parse_input(input)
	aim, horiz, depth := 0, 0, 0
	for _, movement := range *movements {
		switch movement.direction {
		case FORWARD:
			horiz += movement.value
			depth += aim * movement.value
		case DOWN:
			aim += movement.value
		case UP:
			aim -= movement.value
		}
	}

	return horiz * int(math.Abs(float64(depth)))
}
