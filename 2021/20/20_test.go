package main

import (
	"log"
	"os"
	"testing"
)

func TestExamplesTwentyPartOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`, 35},
	}

	for _, test := range tests {
		result := run(test.test, 2)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestAdjacent(t *testing.T) {
	tests := []struct {
		test     Image
		a        int
		b        int
		expected int
	}{
		{
			Image{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1}},
			1, 1, 511,
		},
		{
			Image{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}},
			1, 1, 16,
		},
		{
			Image{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}},
			0, 0, 1,
		},
		{
			Image{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}},
			2, 2, 256,
		},
		{
			Image{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}},
			2, 0, 4,
		},
		{
			Image{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}},
			0, 2, 64,
		},
		{
			Image{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}},
			0, 1, 8,
		},
		{
			Image{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}},
			2, 1, 32,
		},
	}

	for _, test := range tests {
		result := test.test.adjacent(test.a, test.b)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestTwentyPartOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		// !5278
		{string(file), 5278},
	}

	for _, test := range tests {
		result := run(test.test, 2)
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
