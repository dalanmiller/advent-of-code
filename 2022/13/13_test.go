package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

const EXAMPLE = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

func TestExampleThirteen(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{EXAMPLE, 13},
		{`[[4]]
[[5]]`, 1},
	}
	for _, test := range tests {
		result := runPartOne(strings.NewReader(test.test))
		if result != test.expected {
			log.Fatalf("Result %d != expected %d", result, test.expected)
		}
	}
}

func TestThirteenOne(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		// 3650 too low
		// 4442 too low
		// 6251 too high
		{reader, 6240},
	}

	for _, test := range tests {
		result := runPartOne(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesThirteenTwo(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE + "\n\n[[2]]\n[[6]]"), 140},
	}

	for _, test := range tests {
		result := runPartTwo(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestThirteenTwo(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()

	s, _ := io.ReadAll(file)
	reader := strings.NewReader(string(s) + "\n[[2]]\n[[6]]")

	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{reader, 23142},
	}

	for _, test := range tests {
		result := runPartTwo(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
