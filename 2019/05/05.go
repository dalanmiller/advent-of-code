package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Operation int

const (
	ADD       Operation = 1
	MULT                = 2
	STORE               = 3
	OUTPUT              = 4
	JUMPTRUE            = 5
	JUMPFALSE           = 6
	LESSTHAN            = 7
	EQUALS              = 8
)

var operationsMap = map[int]Operation{
	1: ADD,
	2: MULT,
	3: STORE,
	4: OUTPUT,
	5: JUMPTRUE,
	6: JUMPFALSE,
	7: LESSTHAN,
	8: EQUALS,
}

var operationsParamMap = map[Operation]int{
	ADD:       3,
	MULT:      3,
	STORE:     1,
	OUTPUT:    1,
	JUMPTRUE:  2,
	JUMPFALSE: 2,
	LESSTHAN:  3,
	EQUALS:    3,
}

type Mode int

const (
	POSITION  Mode = 0
	IMMEDIATE Mode = 1
)

var modeMap = map[string]Mode{
	"0": POSITION,
	"1": IMMEDIATE,
}

type AddInstruction struct{}

func (ai AddInstruction) execute(cur_pos int, params []int, intcodes []int) {
	intcodes[intcodes[cur_pos+3]] = params[0] + params[1]
}

type MultInstruction struct{}

func (mi MultInstruction) execute(cur_pos int, params []int, intcodes []int) {
	intcodes[intcodes[cur_pos+3]] = params[0] * params[1]
}

type StoreInstruction struct{}

func (si StoreInstruction) execute(input int, cur_pos int, params []int, intcodes []int) {
	intcodes[intcodes[cur_pos+1]] = input
}

type OutputInstruction struct{}

func (oi OutputInstruction) execute(cur_pos int, params []int, intcodes []int) int {
	return params[0]
}

type JumpIfTrueInstruction struct{}

func (jiti JumpIfTrueInstruction) execute(cur_pos int, params []int, intcodes []int) (int, bool) {
	if params[0] != 0 {
		return params[1], true
	}
	return cur_pos, false
}

type JumpIfFalseInstruction struct{}

func (jifi JumpIfFalseInstruction) execute(cur_pos int, params []int, intcodes []int) (int, bool) {
	if params[0] == 0 {
		return params[1], true
	}
	return cur_pos, false
}

type LessThanInstruction struct{}

func (lti LessThanInstruction) execute(cur_pos int, params []int, intcodes []int) {
	if params[0] < params[1] {
		intcodes[intcodes[cur_pos+3]] = 1
	} else {
		intcodes[intcodes[cur_pos+3]] = 0
	}
}

type EqualsInstruction struct{}

func (ei EqualsInstruction) execute(cur_pos int, params []int, intcodes []int) {
	if params[0] == params[1] {
		intcodes[intcodes[cur_pos+3]] = 1
	} else {
		intcodes[intcodes[cur_pos+3]] = 0
	}
}

func parse_instruction(instructions int) (Operation, string) {
	in_str := fmt.Sprint(instructions)
	in_length := len(in_str)
	var opcode int
	var err error
	var parameters string
	if len(in_str) > 1 {
		opcode, err = strconv.Atoi(in_str[in_length-2:])
		parameters = in_str[:in_length-2]
	} else {
		opcode, err = strconv.Atoi(in_str)
	}

	if err != nil {
		log.Fatalf("Could not convert to integer, %s", err)
	}

	operation := operationsMap[opcode]

	return operation, parameters

}

func run(input int, program string) (int, []int) {

	var intcodes []int
	for _, s := range strings.Split(program, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not convert to integer, %s", err)
		}
		intcodes = append(intcodes, i)
	}

	memory := input
	cur_pos := 0
	var intcode int
	// for i, intcode := range intcodes {
	for intcodes[cur_pos] != 99 {
		intcode = intcodes[cur_pos]
		operation, parameters := parse_instruction(intcode)

		expectedLength := operationsParamMap[operation]
		paramLength := len(parameters)
		if paramLength < expectedLength && paramLength >= 1 {
			string_params, err := strconv.Atoi(parameters)
			if err != nil {
				log.Fatalf("Could not convert params to integer, %s", err)
			}
			parameters = fmt.Sprintf("%0*d", expectedLength, string_params)
		} else if paramLength == 0 {
			parameters = fmt.Sprintf("%0*d", expectedLength, 0)
		}

		var modes []Mode
		// Go in reverse order
		for i := len(parameters) - 1; i >= 0; i-- {
			modes = append(modes, modeMap[string(parameters[i])])
		}

		parameter_values := []int{}
		for i, mode := range modes {
			switch mode {
			case POSITION:
				parameter_values = append(parameter_values, intcodes[intcodes[cur_pos+i+1]])
			case IMMEDIATE:
				parameter_values = append(parameter_values, intcodes[cur_pos+i+1])
			}
		}

		jumped := false
		switch operation {
		case ADD:
			AddInstruction{}.execute(cur_pos, parameter_values, intcodes)
		case MULT:
			MultInstruction{}.execute(cur_pos, parameter_values, intcodes)
		case STORE:
			StoreInstruction{}.execute(input, cur_pos, parameter_values, intcodes)
		case OUTPUT:
			memory = OutputInstruction{}.execute(cur_pos, parameter_values, intcodes)
		case JUMPTRUE:
			cur_pos, jumped = JumpIfTrueInstruction{}.execute(cur_pos, parameter_values, intcodes)
		case JUMPFALSE:
			cur_pos, jumped = JumpIfFalseInstruction{}.execute(cur_pos, parameter_values, intcodes)
		case LESSTHAN:
			LessThanInstruction{}.execute(cur_pos, parameter_values, intcodes)
		case EQUALS:
			EqualsInstruction{}.execute(cur_pos, parameter_values, intcodes)
		}

		if jumped == false {
			cur_pos += operationsParamMap[operation] + 1
		}
	}

	return memory, intcodes
}
