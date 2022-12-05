package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

const EXAMPLE = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestExamplesFiveOne(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected string
	}{
		{strings.NewReader(EXAMPLE), "CMZ"},
	}

	for _, test := range tests {
		result := run(test.test, 9000)
		if test.expected != result {
			t.Fatalf("Result %s != expected %s", result, test.expected)
		}
	}
}

func TestFiveOne(t *testing.T) {
	file, _ := os.Open("./input")
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected string
	}{
		{reader, "PTWLTDSJV"},
	}

	for _, test := range tests {
		result := run(test.test, 9000)
		if result != test.expected {
			t.Fatalf("Result %s != expected %s", result, test.expected)
		}
	}
}

func TestExamplesFiveTwo(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected string
	}{
		{strings.NewReader(EXAMPLE), "MCD"},
	}

	for _, test := range tests {
		result := run(test.test, 9001)
		if test.expected != result {
			t.Fatalf("Result %s != expected %s", result, test.expected)
		}
	}
}

func TestFiveTwo(t *testing.T) {
	file, _ := os.Open("./input")
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected string
	}{
		{reader, "WZMFVGGZP"},
	}

	for _, test := range tests {
		result := run(test.test, 9001)
		if result != test.expected {
			t.Fatalf("Result %s != expected %s", result, test.expected)
		}
	}
}
