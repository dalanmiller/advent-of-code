package main

import (
	"container/heap"
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

var test15 = `#############
#C....B.B..C#
###A#.#.#D###
  #A#.#.#D#
  #########`

var example1 = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

var example1A = `#############
#...........#
###B#C#B#D###
  #.#D#C#.#
  #########`

var example1B = `#############
#...........#
###.#C#.#D###
  #A#D#C#A#
  #########`

var example1C = `#############
#...........#
###B#.#B#D###
  #A#D#.#A#
  #########`

var stuck1 = `#############
#D.........B#
###.#B#.#D###
  #A#A#C#C#
  #########
`

func TestBurrowTypeCompleteAndOccupantsTest1(t *testing.T) {
	amphipods := parseInput(test1)

	w := World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
	}

	complete, occupancy := w.burrowTypeCompleteAndOccupants(AMBER)

	if complete == true {
		t.Fatalf("Test1 Burrow should not be complete")
	}

	if occupancy[0] != "" {
		t.Fatalf("First spot should not be occupied")
	}

	if occupancy[1] == "" {
		t.Fatalf("Second spot should be occupied")
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
		{stuck1, 3, 5}, // 14
		{stuck1, 6, 5}, // 15
	}

	for i, test := range tests {
		amphipods := parseInput(test.test)

		w := World{
			Amphipods:    amphipods,
			BurrowHeight: 2,
		}

		moves := w.possibleMoves(amphipods[test.amphipodIndex])

		if len(moves) != test.expectedMoves {
			t.Fatalf("PossibleMovesTest %d incorrect number of moves, calculated %d, expected %d", i, len(moves), test.expectedMoves)
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
		t.Fatalf("Burrow should not be complete")
	}

	// for i, occupied := range occupancy {
	// 	log.Printf("%d: %t", i, occupied)
	// }

	if occupancy[0] != "" {
		t.Fatalf("First spot should not be occupied")
	}

	if occupancy[1] == "" {
		t.Fatalf("Second spot should be occupied")
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
		t.Fatalf("Incorrect number of moves, calculated %d, expected 1", len(moves))
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
				t.Fatalf("Incorrect number of moves, calculated %d, expected 0", len(moves))
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
		t.Fatalf("Expected cost of 9000, got %d", cost)
	}

	a3 := amphipods[2]
	destination = Position{3, 2}
	cost = a3.costToMove(destination)
	if cost != 8 {
		t.Fatalf("Expected cost of 8, got %d", cost)
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
		t.Fatalf("Burrow returned complete when that is not true")
	}

	// Test a complete burrow
	amphipods = parseInput(test3)

	w = World{
		Amphipods:    amphipods,
		BurrowHeight: 2,
	}

	if w.burrowsComplete() != true {
		t.Fatalf("Burrow complete, but received not complete")
	}

}

func TestLCQ(t *testing.T) {
	worlds := []*World{
		{
			CurrentCost:    110,
			AdditionalCost: 1000,
			AmphipodsLeft:  2,
		},
		{
			CurrentCost:    100,
			AdditionalCost: 100,
			AmphipodsLeft:  5,
		},
		{
			CurrentCost:    91,
			AdditionalCost: 1,
			AmphipodsLeft:  2,
		},
		{
			CurrentCost:    90,
			AdditionalCost: 10,
			AmphipodsLeft:  1,
		},
	}

	lcq := LowCostQueue{
		worlds:   []*World{},
		contains: map[*World]bool{},
		fScore:   map[*World]uint16{},
	}

	for i, _ := range worlds {
		lcq.fScore[worlds[i]] = worlds[i].h()
	}

	heap.Init(&lcq)

	for i, _ := range worlds {
		heap.Push(&lcq, worlds[i])
	}

	for i, _ := range worlds {
		if !lcq.Contains(worlds[i]) {
			log.Fatal("LCQ does not contain world after pushing")
		}
	}

	lowestWorld := heap.Pop(&lcq).(*World)
	if lowestWorld.AdditionalCost != 10 || lowestWorld.AmphipodsLeft != 1 || lcq.fScore[lowestWorld] != 90 {
		t.Fatalf("Did not Pop() lowest addl cost world first, got %d %d %d", lowestWorld.AdditionalCost, lowestWorld.AmphipodsLeft, lcq.fScore[lowestWorld])
	}

}

// 2022-06-14
// BenchmarkAStar-10    	       1	6_748_991_875 ns/op	2937985560 B/op	25544750 allocs/op
// PASS
// ok  	23	7.032s

// 2022-06-14
// BenchmarkAStar-10    	       1	7_276_955_875 ns/op	3109453136 B/op	25621226 allocs/op
// PASS
// ok  	23	7.569s

// 2022-06-14
// BenchmarkAStar-10    	       1	7_161_838_375 ns/op	3109385264 B/op	25620755 allocs/op
// PASS
// ok  	23	7.272s

// 2022-06-15
// BenchmarkAStar-10    	       1	2302544208 ns/op	636207480 B/op	 5540297 allocs/op
// PASS
// ok  	23	2.578s

// 2022-06-16
// BenchmarkAStar-10    	       1	3764858500 ns/op	8_681_082_000 B/op	 5423410 allocs/op
// PASS
// ok  	23	4.317s

// 2022-06-16
// BenchmarkAStar-10    	       1	3_693_698_250 ns/op	8_652_258_024 B/op	 5415201 allocs/op
// PASS
// ok  	23	4.172s

// 2022-06-16
// BenchmarkAStar-10    	       1	3_346_304_250 ns/op	805_341_080 B/op	 5745537 allocs/op
// PASS
// ok  	23	3.672s

// 2022-06-16

func BenchmarkAStar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		amphipods := parseInput(example1)

		w := World{
			Amphipods:      amphipods,
			CurrentCost:    0,
			AdditionalCost: 0,
			BurrowHeight:   int8(2),
			AmphipodsLeft:  int8(4 * 2),
		}

		b.ResetTimer()
		aStarRestart(w)
	}
}

