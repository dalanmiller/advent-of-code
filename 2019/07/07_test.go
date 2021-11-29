package main

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		Permutation    []int
		ExpectedLength int
		ExpectedFirst  []int
	}{
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2}, 4, []int{1, 1}},
		{[]int{1, 2, 3}, 27, []int{1, 1, 1}},
		{[]int{1, 2, 3, 4}, 256, []int{1, 1, 1, 1}},
		{[]int{1, 2, 3, 4, 5}, 3125, []int{1, 1, 1, 1, 1}},
	}

	for _, test := range tests {
		result := permutationsWithRepetition(test.Permutation)
		if test.ExpectedLength != len(result) {
			t.Fatalf("Result length %d != expected %d", len(result), test.ExpectedLength)
		}

		if !reflect.DeepEqual(test.ExpectedFirst, result[0]) {
			t.Fatalf("Result first %v != expected %v", result[0], test.ExpectedFirst)
		}
	}
}

func TestIntcodesGenerator(t *testing.T) {
	tests := []struct {
		program  string
		expected []int
	}{
		{"1,1,1", []int{1, 1, 1}},
		{"3,2,1", []int{3, 2, 1}},
		{"1003,2,1", []int{1003, 2, 1}},
	}

	for _, test := range tests {
		result := generate_intcodes(test.program)

		if len(result) != len(test.expected) {
			log.Fatalf("Result length %d != expected length %d", len(result), len(test.expected))
		}
		for i := range result {
			if result[i] != test.expected[i] {
				log.Fatalf("Contents are not the same, at %d, result value %d != expected %d", i, result[i], test.expected[i])
			}
		}
	}
}

func TestParseInstructions(t *testing.T) {
	tests := []struct {
		// Intcode            int
		ExpectedOperation   Operation
		ExpectedParamLength int
		ExpectedParams      []int
		IntCodes            []int
	}{
		{ADD, 3, []int{1, 1, 0}, []int{11101, 1, 1, 0}},
		{MULT, 3, []int{1, 1, 0}, []int{11102, 1, 1, 0}},
		{INPUT, 1, []int{2}, []int{103, 2, 0}},
		{OUTPUT, 1, []int{0}, []int{104, 0, 0}},
		{JUMPTRUE, 2, []int{0, 0}, []int{1105, 0, 0}},
		{JUMPFALSE, 2, []int{0, 0}, []int{1106, 0, 0}},
		{LESSTHAN, 3, []int{0, 0, 1}, []int{11107, 0, 0, 1}},
		{EQUALS, 3, []int{0, 0, 3}, []int{11108, 0, 0, 3}},
	}

	for _, test := range tests {
		operation, parameters := parse_instruction(test.IntCodes[0], 0, test.IntCodes)

		if test.ExpectedOperation != operation {
			t.Fatalf("Operation %v != expected %v", operation, test.ExpectedOperation)
		}

		if test.ExpectedParamLength != len(parameters) {
			t.Fatalf("Mode length %d != expected %d, %v", len(parameters), test.ExpectedParamLength, parameters)
		}

		if !reflect.DeepEqual(test.ExpectedParams, parameters) {
			t.Fatalf("Modes was %v != expected %v", parameters, test.ExpectedParams)
		}
	}
}

func TestExamplesSevenOne(t *testing.T) {
	tests := []struct {
		Program  string
		Expected int
	}{
		{"3,11,3,12,1,11,12,13,4,13,99,0,0,0", 10},
		{"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", 43210},
		{"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", 54321},
		{"3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,00", 65210},
	}

	for _, test := range tests {
		result := run(test.Program, []int{0, 1, 2, 3, 4})
		if test.Expected != result {
			t.Fatalf("Result % d != expected % d", result, test.Expected)
		}
	}
}

