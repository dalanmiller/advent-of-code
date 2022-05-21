package main

import (
	"log"
	"os"
	"testing"
)

func TestExamplesTwentyThreeOneParseInput(t *testing.T) {

	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}
	file_string := string(file)
	result := parseInput(file_string)
	w := World{
		Amphipods:    result,
		BurrowHeight: 2,
	}

	output := w.Print()
	if output != file_string {
		log.Fatalf("Parsed map does not match raw map\n Received\n%s\nExpected\n%s", output, file_string)
	}
}

func TestExamplesTwentyThreeOne(t *testing.T) {
	baseArrangement := `#############
#...........#
###.#.#.#.###
	#A#.#A#.#
	#########
`

	amphipods := parseInput(baseArrangement)
	world := World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
	}

	result := world.Solve()
	if result != 15 {
		log.Fatalf("Result does not equal 15, got %d instead", result)
	}
}

// func TestExamplesOneOne(t *testing.T) {
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

// func TestOneOne(t *testing.T) {
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
// 		if result != test.expected {
// 			t.Fatalf("Result % d != expected % d", result, test.expected)
// 		}
// 	}
// }

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
