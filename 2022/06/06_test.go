package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func TestExamplesSixOne(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader("bvwbjplbgvbhsrlpgdmjqwftvncz"), 5},
		{strings.NewReader("nppdvjthqldpwncqszvftbrmjlhg"), 6},
		{strings.NewReader("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 10},
		{strings.NewReader("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 11},
	}

	for _, test := range tests {
		result := run(test.test, 4)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSixOne(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		{reader, 1625},
	}

	for _, test := range tests {
		result := run(test.test, 4)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesSixTwo(t *testing.T) {
	tests := []struct {
		test     io.Reader
		expected int
	}{
		{strings.NewReader("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 19},
		{strings.NewReader("bvwbjplbgvbhsrlpgdmjqwftvncz"), 23},
		{strings.NewReader("nppdvjthqldpwncqszvftbrmjlhg"), 23},
		{strings.NewReader("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 29},
		{strings.NewReader("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 26},
	}

	for _, test := range tests {
		result := run(test.test, 14)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSixTwo(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		{reader, 2250},
	}

	for _, test := range tests {
		result := run(test.test, 14)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
