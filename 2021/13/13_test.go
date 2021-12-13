package main

import (
	"math"
	"os"
	"testing"
)

func TestFoldLeft(t *testing.T) {
	tests := []struct {
		test     Paper
		x        int
		expected int
	}{
		{Paper{
			Height: 5,
			Width:  5,
			Dots: []Coord{

				{X: 1, Y: 0},
				{X: 1, Y: 2},
				{X: 1, Y: 4},
				{X: 3, Y: 0},
				{X: 3, Y: 1},
				{X: 3, Y: 2},
				{X: 3, Y: 3},
				{X: 3, Y: 4},
			},
		}, 2, 5},
	}

	for _, test := range tests {
		test.test.foldLeft(test.x)
		result := len(test.test.Dots)
		if result != test.expected {
			t.Fatalf("Incorrect dots, got %d, expected %d", result, test.expected)
		}
	}
}
func TestExamplesThirteenOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7`, 17},
	}

	for _, test := range tests {
		result := run(test.test, 1)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestThirteenOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		t.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 802},
	}

	for _, test := range tests {
		result := run(test.test, 1)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestThirteenTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		t.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 120},
	}

	// ###..#..#.#..#.##.#.####..##..#..#.###..
	// #..#.#.#..#..#.#.......#.#..#.#..#.#..#.
	// #..#.##...####.###.......#....#..#.###..
	// ###..#.#..#..#.#.....#...#.##.#..#.#..#.
	// #.#..#.#..#..#.#....#....#..#.#..#.#..#.
	// #..#.#..#.#..#.#....####..###..##..###..

	for _, test := range tests {
		result := run(test.test, math.MaxInt)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
