package main

import (
	"log"
	"os"
	"testing"
)

func TestExamplesSixteenOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int64
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
		expected int64
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

func TestParsingSixteenTwo(t *testing.T) {
	tests := []struct {
		message          string
		expectedValue    int64
		expectedChildren int
	}{
		{
			"10011100000000010100000100001000000000100101000000110010000011110001100000000010000100000100101000001000",
			1,
			2,
		},
		{
			"10011100000000010100000100001000000000100101000000110010000011110001100000000010000100000100101000001100",
			0,
			2,
		},
		{ // Lesser than true
			"10011000000000010100000100001000000000100101000000110010000011110001100000000010000100000100101000001100",
			1,
			2,
		},
		{ // Greater than false
			"10010100000000010100000100001000000000100101000000110010000011110001100000000010000100000100101000001100",
			0,
			2,
		},
		{ // Greater than true
			"10010100000000010100000100001000000000100101000010110010000011110001100000000010000100000100101000001100",
			1,
			2,
		},
		{ // Sum 1, 2, 2, 2
			"11000010000000010011010000001010100000100101000001001010000010",
			7,
			4,
		},
		{ // Prod 1, 2, 2, 2
			"11000110000000010011010000001010100000100101000001001010000010",
			8,
			4,
		},
		{ // Prod 1, 2, 2, 10
			"11000110000000010011010000001010100000100101000001001010001010",
			40,
			4,
		},
		{ // Min 1, 2, 2, 10
			"11001010000000010011010000001010100000100101000001001010001010",
			1,
			4,
		},
		{ // Max 1, 2, 2, 10
			"11001110000000010011010000001010100000100101000001001010001010",
			10,
			4,
		},
		{ // Max 1, 2, 2, 10
			"1100111000000001001101000000101010000010010100000100101000101000000000000000000000000",
			10,
			4,
		},
		{ // Literal 15
			"00010001111",
			15,
			0,
		},
	}

	for _, test := range tests {
		p := Packet{
			Data: test.message,
		}
		p.parsePacket()
		if test.expectedValue != p.Value {
			log.Fatalf("Expected value %d, got %d", test.expectedValue, p.Value)
		}

		if test.expectedChildren != len(p.Children) {
			log.Fatalf("Expected number of children %d, has %d", test.expectedChildren, len(p.Children))
		}
	}
}

func TestExamplesSixteenTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int64
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

func TestSixteenTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int64
	}{
		{string(file), 1264857437203},
	}

	for _, test := range tests {
		result := run(test.test, true)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
