package main

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestParseLines(t *testing.T) {
	tests := []struct {
		test     string
		expected []Line
	}{
		{`15,4 -> 13,4
1,10 -> 1,8
0,0 -> 2,2`,
			[]Line{
				{
					start: Coordinate{x: 15, y: 4},
					end:   Coordinate{x: 13, y: 4},
					coordinates: []Coordinate{
						{x: 15, y: 4},
						{x: 14, y: 4},
						{x: 13, y: 4},
					},
				},
				{
					start: Coordinate{x: 1, y: 10},
					end:   Coordinate{x: 1, y: 8},
					coordinates: []Coordinate{
						{x: 1, y: 10},
						{x: 1, y: 9},
						{x: 1, y: 8},
					},
				},
				{
					start: Coordinate{x: 0, y: 0},
					end:   Coordinate{x: 2, y: 2},
					coordinates: []Coordinate{
						{x: 0, y: 0},
						{x: 1, y: 1},
						{x: 2, y: 2},
					},
				},
			},
		},
	}

	for _, test := range tests {
		result := parse_lines(test.test, true)
		if !reflect.DeepEqual(result, test.expected) {
			log.Fatalf("Parsed line %v does not match expected %v", result, test.expected)
		}
	}
}

func TestExamplesFiveOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`, 5},
	}

	for _, test := range tests {
		result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestFiveOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 7318},
	}

	for _, test := range tests {
		result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesFiveTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`, 12},
	}

	for _, test := range tests {
		result := run_two(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestFiveTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 19939},
	}

	for _, test := range tests {
		result := run_two(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
