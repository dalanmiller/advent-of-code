package main

import (
	"log"
	"os"
	"testing"
)

func TestExamplesFiveOne(t *testing.T) {
	tests := []struct {
		program  string
		input    int
		expected int
	}{
		{"3,0,4,0,99", 1, 1},
		{"1002,4,3,4,33", 0, 0},
		{"1101,50,50,0,4,0,99", 0, 100},
		// {"1101,100,-1,4,0", 0, 0},
	}

	for _, test := range tests {
		result, intcodes := run(test.input, test.program)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d | %d", result, test.expected, intcodes)
		}
	}
}

func TestFiveOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		program  string
		input    int
		expected int
	}{
		{string(file), 1, 9025675},
	}

	for _, test := range tests {
		result, intcodes := run(test.input, test.program)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d | %d", result, test.expected, intcodes)
		}
	}
}

func TestExamplesFiveTwo(t *testing.T) {
	tests := []struct {
		program  string
		input    int
		expected int
	}{
		{"1101,50,50,0,0108,100,0,0,4,0,99", 0, 1},
		{"1101,50,50,0,0107,101,0,0,4,0,99", 0, 0},
		{"1101,50,50,0,0107,99,0,0,4,0,99", 0, 1},
		{"3,3,1107,-1,8,3,4,3,99", 7, 1},
		{"3,3,1107,-1,8,3,4,3,99", 9, 0},
		{"3,3,1108,-1,8,3,4,3,99", 8, 1},
		{"3,3,1108,-1,8,3,4,3,99", 9, 0},
		{"3,9,7,9,10,9,4,9,99,-1,8", 7, 1},
		{"3,9,7,9,10,9,4,9,99,-1,8", 9, 0},
		{"3,9,8,9,10,9,4,9,99,-1,8", 8, 1},
		{"3,9,8,9,10,9,4,9,99,-1,8", 9, 0},

		// Jump tests
		{"5,7,8,99,104,0,99,1,4", 0, 0},
		{"6,7,8,99,104,0,99,0,4", 0, 0},
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 0, 0},
		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 0, 0},
		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 1, 1},
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 1, 1},

		{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", 8, 1000},
		{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", 9, 1001},
	}

	for _, test := range tests {
		result, intcodes := run(test.input, test.program)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d | %d", result, test.expected, intcodes)
		}
	}
}

func TestFiveTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		program  string
		input    int
		expected int
	}{
		{string(file), 5, 11981754},
	}

	for _, test := range tests {
		result, intcodes := run(test.input, test.program)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d | %d", result, test.expected, intcodes)
		}
	}
}
