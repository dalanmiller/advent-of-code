package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

const EXAMPLE = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestExamplesNineOne(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 13},
	}

	for _, test := range tests {
		result := run(test.test, 2)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestNineOne(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		// 6058 too high
		{reader, 6057},
	}

	for _, test := range tests {
		result := run(test.test, 2)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesNineTwo(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 1},
		{strings.NewReader(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`), 36},
	}

	for _, test := range tests {
		result := run(test.test, 10)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestNineTwo(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		{reader, 2514},
	}

	for _, test := range tests {
		result := run(test.test, 10)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
