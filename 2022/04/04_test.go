package main

import (
	"log"
	"os"
	"testing"
)

const EXAMPLE = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestExamplesFourOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{EXAMPLE, 2},
		{"35-60,36-60", 1},
		{"6-94,94-94", 1},
		{"3-95,3-98", 1},
		{"1-1,1-1", 1},
		{"94-94,6-94", 1},
		{"3-98,3-95", 1},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestFourOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		// 491 too low
		{string(file), 550},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesFourTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{EXAMPLE, 4},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestFourTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		// 960 too high
		{string(file), 931},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
