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
		expected []int
	}{
		{`3,4,3`, []int{
			3,
			4,
			3,
		}},
	}

	for _, test := range tests {
		result := parseInput(test.test)
		if !reflect.DeepEqual(test.expected, result) {
			t.Fatalf("Result %v != expected %v", result, test.expected)
		}
	}
}

func TestExamplesSixOne(t *testing.T) {
	tests := []struct {
		test     string
		days     int
		expected int
	}{
		{`1`, 3, 2},
		{`3`, 5, 2},
		{`1,1,1,1,1`, 3, 10},
		{`1,1,1,1,1`, 12, 20},
		{`1,1,1,1,1`, 24, 45},
		{`3,4,3,1,2`, 80, 5934},
	}

	for _, test := range tests {
		result := run(test.test, test.days)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSixOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 385391},
	}

	for _, test := range tests {
		result := run(test.test, 80)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesSixTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`3,4,3,1,2`, 26984457539},
	}

	for _, test := range tests {
		result := run(test.test, 256)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSixTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 1728611055389},
	}

	for _, test := range tests {
		result := run(test.test, 256)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func BenchmarkSixTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {

		file, err := os.ReadFile("./input")
		if err != nil {
			log.Fatalf("could not read file")
		}

		tests := []struct {
			test     string
			expected int
		}{
			{string(file), 1728611055389},
		}

		for _, test := range tests {
			result := run(test.test, 256)
			if result != test.expected {
				b.Fatalf("Result % d != expected % d", result, test.expected)
			}
		}
	}
}