func TestRunProgram(t *testing.T) {
	tests := []struct {
		State    *ProgramState
		Expected int
	}{
		{
			State: &ProgramState{
				AmplifierCode: "1",
				Intcodes:      []int{3, 11, 3, 12, 1, 11, 12, 13, 4, 13, 99, 0, 0, 0},
				CurPos:        0,
				Input:         0,
				Memory:        0,
				InputCount:    1,
				FeedbackMode:  false,
			},
			Expected: 0,
		}, {
			State: &ProgramState{
				AmplifierCode: "2",
				Intcodes:      []int{3, 11, 3, 12, 1, 11, 12, 13, 4, 13, 99, 0, 0, 0},
				CurPos:        0,
				Input:         1,
				Memory:        0,
				InputCount:    1,
				FeedbackMode:  false,
			},
			Expected: 1,
		}, {
			State: &ProgramState{
				AmplifierCode: "3",
				Intcodes:      []int{3, 11, 3, 12, 1, 11, 12, 13, 4, 13, 99, 0, 0, 0},
				CurPos:        0,
				Input:         1,
				Memory:        1,
				Phase:         1,
				InputCount:    1,
				FeedbackMode:  false,
			},
			Expected: 2,
		}, {
			State: &ProgramState{
				AmplifierCode: "4",
				Intcodes:      []int{3, 11, 3, 12, 1, 11, 12, 13, 4, 13, 99, 0, 0, 0},
				CurPos:        0,
				Input:         1,
				Memory:        1,
				Phase:         1,
				InputCount:    1,
				FeedbackMode:  false,
			},
			Expected: 2,
		},
		{
			State: &ProgramState{
				AmplifierCode: "5.true",
				Intcodes:      []int{1105, 1, 4, 99, 3, 12, 4, 12, 99, 0, 0, 0, 0},
				CurPos:        0,
				Input:         0,
				Memory:        1,
				Phase:         1,
				InputCount:    1,
				FeedbackMode:  false,
			},
			Expected: 1,
		},
		{
			State: &ProgramState{
				AmplifierCode: "5.false",
				Intcodes:      []int{1105, 0, 4, 99, 3, 12, 4, 12, 99, 0, 0, 0, 0},
				CurPos:        0,
				Input:         1,
				Memory:        0,
				Phase:         1,
				InputCount:    1,
				FeedbackMode:  false,
			},
			Expected: 0,
		},
		{
			State: &ProgramState{
				AmplifierCode: "6.true",
				Intcodes:      []int{1106, 0, 4, 99, 3, 12, 4, 12, 99, 0, 0, 0, 0},
				CurPos:        0,
				Input:         1,
				Memory:        0,
				Phase:         1,
				InputCount:    1,
				FeedbackMode:  false,
			},
			Expected: 1,
		},
		{
			State: &ProgramState{
				AmplifierCode: "7.true",
				Intcodes:      []int{1107, 1, 2, 12, 4, 12, 99, 0, 00, 0, 0, 0, 0},
				CurPos:        0,
				Input:         1,
				Memory:        0,
				Phase:         1,
				InputCount:    1,
				FeedbackMode:  false,
			},
			Expected: 1,
		},
		{
			State: &ProgramState{
				AmplifierCode: "7.false",
				Intcodes:      []int{1107, 3, 2, 12, 4, 12, 99, 0, 00, 0, 0, 0, 0},
				CurPos:        0,
				Input:         1,
				Memory:        0,
				Phase:         1,
				InputCount:    1,
				FeedbackMode:  false,
			},
			Expected: 0,
		},
		{
			State: &ProgramState{
				AmplifierCode: "8.true",
				Intcodes:      []int{1108, 2, 2, 12, 4, 12, 99, 0, 0, 0, 0, 0, 0},
				CurPos:        0,
				Input:         1,
				Memory:        0,
				Phase:         1,
				InputCount:    1,
				FeedbackMode:  false,
			},
			Expected: 1,
		},
		{
			State: &ProgramState{
				AmplifierCode: "8.false",
				Intcodes:      []int{1108, 1, 2, 12, 4, 12, 99, 0, 0, 0, 0, 0, 0},
				CurPos:        0,
				Input:         1,
				Memory:        0,
				Phase:         1,
				InputCount:    1,
				FeedbackMode:  false,
			},
			Expected: 0,
		},
	}

	for _, test := range tests {
		run_program(test.State)
		if test.State.Memory != test.Expected {
			log.Fatalf("State %s, Memory %d != expected %d", test.State.AmplifierCode, test.State.Memory, test.Expected)
		}
	}
}

