package main

import (
	"log"
	"os"
	"testing"
)

const ExampleOne = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestExamplesOneOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{ExampleOne, 24000},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestOneOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 72478},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesOneTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{ExampleOne, 45000},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestOneTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 210367},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
