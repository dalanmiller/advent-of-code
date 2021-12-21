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
			// {X: 0, Y: 0},
			{X: 1, Y: 0},
			// {X: 2, Y: 0},
			{X: 0, Y: 1},

			{X: 2, Y: 1},
			// {X: 0, Y: 2},
			{X: 1, Y: 2},
			// {X: 2, Y: 2},
		}},
	}

	for _, test := range tests {
		result := adjacentCoords(test.i, test.j, test.width, test.height)
		if !reflect.DeepEqual(test.expected, result) {
			t.Fatalf("Result %v != expected %v", result, test.expected)
		}
	}
}

func TestExamplesFifteenOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`119
919
911`, 4},
		{`999
999
999`, 36},
		{`1163751742
1381373672
2136511328`, 33},
		{`1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`, 40},
	}

	for _, test := range tests {
		result := run(test.test, 1)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestFifteenOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 745},
	}

	for _, test := range tests {
		result := run(test.test, 1)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

// func TestExamplesFifteenTwo(t *testing.T) {
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

// func TestFifteenTwo(t *testing.T) {
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
