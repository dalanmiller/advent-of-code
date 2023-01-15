package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

const EXAMPLE = `1
2
-3
3
-2
0
4`

func TestExamplesTwoOne(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 3},
		// 		{strings.NewReader(`1000
		// 2000
		// 3000`), 6000},
		// 		{strings.NewReader(`-8
		// 1
		// 2
		// 3`), 1},
	}

	for _, test := range tests {
		if result := run(test.test); test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestTwoOne(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		// 14320, too low
		// 3845, too low
		{reader, 14526},
	}

	for _, test := range tests {
		if result := run(test.test); result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesTwoTwo(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 1623178306},
	}

	for _, test := range tests {
		result := runPartTwo(test.test, 811589153)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestTwoTwo(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		{reader, 9738258246847},
	}

	for _, test := range tests {
		result := runPartTwo(test.test, 811589153)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
