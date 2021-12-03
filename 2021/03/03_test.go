package main

import (
	"log"
	"os"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		test     string
		expected [][]int
	}{
		{`00100`, [][]int{
			{0, 0, 1, 0, 0},
		}},
	}

	for _, test := range tests {
		result := parse_input(test.test)
		for i, line := range result {
			for j, n := range line {
				if n != test.expected[i][j] {
					t.Fatalf("Result % d != expected % d", result, test.expected)
				}
			}
		}
	}
}
func TestExamplesThreeOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`, 198},
	}

	for _, test := range tests {
		result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestThreeOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 3895776},
	}

	for _, test := range tests {
		result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesThreeTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`, 230},
	}

	for _, test := range tests {
		result := run_two(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestOneTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 7928162},
	}

	for _, test := range tests {
		result := run_two(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
