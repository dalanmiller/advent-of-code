package main

import (
	"log"
	"os"
	"testing"
)

func TestAdjacent(t *testing.T) {
	tl := Pixel{0, 0}
	br := Pixel{2, 2}

	tr := Pixel{2, 0}
	bl := Pixel{0, 2}

	ml := Pixel{0, 1}
	mr := Pixel{2, 1}

	tm := Pixel{1, 0}
	bm := Pixel{1, 2}
	tests := []struct {
		image    Image
		pixel    Pixel
		expected int
	}{
		{
			Image{
				Known: map[Pixel]bool{tl: true, br: true},
			},
			Pixel{1, 1},
			257, // 100000001
		},
		{
			Image{
				Known: map[Pixel]bool{tr: true, bl: true},
			},
			Pixel{1, 1},
			68, // 001000100
		},
		{
			Image{
				Known: map[Pixel]bool{ml: true, mr: true},
			},
			Pixel{1, 1},
			40, // 000101000
		},
		{
			Image{
				Known: map[Pixel]bool{tm: true, bm: true},
			},
			Pixel{1, 1},
			130, // 010000010
		},
	}

	for _, test := range tests {
		result := test.image.adjacent(test.pixel)
		if result != test.expected {
			t.Fatalf(
				"Incorrect binary representation, got %d, expected %d",
				test.image.adjacent(test.pixel),
				test.expected,
			)
		}
	}
}

func TestTwentyPartOne(t *testing.T) {
	file, _ := os.ReadFile("./input")
	input := string(file)

	tests := []struct {
		test       string
		iterations int
		expected   int
	}{
		// <5278
		// >5073
		// !5167
		// !5400
		// !5020
		// +5081
		{input, 0, 4916},
		{input, 2, 5081},
	}

	for _, test := range tests {
		result := run(test.test, test.iterations)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func BenchmarkTwentyPartOne(b *testing.B) {
	file, _ := os.ReadFile("./input")
	input := string(file)

	for i := 0; i < b.N; i++ {
		result := run(input, 2)
		if result != 5171 {
			b.Fatalf("Result % d != expected % d", result, 5121)
		}
	}
}

func TestTwentyTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test       string
		iterations int
		expected   int
	}{
		{string(file), 50, 15088},
	}

	for _, test := range tests {
		result := run(test.test, test.iterations)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