// ok  	23	9.799s
// ok  	23	3.741s
func TestAStar(t *testing.T) {

	tests := []struct {
		test         string
		BurrowHeight int8
		expected     uint16
	}{
		{
			test4, 2, 1300,
		},
		{
			test7, 2, 1400,
		},
		{
			test15, 2, 1370,
		},
		{
			test8, 2, 15422,
		},
		{
			test10, 2, 13326,
		},
		{
			test9, 2, 0,
		},
		{
			test12, 2, 46,
		},
		{
			example1A, 2, 12510, // Should be less than 12521
		},
		{
			example1B, 2, 12409, // Should be less than 12521
		},
		{
			example1C, 2, 12101, // Should be less than 12521
		},
		{
			example1, 2, 12521,
		},
	}

	for i, test := range tests {
		amphipods := parseInput(test.test)

		w := World{
			Amphipods:      amphipods,
			CurrentCost:    0,
			BurrowHeight:   test.BurrowHeight,
			AdditionalCost: 0,
			AmphipodsLeft:  int8(4 * test.BurrowHeight),
		}

		cost := aStarRestart(w)

		if cost != test.expected {
			t.Fatalf("Test %d failed, total cost was %d, expected %d", i, cost, test.expected)
		}
	}
}

func TestTwentyThreeOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		t.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected uint16
	}{
		// Tried 6363
		// Tried 8363
		// Tried 20263
		// Tried 35115
		{string(file), 10607},
	}

	for _, test := range tests {
		result := run(test.test, 2)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

// func TestAStarPartTwo(t *testing.T) {

// 	examplePartTwo := `#############
// #...........#
// ###B#C#B#D###
//   #D#C#B#A#
//   #D#B#A#C#
//   #A#D#C#A#
//   #########`

// 	tests := []struct {
// 		test         string
// 		BurrowHeight int8
// 		expected     uint16
// 	}{
// 		{
// 			examplePartTwo, 4, 44169,
// 		},
// 	}

// 	for i, test := range tests {
// 		amphipods := parseInput(test.test)

// 		w := World{
// 			Amphipods:      amphipods,
// 			CurrentCost:    0,
// 			BurrowHeight:   test.BurrowHeight,
// 			AdditionalCost: 0,
// 			AmphipodsLeft:  int8(4 * test.BurrowHeight),
// 		}

// 		cost := aStarRestart(w)

// 		if cost != test.expected {
// 			t.Fatalf("Test %d failed, total cost was %d, expected %d", i, cost, test.expected)
// 		}
// 	}
// }

func TestTwentyThreeTwo(t *testing.T) {
	file, err := os.ReadFile("./input2")
	if err != nil {
		t.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected uint16
	}{
		{string(file), 59071},
	}

	for _, test := range tests {
		result := run(test.test, 4)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
