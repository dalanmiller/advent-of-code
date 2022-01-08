package main

import (
	"log"
	"os"
	"testing"
)

func TestExamplesTwentyOneOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`Player 1 starting position: 4
Player 2 starting position: 8`, 739785},
	}

	for _, test := range tests {
		result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result %d != expected %d", result, test.expected)
		}
	}
}

func TestTwentyOneOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 518418},
	}

	for _, test := range tests {
		result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesTwentyOneTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`Player 1 starting position: 4
		Player 2 starting position: 8`,
			//11997614504960504
			//8746260974116208640
			444356092776315},
	}

	for _, test := range tests {
		result := run_two(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestTwentyOneTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 116741133558209},
	}

	for _, test := range tests {
		result := run_two(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