func TestRedditRunProgram(t *testing.T) {
	tests := []struct {
		State          *ProgramState
		ExpectedOutput int
		ExpectedCurPos int
	}{
		{
			State: &ProgramState{
				AmplifierCode: "0",
				Intcodes:      []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
				Phase:         9,
				CurPos:        0,
				Input:         0,
				Memory:        0,
				InputCount:    1,
				FeedbackMode:  true,
			},
			ExpectedOutput: 5,
			ExpectedCurPos: 18,
		},
		{
			State: &ProgramState{
				AmplifierCode: "1",
				Intcodes:      []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
				Phase:         8,
				CurPos:        0,
				Input:         5,
				Memory:        0,
				InputCount:    1,
				FeedbackMode:  true,
			},
			ExpectedOutput: 14,
			ExpectedCurPos: 18,
		}, {
			State: &ProgramState{
				AmplifierCode: "2",
				Intcodes:      []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
				Phase:         7,
				CurPos:        0,
				Input:         14,
				Memory:        0,
				InputCount:    1,
				FeedbackMode:  true,
			},
			ExpectedOutput: 31,
			ExpectedCurPos: 18,
		}, {
			State: &ProgramState{
				AmplifierCode: "3",
				Intcodes:      []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
				Phase:         6,
				CurPos:        0,
				Input:         31,
				Memory:        0,
				InputCount:    1,
				FeedbackMode:  true,
			},
			ExpectedOutput: 64,
			ExpectedCurPos: 18,
		}, {
			State: &ProgramState{
				AmplifierCode: "4",
				Intcodes:      []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
				Phase:         5,
				CurPos:        0,
				Input:         64,
				Memory:        0,
				InputCount:    1,
				FeedbackMode:  true,
			},
			ExpectedOutput: 129,
			ExpectedCurPos: 18,
		},
		{
			State: &ProgramState{
				AmplifierCode: "0",
				Intcodes:      []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 5, 5, 5},
				Phase:         9,
				CurPos:        18,
				Input:         129,
				Memory:        0,
				InputCount:    2,
				FeedbackMode:  true,
			},
			ExpectedOutput: 263,
			ExpectedCurPos: 18,
		},
	}

	for _, test := range tests {
		run_program(test.State)
		if test.State.Memory != test.ExpectedOutput {
			log.Fatalf("State %s, Memory %d != expected %d", test.State.AmplifierCode, test.State.Memory, test.ExpectedOutput)
		}

		if test.State.CurPos != test.ExpectedCurPos {
			log.Fatalf("State %s, CurPos %d != expected %d", test.State.AmplifierCode, test.State.CurPos, test.ExpectedCurPos)
		}
	}
}

func TestSevenOne(t *testing.T) {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("could not read input, yikes")
	}

	tests := []struct {
		Program  string
		Expected int
	}{
		{string(file), 368584},
	}

	for _, test := range tests {
		result := run(test.Program, []int{0, 1, 2, 3, 4})
		if test.Expected != result {
			t.Fatalf("Result % d != expected % d", result, test.Expected)
		}
	}
}

func BenchmarkSevenOne(b *testing.B) {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("could not read input, yikes")
	}

	tests := []struct {
		Program  string
		Expected int
	}{
		{string(file), 368584},
	}

	for _, test := range tests {
		for i := 0; i < b.N; i++ {
			result := run(test.Program, []int{0, 1, 2, 3, 4})
			if test.Expected != result {
				b.Fatalf("Result % d != expected % d", result, test.Expected)
			}
		}
	}
}

func TestExamplesSevenTwo(t *testing.T) {
	tests := []struct {
		Program  string
		Expected int
	}{
		{"3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5", 139629729},
		{"3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10", 18216},
	}

	for _, test := range tests {
		result := run(test.Program, []int{5, 6, 7, 8, 9})
		if test.Expected != result {
			t.Fatalf("Result % d != expected % d", result, test.Expected)
		}
	}
}

func TestSevenTwo(t *testing.T) {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("could not read input, yikes")
	}

	tests := []struct {
		Program  string
		Expected int
	}{
		{string(file), 35993240},
	}

	for _, test := range tests {
		result := run(test.Program, []int{5, 6, 7, 8, 9})
		if test.Expected != result {
			t.Fatalf("Result %d != expected %d", result, test.Expected)
		}
	}
}
