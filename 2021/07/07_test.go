package main

import (
	"log"
	"os"
	"testing"
)

func TestExamplesSevenOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{"16,1,2,0,4,2,7,1,2,14", 37},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSevenOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 348996},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestFuelDistanceCalculator(t *testing.T) {
	tests := []struct {
		distance int
		expected int
	}{
		{4, 10},
		{1, 1},
		{2, 3},
		{3, 6},
		{5, 15},
	}

	memo := make(map[int]int)
	for _, test := range tests {
		result := calculateFuelRequirement(test.distance, 1, memo)
		if result != test.expected {
			log.Fatalf("Distance calculation failed, got %d, expected %d", result, test.expected)
		}
	}
}
func TestExamplesSevenTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{"16,1,2,0,4,2,7,1,2,14", 168},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSevenTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 98231647},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func BenchmarkRun(b *testing.B) {
	file, _ := os.ReadFile("./input")

	for i := 0; i < b.N; i++ {
		_, result := run(string(file))
		result++
	}
}
