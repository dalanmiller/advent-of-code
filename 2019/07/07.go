package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Operation int

const (
	ADD Operation = iota + 1
	MULT
	INPUT
	OUTPUT
	JUMPTRUE
	JUMPFALSE
	LESSTHAN
	EQUALS
)

var operationsMap = map[int]Operation{
	1: ADD,
	2: MULT,
	3: INPUT,
	4: OUTPUT,
	5: JUMPTRUE,
	6: JUMPFALSE,
	7: LESSTHAN,
	8: EQUALS,
}

var operationsParamMap = map[Operation]int{
	ADD:       3,
	MULT:      3,
	INPUT:     1,
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

type InputInstruction struct{}

func (ii InputInstruction) execute(input int, cur_pos int, params []int, intcodes []int) {
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

func parse_instruction(instructions int, cur_pos int, intcodes []int) (Operation, []int) {
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
		log.Fatalf("Could not convert to integer, %s, %s", err, in_str)
	}

	operation := operationsMap[opcode]

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

	return operation, parameter_values

}

func generate_intcodes(program string) []int {
	var intcodes []int
	for _, s := range strings.Split(program, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not convert to integer, %s", err)
		}
		intcodes = append(intcodes, i)
	}

	return intcodes
}

func run_program(ps *ProgramState) {
	var intcode, value int
	for ps.Intcodes[ps.CurPos] != 99 {
		intcode = ps.Intcodes[ps.CurPos]
		operation, parameters := parse_instruction(intcode, ps.CurPos, ps.Intcodes)

		jumped := false
		switch operation {
		case ADD:
			AddInstruction{}.execute(ps.CurPos, parameters, ps.Intcodes)
		case MULT:
			MultInstruction{}.execute(ps.CurPos, parameters, ps.Intcodes)
		case INPUT:
			switch ps.InputCount {
			case 1:
				value = ps.Phase
			case 2:
				value = ps.Input
			default:
				value = ps.Memory
			}
			InputInstruction{}.execute(value, ps.CurPos, parameters, ps.Intcodes)
			ps.InputCount++
		case OUTPUT:
			ps.Memory = OutputInstruction{}.execute(ps.CurPos, parameters, ps.Intcodes)
		case JUMPTRUE:
			ps.CurPos, jumped = JumpIfTrueInstruction{}.execute(ps.CurPos, parameters, ps.Intcodes)
		case JUMPFALSE:
			ps.CurPos, jumped = JumpIfFalseInstruction{}.execute(ps.CurPos, parameters, ps.Intcodes)
		case LESSTHAN:
			LessThanInstruction{}.execute(ps.CurPos, parameters, ps.Intcodes)
		case EQUALS:
			EqualsInstruction{}.execute(ps.CurPos, parameters, ps.Intcodes)
		}

		if jumped == false {
			ps.CurPos += operationsParamMap[operation] + 1
		}

		if operation == OUTPUT && ps.FeedbackMode {
			return
		}
	}
}

func permutationsWithRepetitionDescend(input []int, data []int, last int, index int, output *[][]int) {
	for i := 0; i < len(input); i++ {
		data[index] = input[i]
		if index == last {
			*output = append(*output, append([]int(nil), data...))
		} else {
			permutationsWithRepetitionDescend(input, data, last, index+1, output)
		}
	}
}

func permutationsWithRepetition(input []int) [][]int {
	length := len(input)
	data := make([]int, length)
	sort.Ints(input[:])
	output := make([][]int, 0)
	permutationsWithRepetitionDescend(input, data, length-1, 0, &output)
	return output
}

type ProgramState struct {
	AmplifierCode string
	Input         int
	Phase         int
	Memory        int
	Intcodes      []int
	CurPos        int
	InputCount    int
	FeedbackMode  bool
}

func run(program string, phases []int) int {

	intcodes := generate_intcodes(program)

	phase_permutations := permutationsWithRepetition(phases)

	// I could've just not had repetitions, but that's the implementation I
	// . foolishly made, and now I need to make filter for sets with unique
	// . sets.
	unique_phase_permutations := make([][]int, 0)
	for _, phases := range phase_permutations {
		seen := make(map[int]bool)
		for _, phase := range phases {
			if _, ok := seen[phase]; ok {
				break
			} else {
				seen[phase] = true
			}
		}
		if len(seen) == 5 {
			unique_phase_permutations = append(unique_phase_permutations, phases)
		}
	}

	max := 0
	stateMap := make(map[string]*ProgramState)
	for _, phases := range unique_phase_permutations {
		for i, letter := range []string{"a", "b", "c", "d", "e"} {
			new_intcodes := make([]int, len(intcodes))
			copy(new_intcodes, intcodes)
			stateMap[letter] = &ProgramState{
				AmplifierCode: letter,
				Phase:         phases[i],
				Intcodes:      new_intcodes,
				CurPos:        0,
				InputCount:    1,
				FeedbackMode:  false,
			}
		}

		// Check if part one or part two
		var val int
		if phases[0] < 5 {

			// Start with input 0
			val = 0
			for _, letter := range []string{"a", "b", "c", "d", "e"} {
				ps := stateMap[letter]
				ps.Input = val
				run_program(ps)
				val = ps.Memory
			}

			// if output here is greater than max, then set as new max
			if val > max {
				max = val
			}

		} else {

			// Set FeedbackMode true on all PStates
			for _, ps := range stateMap {
				ps.FeedbackMode = true
			}

			// Set run to 0 so we only use phase the first time
			run := 0
			val = 0
			for {
				for _, letter := range []string{"a", "b", "c", "d", "e"} {
					ps := stateMap[letter]
					if run == 0 {
						ps.InputCount = 1
					} else {
						ps.InputCount = 2
					}
					ps.Input = val
					run_program(ps)
					val = ps.Memory
				}
				run++

				// Since we only stop when one of the machines
				// . hits state of 99, we need to cycle through
				// . and check each one.
				var done bool
				for _, ps := range stateMap {
					if ps.Intcodes[ps.CurPos] == 99 {
						done = true
					}
				}
				if done {
					break
				}
			}

			// Once run is complete, check if the output from amplifier E
			// . is greater than previous max.
			if stateMap["e"].Memory > max {
				max = stateMap["e"].Memory
			}
		}
	}
	return max
}
