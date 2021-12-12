package main

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestAdjacentCoords(t *testing.T) {
	tests := []struct {
		i        int
		j        int
		width    int
		height   int
		expected []Coord
	}{
		{1, 1, 3, 3, []Coord{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 2, Y: 0},
			{X: 0, Y: 1},

			{X: 2, Y: 1},
			{X: 0, Y: 2},
			{X: 1, Y: 2},
			{X: 2, Y: 2},
		}},
	}

	for _, test := range tests {
		result := adjacentCoords(test.i, test.j, test.width, test.height)
		if !reflect.DeepEqual(test.expected, result) {
			t.Fatalf("Result %v != expected %v", result, test.expected)
		}
	}
}
func TestExamplesElevenOne(t *testing.T) {
	tests := []struct {
		test     string
		days     int
		expected int
	}{
		{`11111
19991
19191
19991
11111`, 2, 9},
		{`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`, 2, 35},
		{`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`, 10, 204},
		{`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`, 100, 1656},
	}

	for _, test := range tests {
		result := run(test.test, test.days)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestElevenOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		days     int
		expected int
	}{
		{string(file), 100, 1683},
	}

	for _, test := range tests {
		result := run(test.test, test.days)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestElevenTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		days     int
		expected int
	}{
		{string(file), 1000, 998},
	}

	for _, test := range tests {
		result := run(test.test, test.days)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
