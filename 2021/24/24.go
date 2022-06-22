package main

import (
	"log"
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

type Register int8

const (
	none Register = iota - 1
	w
	x
	y
	z
)

var registerMap = map[string]Register{
	"w": w,
	"x": x,
	"y": y,
	"z": z,
}

var functionMap = map[Operation]func(o1, o2 int64) int64{
	ADD:      func(o1, o2 int64) int64 { return o1 + o2 },
	MULTIPLY: func(o1, o2 int64) int64 { return o1 * o2 },
	DIVIDE:   func(o1, o2 int64) int64 { return o1 / o2 },
	EQUAL: func(o1, o2 int64) int64 {
		if o1 == o2 {
			return 1
		} else {
			return 0
		}
	},
	MODULO: func(o1, o2 int64) int64 { return o1 % o2 },
}

type Computer struct {
	Numbers      [14]int64
	Instructions []Instruction
	Memory       [4]int64
	P            int // Progress into instruction set
	N            int
}

func (c *Computer) Cycle() bool {
	instruction := c.Instructions[c.P]

	if Operation(instruction.Op) == INPUT {
		c.Memory[instruction.Var] = c.Numbers[c.N]
		c.N++
		// If no error, we then we have a value, otherwise we have a variable
	} else if v, err := strconv.ParseInt(instruction.Val, 10, 64); err == nil {

		c.Memory[instruction.Var] = functionMap[instruction.Op](
			c.Memory[instruction.Var], v,
		)
	} else {
		c.Memory[instruction.Var] = functionMap[instruction.Op](
			c.Memory[instruction.Var], c.Memory[registerMap[instruction.Val]],
		)
	}
	c.P++

	if c.P >= len(c.Instructions) {
		return false
	}
	return true
}

type Instruction struct {
	Op  Operation
	Var Register
	Val string
}

func parseInput(input string) []Instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]Instruction, 0, len(lines))

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		// variable, _ := strconv.Atoi(tokens[1])

		switch len(tokens) {
		case 2:
			// inp z
			instructions = append(instructions, Instruction{
				Op:  Operation(tokens[0]),
				Var: registerMap[tokens[1]],
			})
		case 3:
			// add z w
			// mod z 2
			// div w 2

			instructions = append(instructions, Instruction{
				Op:  Operation(tokens[0]),
				Var: registerMap[tokens[1]],
				Val: tokens[2],
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

func ALU(nums [14]int64, instructions []Instruction) int64 {
	c := Computer{
		nums,
		instructions,
		[4]int64{},
		0,
		0,
	}

	for c.Cycle() {

	}

	return c.Memory[Register(3)]
}

// Put into CSV
// $ xsv table blocks.csv
//
// 1         2         3         4          5         6         7         8         9          10        11         12        13        14
// inp w     inp w     inp w     inp w      inp w     inp w     inp w     inp w     inp w      inp w     inp w      inp w     inp w     inp w
// mul x 0   mul x 0   mul x 0   mul x 0    mul x 0   mul x 0   mul x 0   mul x 0   mul x 0    mul x 0   mul x 0    mul x 0   mul x 0   mul x 0
// add x z   add x z   add x z   add x z    add x z   add x z   add x z   add x z   add x z    add x z   add x z    add x z   add x z   add x z
// mod x 26  mod x 26  mod x 26  mod x 26   mod x 26  mod x 26  mod x 26  mod x 26  mod x 26   mod x 26  mod x 26   mod x 26  mod x 26  mod x 26

// div z 1   div z 1   div z 1   div z 26   div z 1   div z 1   div z 26  div z 1   div z 26   div z 1   div z 26   div z 26  div z 26  div z 26

// add x 15  add x 14  add x 11  add x -13  add x 14  add x 15  add x -7  add x 10  add x -12  add x 15  add x -16  add x -9  add x -8  add x -8

// eql x w   eql x w   eql x w   eql x w    eql x w   eql x w   eql x w   eql x w   eql x w    eql x w   eql x w    eql x w   eql x w   eql x w
// eql x 0   eql x 0   eql x 0   eql x 0    eql x 0   eql x 0   eql x 0   eql x 0   eql x 0    eql x 0   eql x 0    eql x 0   eql x 0   eql x 0
// mul y 0   mul y 0   mul y 0   mul y 0    mul y 0   mul y 0   mul y 0   mul y 0   mul y 0    mul y 0   mul y 0    mul y 0   mul y 0   mul y 0
// add y 25  add y 25  add y 25  add y 25   add y 25  add y 25  add y 25  add y 25  add y 25   add y 25  add y 25   add y 25  add y 25  add y 25
// mul y x   mul y x   mul y x   mul y x    mul y x   mul y x   mul y x   mul y x   mul y x    mul y x   mul y x    mul y x   mul y x   mul y x
// add y 1   add y 1   add y 1   add y 1    add y 1   add y 1   add y 1   add y 1   add y 1    add y 1   add y 1    add y 1   add y 1   add y 1
// mul z y   mul z y   mul z y   mul z y    mul z y   mul z y   mul z y   mul z y   mul z y    mul z y   mul z y    mul z y   mul z y   mul z y
// mul y 0   mul y 0   mul y 0   mul y 0    mul y 0   mul y 0   mul y 0   mul y 0   mul y 0    mul y 0   mul y 0    mul y 0   mul y 0   mul y 0
// add y w   add y w   add y w   add y w    add y w   add y w   add y w   add y w   add y w    add y w   add y w    add y w   add y w   add y w

// add y 4   add y 16  add y 14  add y 3    add y 11  add y 13  add y 11  add y 7   add y 12   add y 15  add y 13   add y 1   add y 15  add y 4

// mul y x   mul y x   mul y x   mul y x    mul y x   mul y x   mul y x   mul y x   mul y x    mul y x   mul y x    mul y x   mul y x   mul y x
// add z y   add z y   add z y   add z y    add z y   add z y   add z y   add z y   add z y    add z y   add z y    add z y   add z y   add z y

// Every 6th line "div z N"
// Every 7th line "add x N"
// Every 16th line "add y N"
var diffs = [14][3]int{
	{1, 15, 4},
	{1, 14, 16},
	{1, 11, 14},
	{26, -13, 3},
	{1, 14, 11},
	{1, 15, 13},
	{26, -7, 11},
	{1, 10, 7},
	{26, -12, 12},
	{1, 15, 15},
	{26, -16, 13},
	{26, -9, 1},
	{26, -8, 15},
	{26, -8, 4},
}

// Yup pulling out the functions from the input
// . and making it it's own native function
func block(w, x, y, z, zdiv, xadd, yadd int) int {
	x = z
	x %= 26
	z /= zdiv
	x += xadd
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = 25*x + 1
	z *= y
	y = w + yadd
	y *= x
	z += y

	return z
}

func run(input string) int {
	instructions := parseInput(input)

	for i := 99999999999999; i >= 11111111111111; i-- {
		// Check if i is valid, otherwise continue
		if !validateMonadNumber(i) {
			continue
		}
		log.Print(i)

		nums := [14]int64{}
		for j, s := range strings.Split(strconv.Itoa(i), "") {
			val, _ := strconv.ParseInt(s, 10, 64)
			nums[j] = val
		}

		if ALU(nums, instructions) == 0 {
			return i
		}
	}

	return -1
}
