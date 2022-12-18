package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

const EXAMPLE = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestExamplesTwelveOne(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 31},
	}

	for _, test := range tests {
		result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestTwelveOne(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		// 6075, too high
		// 53, not right
		// 991, too high
		// 450, not right
		// 992, not right
		// 598, not right
		// 744, not right
		{reader, 440},
	}

	for _, test := range tests {
		result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesTwelveTwo(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 29},
	}

	for _, test := range tests {
		result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

// func TestTwelveTwo(t *testing.T) {
// 	file, _ := os.Open("./input")
// 	defer file.Close()
// 	reader := bufio.NewReader(file)

// 	tests := []struct {
// 		test     *bufio.Reader
// 		expected int
// 	}{
// 		{reader, 0},
// 	}

// 	for _, test := range tests {
// 		result := run(test.test)
// 		if result != test.expected {
// 			t.Fatalf("Result % d != expected % d", result, test.expected)
// 		}
// 	}
// }
