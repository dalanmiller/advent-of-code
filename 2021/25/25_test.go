package main

import (
	"os"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`

	rc, dc, _ := parseInput(input)

	if len(rc) != 23 {
		t.Fatalf("Incorrect number of right cucs, got %d", len(rc))
	}

	if len(dc) != 26 {
		t.Fatalf("Incorrect number of down cucs, got %d", len(dc))
	}
}

func TestExamplesTwentyFive(t *testing.T) {
	tests := []struct {
		test          string
		expectedMoves int
	}{
		{`........v.
>.......v.
........v.`, 8},
		{`........v>
........v.
........v.`, 9},
		{`..v
>>>
...
...`, 1},
	}

	for i, test := range tests {
		n := run(test.test)
		if n != test.expectedMoves {
			t.Fatalf("Test %d: did not match expected moves, got %d, expected %d", i, n, test.expectedMoves)
		}
	}
}

func TestGridPrinter(t *testing.T) {
	test := `..v
>>>
...`

	_, _, pm := parseInput(test)

	result := printGrid(3, 3, pm)
	if result != test {
		t.Fatalf("Grid did not match input, got: \n\n %s", result)
	}
}

func TestTwentyFiveExample(t *testing.T) {
	test := `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`

	result := run(test)
	if result != 58 {
		t.Fatalf("Example did not take 58 steps, got %d", result)
	}
}

func TestOneOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		t.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 380},
	}

	for _, test := range tests {
		result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

//func TestExamplesOneTwo(t *testing.T) {
//	tests := []struct {
//		test     string
//		expected int
//	}{
//		{"", 0},
//	}
//
//	for _, test := range tests {
//		result := run(test.test)
//		if test.expected != result {
//			t.Fatalf("Result % d != expected % d", result, test.expected)
//		}
//	}
//}
//
//func TestOneTwo(t *testing.T) {
//	file, err := os.ReadFile("./input")
//	if err != nil {
//		log.Fatalf("could not read file")
//	}
//
//	tests := []struct {
//		test     string
//		expected int
//	}{
//		{string(file), 0},
//	}
//
//	for _, test := range tests {
//		result := run(test.test)
//		if result[0] != test.expected {
//			t.Fatalf("Result % d != expected % d", result, test.expected)
//		}
//	}
//}
