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
  #########`

var test2 = `#############
#AA.......A.#
###.#.#.#.###
  #D#.#.#D#
  #########`

var test3 = `#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########`

var test4 = `#############
#C.........C#
###A#B#.#D###
  #A#B#.#D#
  #########`

var test5 = `#############
#.....D.....#
###A#B#C#.###
  #A#B#C#.#
  #########`

var test6 = `#############
#...........#
###A#B#C#.###
  #A#B#C#D#
  #########`

var test7 = `#############
#C.........C#
###A#.#B#D###
  #A#.#B#D#
  #########`

var test8 = `#############
#C.........C#
###B#.#D#A###
  #D#.#B#A#
  #########`

var test9 = `#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########`

var test10 = `#############
#CD.....A.CD#
###.#.#.#.###
  #B#.#B#A#
  #########`

var test11 = `#############
#...........#
###B#B#D#D###
  #C#A#A#C#
  #########`

// A: 6, 7
// B: 2 +

var test12 = `#############
#...........#
###B#A#.#.###
  #A#B#.#.#
  #########`

// A: 2 + 4 = 6
// B: 4 = 40

var test13 = `#############
#D..........#
###.#.#.#C###
  #.#.#.#.#
  #########`

var test14 = `#############
#.A.A.......#
###.#.#.#C###
  #B#.#.#.#
  #########`

func TestBurrowTypeCompleteAndOccupantsTest1(t *testing.T) {
	amphipods := parseInput(test1)

	w := World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
	}

	complete, occupancy := w.burrowTypeCompleteAndOccupants(AMBER)

	if complete == true {
		log.Fatalf("Test1 Burrow should not be complete")
	}

	if occupancy[0] == true {
		log.Fatalf("First spot should not be occupied")
	}

	if occupancy[1] == false {
		log.Fatalf("Second spot should be occupied")
	}
}

func TestPossibleMovesTest(t *testing.T) {
	tests := []struct {
		test          string
		amphipodIndex int
		expectedMoves int
	}{
		{test1, 0, 1},  // 0
		{test2, 3, 1},  // 1
		{test3, 1, 0},  // 2
		{test3, 2, 0},  // 3
		{test3, 3, 0},  // 4
		{test3, 4, 0},  // 5
		{test7, 5, 0},  // 6
		{test7, 0, 0},  // 7
		{test7, 1, 0},  // 8
		{test8, 0, 0},  // 9
		{test8, 1, 0},  // 10
		{test8, 2, 1},  // 11
		{test13, 0, 0}, // 12
		{test14, 2, 0}, // 13
	}

	for i, test := range tests {
		amphipods := parseInput(test.test)

		w := World{
			Amphipods:    amphipods,
			BurrowHeight: 2,
		}

		moves := w.possibleMoves(amphipods[test.amphipodIndex])

		if len(moves) != test.expectedMoves {
			log.Fatalf("PossibleMovesTest %d incorrect number of moves, calculated %d, expected %d", i, len(moves), test.expectedMoves)
		}
	}

}

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

	// for i, occupied := range occupancy {
	// 	log.Printf("%d: %t", i, occupied)
	// }

	if occupancy[0] == true {
		log.Fatalf("First spot should not be occupied")
	}

	if occupancy[1] == false {
		log.Fatalf("Second spot should be occupied")
	}
}

func TestPossibleMovesTest5(t *testing.T) {
	amphipods := parseInput(test5)

	w := World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
	}

	moves := w.possibleMoves(amphipods[0])

	if len(moves) != 1 {
		log.Fatalf("Incorrect number of moves, calculated %d, expected 1", len(moves))
	}

}

func TestPossibleMovesTestAllNone(t *testing.T) {

	tests := []struct {
		test string
	}{
		{test6},
		{test9},
	}

	for _, test := range tests {
		amphipods := parseInput(test.test)

		w := World{
			Amphipods:    amphipods,
			BurrowHeight: 2,
		}

		for i := range amphipods {
			moves := w.possibleMoves(amphipods[i])

			if amphipods[i].Type != "" && len(moves) != 0 {
				log.Fatalf("Incorrect number of moves, calculated %d, expected 0", len(moves))
			}
		}
	}

}

func TestCostOfMoveTest2(t *testing.T) {
	amphipods := parseInput(test2)

	d1 := amphipods[3]
	destination := Position{9, 2}
	cost := d1.costToMove(destination)

	if cost != 9000 {
		log.Fatalf("Expected cost of 9000, got %d", cost)
	}

	a3 := amphipods[2]
	destination = Position{3, 2}
	cost = a3.costToMove(destination)
	if cost != 8 {
		log.Fatalf("Expected cost of 8, got %d", cost)
	}

}

func TestBurrowComplete(t *testing.T) {
	// Test an obviously incomplete arrangement
	amphipods := parseInput(test2)
	w := World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
	}

	if w.burrowsComplete() {
		log.Fatalf("Burrow returned complete when that is not true")
	}

	// Test a complete burrow
	amphipods = parseInput(test3)

	w = World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
	}

	if w.burrowsComplete() != true {
		log.Fatalf("Burrow complete, but received not complete")
	}

}

func TestAStar(t *testing.T) {

	tests := []struct {
		test         string
		BurrowHeight int
		expected     int
	}{
		{
			test4, 2, 1300,
		}, {
			test7, 2, 1400,
		}, {
			test8, 2, 29418,
		}, {
			test10, 2, 13326,
		}, {
			test9, 2, -1,
		},
		{
			test12, 2, 164,
		},
	}

	for i, test := range tests {
		amphipods := parseInput(test.test)

		w := World{
			Amphipods:    amphipods,
			BurrowHeight: int8(test.BurrowHeight),
		}

		cost := aStarRestart(&w)

		if cost != test.expected {
			log.Fatalf("Test %d failed, total cost was %d, expected %d", i, cost, test.expected)
		}
	}
}

func TestTwentyThreeOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		// Tried 6363
		// Tried 8363
		// Tried 20263
		{string(file), 20263},
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
