package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type OperationType string

const (
	ADDX OperationType = "addx"
	NOOP OperationType = "noop"
)

var cycleMap = map[OperationType]int{
	ADDX: 2,
	NOOP: 1,
}

type operation struct {
	OperationType OperationType
	Value         int
}

func readInput(input io.Reader) []operation {
	s := bufio.NewScanner(input)

	ops := []operation{}
	for s.Scan() {
		split := strings.Split(s.Text(), " ")
		op := OperationType(split[0])

		var v int
		if op != NOOP {
			v, _ = strconv.Atoi(split[1])
		}

		ops = append(ops, operation{
			op,
			v,
		})
	}

	return ops
}

func run(input io.Reader) int {
	ops := readInput(input)

	x := 1
	cycle := 0
	signalStrength := 0

	// Create and fill screen
	screen := [6][40]string{}
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 39; j++ {
			screen[i][j] = "."
		}
	}

	for _, op := range ops {

		for i := 0; i < cycleMap[op.OperationType]; i++ {

			pixel := cycle % 40
			if pixel == x || pixel == x-1 || pixel == x+1 {
				line := int(cycle / 40)
				screen[line][pixel] = "#"
			}

			cycle++
			if (cycle-20)%40 == 0 {
				signalStrength += cycle * x
			}
		}

		switch op.OperationType {
		case ADDX:
			x += op.Value
		}
	}

	for _, line := range screen {
		var s strings.Builder
		for _, pixel := range line {
			s.WriteString(pixel)
		}
		fmt.Println(s.String())
	}

	return signalStrength
}
