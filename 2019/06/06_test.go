package main

import (
	"log"
	"os"
	"testing"
)

func TestExamplesSixOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{"COM)B", 1},
		{"COM)B\nB)C\nC)D", 6},
		{"COM)B\nB)C\nB)D\nB)E\nB)F\nB)G", 11},
		{"COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L", 42},
	}

	for _, test := range tests {
		result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSixOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 1822},
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
// 		test     int
// 		expected int
// 	}{
// 		{0, 0},
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
// 		test     int
// 		expected int
// 	}{
// 		{0, 0},
// 	}

// 	for _, test := range tests {
// 		result := run(test.test)
// 		if result[0] != test.expected {
// 			t.Fatalf("Result % d != expected % d", result, test.expected)
// 		}
// 	}
// }
