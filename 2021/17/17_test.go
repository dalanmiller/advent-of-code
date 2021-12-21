package main

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		test     string
		expected *TargetBox
	}{
		{
			"target area: x=20..30, y=-10..-5",
			&TargetBox{
				XStart: 20,
				XEnd:   30,
				YStart: -5,
				YEnd:   -10,
			},
		},
		{
			"target area: x=50..150, y=-100..0",
			&TargetBox{
				XStart: 50,
				XEnd:   150,
				YStart: 0,
				YEnd:   -100,
			},
		},
	}

	for _, test := range tests {
		result := parseInput(test.test)
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %v != expected %v", result, test.expected)
		}
	}
}

func TestGenerateVelocities(t *testing.T) {
	tests := []struct {
		lowerX, upperX, lowerY, upperY int
		expected                       []VelocityEstimate
	}{
		{
			1, 2, -1, 3,
			[]VelocityEstimate{
				{1, -1, 0},
				{1, 0, 0},
				{1, 1, 0},
				{1, 2, 0},
				{1, 3, 0},
				{2, -1, 0},
				{2, 0, 0},
				{2, 1, 0},
				{2, 2, 0},
				{2, 3, 0},
			},
		},
	}

	for _, test := range tests {
		result := generatePossibleVelocities(test.lowerX, test.upperX, test.lowerY, test.upperY)
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result %v != expected %v", result, test.expected)
		}
	}
}

func TestExamplesSeventeenOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{
			"target area: x=20..30, y=-10..-5",
			45,
		},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSeventeenOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 3160},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesSeventeenTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{"target area: x=20..30, y=-10..-5", 112},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSeventeenTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 1928},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dalanmiller/aoc/2021/17
// BenchmarkSeventeenTwo-10    	    2311	    516646 ns/op	 1999894 B/op	      13 allocs/op
// PASS
// ok  	github.com/dalanmiller/aoc/2021/17	2.571s

func BenchmarkSeventeenTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, _ := os.ReadFile("./input")
		maxY, valids := run(string(file))
		if maxY != 3160 {
			b.Fatalf("Max Y fail, wanted %d, got %d", 3160, maxY)
		}

		if valids != 1928 {
			b.Fatalf("Valid solutions fail")
		}
	}
}
