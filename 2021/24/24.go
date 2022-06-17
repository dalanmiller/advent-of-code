package main

import (
	"strconv"
	"strings"
)

type Operation string

const (
	INPUT    Operation = "inp"
	ADD      Operation = "add"
	MULTIPLY Operation = "mul"
	DIVIDE   Operation = "div"
	EQUAL    Operation = "eql"
	MODULO   Operation = "mod"
)

var functionMap = map[Operation]func(o1, o2 int) int{
	ADD:      func(o1, o2 int) int { return o1 + o2 },
	MULTIPLY: func(o1, o2 int) int { return o1 * o2 },
	DIVIDE:   func(o1, o2 int) int { return o1 / o2 },
	EQUAL: func(o1, o2 int) int {
		if o1 == o2 {
			return 1
		} else {
			return 0
		}
	},
	MODULO: func(o1, o2 int) int { return o1 % o2 },
}

type Instruction struct {
	Op  Operation
	Var string
	Val int
}

func parseInput(input string) []Instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]Instruction, 0, len(lines))

	for _, line := range lines {
		tokens := strings.Split(line, " ")

		switch len(tokens) {
		case 2:
			// inp z
			instructions = append(instructions, Instruction{
				Op:  Operation(tokens[0]),
				Var: tokens[1],
			})
		case 3:
			// add z w
			// mod z 2
			// div w 2

			val, _ := strconv.Atoi(tokens[2])
			instructions = append(instructions, Instruction{
				Op:  Operation(tokens[0]),
				Var: tokens[1],
				Val: val,
			})
		}
	}

	return instructions
}

func validateMonadNumber(i int) bool {
	nums := strings.Split(strconv.Itoa(i), "")

	for _, num := range nums {
		if num == "0" {
			return false
		}
	}
	return true
}

func run(input string) int {
	instructions := parseInput(input)

	for i := 999999999999; i >= 111111111111; i-- {
		// Check if i is valid, otherwise continue
		if !validateMonadNumber(i) {
			continue
		}

		nums := make([]int, 14)
		for i, s := range strings.Split(strconv.Itoa(i), "") {
			val, _ := strconv.Atoi(s)
			nums[i] = val
		}

		i := 0
		memory := map[string]int{}
		var last string
		for _, instruction := range instructions {
			if instruction.Op == INPUT {
				memory[instruction.Var] = nums[i]
				last = instruction.Var
				i++
			} else {
				memory[instruction.Var] = functionMap[instruction.Op](
					memory[last], instruction.Val,
				)
			}
		}
		if memory["z"] == 0 {
			return i
		}
	}

	return -1
}
