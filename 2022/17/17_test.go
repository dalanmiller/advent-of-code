package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

const EXAMPLE = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func TestCanMoveVert(t *testing.T) {
	ch := chamber{}
	sh := shape{
		VERT, point{2, 5},
	}

	if ch.canMove(&sh, RIGHT) != true {
		t.Fatalf("FAILED to move VERT to the RIGHT when possible")
	}

	if ch.canMove(&sh, LEFT) != true {
		t.Fatalf("Expected true, got false when moving LEFT and blocked")
	}

	// Top
	ch = chamber{{3, 5}: true}
	if ch.canMove(&sh, RIGHT) != false {
		t.Fatalf("Expected false, got true when moving RIGHT and blocked at top")
	}

	// Top mid
	ch = chamber{{3, 4}: true}
	if ch.canMove(&sh, RIGHT) != false {
		t.Fatalf("Expected false, got true when moving RIGHT and blocked at top mid")
	}

	// Bot mid
	ch = chamber{{3, 3}: true}
	if ch.canMove(&sh, RIGHT) != false {
		t.Fatalf("Expected false, got true when moving RIGHT and blocked at bot mid")
	}

	// bot
	ch = chamber{{3, 2}: true}
	if ch.canMove(&sh, RIGHT) != false {
		t.Fatalf("Expected false, got true when moving RIGHT and blocked at bot")
	}

	// Test top
	ch = chamber{{1, 5}: true}

	if ch.canMove(&sh, LEFT) != false {
		t.Fatalf("Expected true, got false when moving LEFT and blocked")
	}

	// Test top mid
	ch = chamber{{1, 4}: true}

	if ch.canMove(&sh, LEFT) != false {
		t.Fatalf("Expected true, got false when moving LEFT and blocked")
	}

	// Test bot mid
	ch = chamber{{1, 3}: true}

	if ch.canMove(&sh, LEFT) != false {
		t.Fatalf("Expected true, got false when moving LEFT and blocked")
	}

	// Test bot
	ch = chamber{{1, 2}: true}

	if ch.canMove(&sh, LEFT) != false {
		t.Fatalf("Expected true, got false when moving LEFT and blocked")
	}
}

// ..#
// ..#
// ###

func TestCanMoveL(t *testing.T) {
	sh := shape{
		L, point{2, 5},
	}

	tests := []struct {
		ch      chamber
		blocker point
		d       jet
		e       bool
	}{
		{chamber{}, point{0, 0}, RIGHT, true},
		{chamber{}, point{0, 0}, LEFT, false},
		// Right
		{chamber{}, point{3, 5}, RIGHT, false},
		{chamber{}, point{3, 4}, RIGHT, false},
		{chamber{}, point{3, 3}, RIGHT, false},
		// Left
		{chamber{}, point{1, 5}, LEFT, false},
		{chamber{}, point{1, 4}, LEFT, false},
		{chamber{}, point{0, 3}, LEFT, false},
		// Down
		{chamber{}, point{2, 2}, DOWN, false},
		{chamber{}, point{1, 2}, DOWN, false},
		{chamber{}, point{0, 2}, DOWN, false},
	}

	for i, test := range tests {
		test.ch[test.blocker] = true
		if result := test.ch.canMove(&sh, test.d); result != test.e {
			t.Fatalf("%d | expected %v, got %v when moving %v", i, test.e, result, string(test.d))
		}
	}
}

// .#.
// ###
// .#.

func TestCanMovePlus(t *testing.T) {
	sh := shape{
		PLUS, point{2, 5},
	}

	tests := []struct {
		ch      chamber
		blocker point
		d       jet
		e       bool
	}{
		{chamber{}, point{0, 0}, RIGHT, true},
		{chamber{}, point{0, 0}, LEFT, true},
		{chamber{}, point{0, 0}, DOWN, true},
		// Right
		{chamber{}, point{3, 5}, RIGHT, false},
		{chamber{}, point{4, 4}, RIGHT, false},
		{chamber{}, point{3, 3}, RIGHT, false},
		// Left
		{chamber{}, point{1, 5}, LEFT, false},
		{chamber{}, point{0, 4}, LEFT, false},
		{chamber{}, point{1, 3}, LEFT, false},
		// Down
		{chamber{}, point{1, 3}, DOWN, false},
		{chamber{}, point{2, 2}, DOWN, false},
		{chamber{}, point{3, 3}, DOWN, false},
	}

	for i, test := range tests {
		test.ch[test.blocker] = true
		if result := test.ch.canMove(&sh, test.d); result != test.e {
			t.Fatalf("%d | expected %v, got %v when moving %v", i, test.e, result, string(test.d))
		}
	}
}

// ####

func TestCanMoveHoriz(t *testing.T) {
	sh := shape{
		HORIZ, point{2, 5},
	}

	tests := []struct {
		ch      chamber
		blocker point
		d       jet
		e       bool
	}{
		{chamber{}, point{0, 0}, RIGHT, true},
		{chamber{}, point{0, 0}, LEFT, true},
		{chamber{}, point{0, 0}, DOWN, true},
		// LEFT
		{chamber{}, point{1, 5}, LEFT, false},
		// RIGHT
		{chamber{}, point{6, 5}, RIGHT, false},
		// Down
		{chamber{}, point{2, 4}, DOWN, false},
		{chamber{}, point{3, 4}, DOWN, false},
		{chamber{}, point{4, 4}, DOWN, false},
		{chamber{}, point{5, 4}, DOWN, false},
		{chamber{}, point{6, 4}, DOWN, true},
		{chamber{}, point{1, 4}, DOWN, true},
	}

	for i, test := range tests {
		test.ch[test.blocker] = true
		if result := test.ch.canMove(&sh, test.d); result != test.e {
			t.Fatalf("%d | expected %v, got %v when moving %v", i, test.e, result, string(test.d))
		}
	}
}

func TestExamplesSeventeenOne(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 3068},
	}

	for _, test := range tests {
		result := run(test.test, 2022)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSeventeenOne(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		{reader, 3124},
	}

	for _, test := range tests {
		result := run(test.test, 2022)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func BenchmarkSeventeenOne(b *testing.B) {
	for x := 0; x <= b.N; x++ {
		b.StopTimer()
		file, _ := os.Open("./input")
		defer file.Close()
		reader := bufio.NewReader(file)

		tests := []struct {
			test     *bufio.Reader
			expected int
		}{
			{reader, 1514285714288},
		}
		b.StartTimer()
		result := run(tests[0].test, 1000000000000)
		b.StopTimer()
		if result != tests[0].expected {
			b.Fatalf("Result %d != expected %d", result, tests[0].expected)
		}
	}
}

func TestExamplesSeventeenTwo(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 1514285714288},
	}

	for _, test := range tests {
		if result := run(test.test, 1000000000000); test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSeventeenTwo(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		{reader, 1561176470569},
	}

	for _, test := range tests {
		if result := run(test.test, 1000000000000); result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
