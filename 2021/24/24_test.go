package main

import (
	"log"
	"os"
	"testing"
)

func TestALU(t *testing.T) {
	tests := []struct {
		inputs       [14]int64
		instructions string
		expected     int64
	}{
		{[14]int64{1}, "inp z\nmul z -1", -1},
		{[14]int64{1, 3}, "inp z\ninp x\nmul z 3\neql z x", 1},
		{[14]int64{5}, "inp w\nadd z w\nmod z 2\ndiv w 2\nadd y w\nmod y 2\ndiv w 2\nadd x w\nmod x 2\ndiv w 2\nmod w 2", 1},
	}

	for _, test := range tests {
		instructions := parseInput(test.instructions)
		z := ALU(test.inputs, instructions)
		if test.expected != z {
			t.Fatalf("Result % d != expected % d", z, test.expected)
		}
	}
}

func TestTwentyFourOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

// func TestExamplesOneTwo(t *testing.T) {
// 	tests := []struct {
// 		test     string
// 		expected int
// 	}{
// 		{"", 0},
// 	}

// 	for _, test := range tests {
// 		result := run(test.test)
// 		if test.expected != result {
// 			t.Fatalf("Result % d != expected % d", result, test.expected)
// 		}
// 	}
// }

// func TestOneTwo(t *testing.T) {
// 	file, err := os.ReadFile("./input")
// 	if err != nil {
// 		log.Fatalf("could not read file")
// 	}

// 	tests := []struct {
// 		test     string
// 		expected int
// 	}{
// 		{string(file), 0},
// 	}

// 	for _, test := range tests {
// 		result := run(test.test)
// 		if result[0] != test.expected {
// 			t.Fatalf("Result % d != expected % d", result, test.expected)
// 		}
// 	}
// }
