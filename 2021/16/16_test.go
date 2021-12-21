package main

import (
	"log"
	"os"
	"testing"
)

func TestExamplesSixteenOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		// {"D2FE28", 2021},
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}

	for _, test := range tests {
		result := run(test.test, false)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestSixteenOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 875},
	}

	for _, test := range tests {
		result := run(test.test, false)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesSixteenTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	}

	for _, test := range tests {
		result := run(test.test, true)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

// func TestSixteenTwo(t *testing.T) {
// 	file, err := os.ReadFile("./input")
// 	if err != nil {
// 		log.Fatalf("could not read file")
// 	}

// 	tests := []struct {
// 		test     string
// 		expected int
// 	}{
// 		{string(file), 0},
// 	}

// 	for _, test := range tests {
// 		result := run(test.test, true)
// 		if result != test.expected {
// 			t.Fatalf("Result % d != expected % d", result, test.expected)
// 		}
// 	}
// }
