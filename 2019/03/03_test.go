package main

import (
	"log"
	"os"
	"testing"
)

func TestIntersection(t *testing.T) {
	a := Path{
		Lines: []Line{
			{
				Begin:     Coordinate{X: 0, Y: 0},
				End:       Coordinate{X: 0, Y: 2},
				Direction: UP,
			},
			{
				Begin:     Coordinate{X: 0, Y: 2},
				End:       Coordinate{X: 2, Y: 2},
				Direction: RIGHT,
			},
			{
				Begin:     Coordinate{X: 2, Y: 2},
				End:       Coordinate{X: 2, Y: -1},
				Direction: DOWN,
			},
		},
	}
	b := Path{
		Lines: []Line{
			{
				Begin:     Coordinate{X: 1, Y: 1},
				End:       Coordinate{X: -1, Y: 1},
				Direction: LEFT,
			},
			{
				Begin:     Coordinate{X: -1, Y: 1},
				End:       Coordinate{X: -1, Y: 0},
				Direction: DOWN,
			},
			{
				Begin:     Coordinate{X: -1, Y: 0},
				End:       Coordinate{X: 3, Y: 0},
				Direction: RIGHT,
			},
		},
	}

	full_coordinates := a.full_coordinates()
	if len(full_coordinates) != 8 {
		log.Fatalf("Path A coordinates != 8, got %d", len(full_coordinates))
	}

	if len(a.intersect(&b)) != 2 {
		log.Fatalf("Did not find an intersection on Paths")
	}

	found_intersects := a.intersect(&b)
	expected_coordinate1 := Coordinate{X: 0, Y: 1}
	expected_coordinate2 := Coordinate{X: 2, Y: 0}

	if _, ok := found_intersects[expected_coordinate1]; !ok {
		log.Fatalf("Incorrect coordinate intersection 1, got %d", found_intersects[expected_coordinate1])
	}

	if _, ok := found_intersects[expected_coordinate2]; !ok {
		log.Fatalf("Incorrect coordinate intersection 2, got %d", found_intersects[expected_coordinate2])
	}

}

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		coordinates Coordinate
		expected    int
	}{
		{Coordinate{X: 1, Y: 1}, 2},
		{Coordinate{X: -1, Y: -1}, 2},
		{Coordinate{X: -5, Y: 5}, 10},
		{Coordinate{X: 2, Y: -2}, 4},
	}

	for _, test := range tests {
		result := test.coordinates.manhattan()
		if result != test.expected {
			log.Fatalf("Incorrect distance, got %d, expected %d", result, test.expected)
		}
	}
}

func TestExamplesThreeOne(t *testing.T) {
	tests := []struct {
		wires    string
		expected int
	}{
		{"R8,U5,L5,D3\nU7,R6,D4,L4", 6},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}

	for _, test := range tests {
		result, _ := run(test.wires)
		if result != test.expected {
			log.Fatalf("Result %d is not expected %d", result, test.expected)
		}
	}

}

func TestThreeOne(t *testing.T) {

	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		wires    string
		expected int
	}{
		{string(file), 721},
	}

	for _, test := range tests {
		result, _ := run(test.wires)
		if result != test.expected {
			log.Fatalf("Fail: got %d, expected %d", result, test.expected)
		}
	}
}

func TestExamplesThreeTwo(t *testing.T) {
	tests := []struct {
		wires    string
		expected int
	}{
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R834", 610},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410},
	}

	for _, test := range tests {
		_, steps := run(test.wires)
		if steps != test.expected {
			log.Fatalf("Result %d is not expected %d", steps, test.expected)
		}
	}

}

func TestThreeTwo(t *testing.T) {

	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		wires    string
		expected int
	}{
		{string(file), 7388},
	}

	for _, test := range tests {
		_, steps := run(test.wires)
		if steps != test.expected {
			log.Fatalf("Fail: got %d, expected %d", steps, test.expected)
		}
	}
}
