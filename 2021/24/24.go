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

// var functionMap = map[Operation]func(o1, o2 int16) int16{
// 	ADD:      func(o1, o2 int16) int16 { return o1 + o2 },
// 	MULTIPLY: func(o1, o2 int16) int16 { return o1 * o2 },
// 	DIVIDE:   func(o1, o2 int16) int16 { return o1 / o2 },
// 	EQUAL: func(o1, o2 int16) int16 {
// 		if o1 == o2 {
// 			return 1
// 		} else {
// 			return 0
// 		}
// 	},
// 	MODULO: func(o1, o2 int16) int16 { return o1 % o2 },
// }

// type Computer struct {
// 	Numbers      [14]int16
// 	Instructions []Instruction
// 	Memory       [4]int16
// 	P            int // Progress into instruction set
// 	N            int
// }

// func (c *Computer) Cycle() bool {
// 	instruction := c.Instructions[c.P]

// 	if Operation(instruction.Op) == INPUT {
// 		c.Memory[instruction.Var] = c.Numbers[c.N]
// 		c.N++
// 		// If no error, we then we have a value, otherwise we have a variable
// 	} else if v, err := strconv.ParseInt(instruction.Val, 10, 16); err == nil {

// 		c.Memory[instruction.Var] = functionMap[instruction.Op](
// 			c.Memory[instruction.Var], int16(v),
// 		)
// 	} else {
// 		c.Memory[instruction.Var] = functionMap[instruction.Op](
// 			c.Memory[instruction.Var], c.Memory[registerMap[instruction.Val]],
// 		)
// 	}
// 	c.P++

// 	if c.P >= len(c.Instructions) {
// 		return false
// 	}
// 	return true
// }

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

// func validateMonadNumber(i int) bool {
// 	nums := strings.Split(strconv.Itoa(i), "")

// 	for _, num := range nums {
// 		if num == "0" {
// 			return false
// 		}
// 	}
// 	return true
// }

// func ALU(nums [14]int16, instructions []Instruction) int16 {
// 	c := Computer{
// 		nums,
// 		instructions,
// 		[4]int16{},
// 		0,
// 		0,
// 	}

// 	for c.Cycle() {

// 	}

// 	return c.Memory[Register(3)]
// }

// var failMap = map[int]map[int]bool{}
// var LARGEST = [14]int{}
// var CURRENT = [14]int{}

// func ALU2(i, memory int) bool {

// 	prevMemory := memory

// 	if failMap[i] != nil && failMap[i][prevMemory] {
// 		return false
// 	}

// 	for j := 9; j > 0; j-- {
// 		CURRENT[i] = j
// 		z := block(
// 			j, memory, diffs[i][0], diffs[i][1], diffs[i][2],
// 		)

// 		if i < 13 {
// 			if ALU2(i+1, z) {
// 				LARGEST[i] = j
// 				return true
// 			}
// 		} else if z == 0 {
// 			LARGEST[i] = j
// 			return true
// 		}
// 	}

// 	if failMap[i] == nil {
// 		failMap[i] = make(map[int]bool, 1000)
// 	}
// 	failMap[i][prevMemory] = false

// 	return false
// }

// Put into CSV
// $ xsv table blocks.csv
//
// 1         2         3         4          5         6         7         8         9          10        11         12        13        14
// inp w     inp w     inp w     inp w      inp w     inp w     inp w     inp w     inp w      inp w     inp w      inp w     inp w     inp w
// mul x 0   mul x 0   mul x 0   mul x 0    mul x 0   mul x 0   mul x 0   mul x 0   mul x 0    mul x 0   mul x 0    mul x 0   mul x 0   mul x 0
// add x z   add x z   add x z   add x z    add x z   add x z   add x z   add x z   add x z    add x z   add x z    add x z   add x z   add x z
// mod x 26  mod x 26  mod x 26  mod x 26   mod x 26  mod x 26  mod x 26  mod x 26  mod x 26   mod x 26  mod x 26   mod x 26  mod x 26  mod x 26

// This below row determines if the instruction is one of two different behaviors
// . 1. Push the input variable + a constant into z or
// . 2. It pops a previously saved variable (plus constant) and another constant and
// . compared it with the current digit
// . if comparison is successful: z loses its lease significant base-26 digit
// . otherwise if not successful: push some other value into it
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
// var diffs = [14][3]int{
// 	{1, 15, 4},
// 	{1, 14, 16},
// 	{1, 11, 14},
// 	{26, -13, 3},
// 	{1, 14, 11},
// 	{1, 15, 13},
// 	{26, -7, 11},
// 	{1, 10, 7},
// 	{26, -12, 12},
// 	{1, 15, 15},
// 	{26, -16, 13},
// 	{26, -9, 1},
// 	{26, -8, 15},
// 	{26, -8, 4},
// }

// type blockInput struct {
// 	w, z, zdiv, xadd, yadd int
// }

// var blockMap = make(map[blockInput]int, 10000)

// Yup pulling out the functions from the input
// . and making it it's own function
// func block(w, z, zdiv, xadd, yadd int) int {
// var x, y int
// x = z
// x %= 26
// z /= zdiv
// x += xadd
// if x == int(w) {
// 	x = 0
// } else {
// 	x = 1
// }
// y = 25*x + 1
// z *= y
// y = int(w) + yadd
// y *= x
// z += y

// x := 1
// y := 0

// 	return z
// }

type Constraint struct {
	i, j, diff int
}

type Stack []Constraint

func (s *Stack) Push(c Constraint) {
	*s = append(*s, c)
}

func (s *Stack) Pop() Constraint {
	pop := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return pop
}

func generateConstraints(instructions []Instruction) []Constraint {
	constraints := []Constraint{}
	s := &Stack{}

	for i := 0; i < 14; i++ {

		iter := i * 18

		// Check this instruction which will determine the block type:
		// 1. This is a push block
		// 2. This is a pop block
		inst := instructions[iter+4]
		decisionV, _ := strconv.Atoi(inst.Val)

		if inst.Var == 3 && decisionV == 1 {
			inst := instructions[iter+15]
			v, _ := strconv.Atoi(inst.Val)

			if inst.Op != ADD && inst.Val != "y" {
				log.Fatal("Incorrect instruction")
			}

			s.Push(Constraint{
				i, 0, v,
			})

		} else {

			inst := instructions[iter+5]
			v, _ := strconv.Atoi(inst.Val)

			if inst.Op != ADD && inst.Val != "x" {
				log.Fatal("Incorrect instruction")
			}

			// Pop off the end and remove last element
			p := s.Pop()

			constraints = append(constraints, Constraint{
				i, p.i, p.diff + v,
			})

		}
	}

	return constraints
}

func run(input string) (int, int) {
	instructions := parseInput(input)

	constraints := generateConstraints(instructions)

	nmax := make([]int, 14)
	nmin := make([]int, 14)

	for _, constraint := range constraints {
		if constraint.diff > 0 {
			nmax[constraint.i], nmax[constraint.j] = 9, 9-constraint.diff
			nmin[constraint.i], nmin[constraint.j] = 1+constraint.diff, 1
		} else {
			nmax[constraint.i], nmax[constraint.j] = 9+constraint.diff, 9
			nmin[constraint.i], nmin[constraint.j] = 1, 1-constraint.diff
		}
	}

	maxNum, minNum := 0, 0
	for _, n := range nmax {
		maxNum = maxNum*10 + n
	}
	for _, n := range nmin {
		minNum = minNum*10 + n
	}

	return maxNum, minNum
}
