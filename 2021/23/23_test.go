package main

import (
	"log"
	"os"
	"testing"
)

var test1 = `#############
#...........#
###.#.#.#.###
  #D#.#.#.#
  ########`

func TestBurrowTypeCompleteAndOccupantsTest1(t *testing.T) {
	amphipods := parseInput(test1)
	log.Print(amphipods)

	w := World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
	}

	complete, occupancy := w.burrowTypeCompleteAndOccupants(AMBER)

	if complete == true {
		log.Fatalf("Test1 Burrow should not be complete")
	}

	for i, occupied := range occupancy {
		log.Printf("%d: %t", i, occupied)
	}

	if occupancy[0] == true {
		log.Fatalf("First spot should not be occupied")
	}

	if occupancy[1] == false {
		log.Fatalf("Second spot should be occupied")
	}
}

func TestPossibleMovesTest1(t *testing.T) {
	amphipods := parseInput(test1)

	if len(amphipods) != 1 {
		log.Print(amphipods)
		log.Fatalf("Got too many amphipods, got %d, wanted 1", len(amphipods))
	}

	w := World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
	}

	moves := w.possibleMoves(amphipods[0])

	for _, position := range moves {
		log.Printf("P - X:%d, Y:%d ", position.Column, position.Row)
	}

	if len(moves) != 8 {
		log.Fatalf("Incorrect number of moves, calculated %d, expected 8", len(moves))
	}

}

var test2 = `#############
#AA.......A.#
###.#.#.#.###
  #D#.#.#D#
  ########`

func TestBurrowTypeCompleteAndOccupantsTest2(t *testing.T) {
	amphipods := parseInput(test2)

	w := World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
	}

	complete, occupancy := w.burrowTypeCompleteAndOccupants(AMBER)

	if complete == true {
		log.Fatalf("Burrow should not be complete")
	}

	for i, occupied := range occupancy {
		log.Printf("%d: %t", i, occupied)
	}

	if occupancy[0] == true {
		log.Fatalf("First spot should not be occupied")
	}

	if occupancy[1] == false {
		log.Fatalf("Second spot should be occupied")
	}
}

func TestPossibleMovesTest2(t *testing.T) {
	amphipods := parseInput(test2)

	if len(amphipods) != 5 {
		log.Fatalf("Got too many amphipods, got %d, wanted 5", len(amphipods))
	}

	w := World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
	}

	moves := w.possibleMoves(amphipods[0])

	for _, position := range moves {
		log.Printf("P - X:%d, Y:%d ", position.Column, position.Row)
	}

	if len(moves) != 4 {
		log.Fatalf("Incorrect number of moves, calculated %d, expected 4", len(moves))
	}

}

func TestExamplesTwentyOneOneParseInput(t *testing.T) {

	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}
	result := parseInput(string(file))
	w := World{
		Amphipods:    result,
		BurrowHeight: 2,
	}

	w.print()
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
