package main

import (
	"log"
	"os"
	"reflect"
	"testing"
)

// 1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2).
// 2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).
// 2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).
// 1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.

func TestExamplesTwoOne(t *testing.T) {
	tests := []struct {
		testBytes []byte
		expected  []int
	}{
		{[]byte("1,0,0,0,99"), []int{2, 0, 0, 0, 99}},
		{[]byte("2,3,0,3,99"), []int{2, 3, 0, 6, 99}},
		{[]byte("2,4,4,5,99,0"), []int{2, 4, 4, 5, 99, 9801}},
		{[]byte("1,1,1,4,99,5,6,0,99"), []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, test := range tests {
		result := run(test.testBytes)
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestTwoOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		testBytes []byte
		expected  int
	}{
		{file, 3409710},
	}

	for _, test := range tests {
		result := run(test.testBytes)
		if result[0] != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestTwoTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		testBytes []byte
		expected  int
	}{
		{file, 19690720},
	}

	for _, test := range tests {
		noun, verb := run_two(test.testBytes, test.expected)
		if noun == 0 && verb == 0 {
			log.Fatalf("%d, %d, %d", noun, verb, noun+verb*100)
		} else {
			log.Printf("Noun: %d, Verb: %d", noun, verb)
		}
	}
}
