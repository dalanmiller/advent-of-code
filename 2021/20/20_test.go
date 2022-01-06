package main

import (
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
				List:   &[]Pixel{tl, br},
				Lookup: map[Pixel]bool{tl: true, br: true},
			},
			Pixel{1, 1},
			257, // 100000001
		},
		{
			Image{
				List:   &[]Pixel{tr, bl},
				Lookup: map[Pixel]bool{tr: true, bl: true},
			},
			Pixel{1, 1},
			68, // 001000100
		},
		{
			Image{
				List:   &[]Pixel{ml, mr},
				Lookup: map[Pixel]bool{ml: true, mr: true},
			},
			Pixel{1, 1},
			40, // 000101000
		},
		{
			Image{
				List:   &[]Pixel{tm, bm},
				Lookup: map[Pixel]bool{tm: true, bm: true},
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

func TestString(t *testing.T) {
	tl := Pixel{0, 0}
	br := Pixel{2, 2}
	tests := []struct {
		image    Image
		expected string
	}{
		{
			Image{
				List:   &[]Pixel{tl, br},
				Lookup: map[Pixel]bool{tl: true, br: true},
				LeastX: 0,
				MaxX:   2,
				LeastY: 0,
				MaxY:   2,
			},
			`#..
...
..#
`,
		},
	}

	for _, test := range tests {
		if test.image.String() != test.expected {
			t.Fatalf(
				"Incorrect binary representation, got %s, expected %s",
				test.image.String(),
				test.expected,
			)
		}
	}
}

func TestExamplesTwentyPartOne(t *testing.T) {
	tests := []struct {
		test       string
		iterations int
		expected   int
	}{
		{`..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

		#..#.
		#....
		##..#
		..#..
		..###`, 0, 10},
		{`..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

		#..#.
		#....
		##..#
		..#..
		..###`, 1, 24},
		{`..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

		#..#.
		#....
		##..#
		..#..
		..###`, 2, 35},

		{`...............#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

		...
		.#.
		...`, 1, 2},
	}

	for _, test := range tests {
		result := run(test.test, test.iterations)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
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
		// !5278 actual <
		// !5073, actual >
		{input, 0, 4916},
		{input, 1, 5073},
		{input, 2, 5278},
	}

	for _, test := range tests {
		result := run(test.test, test.iterations)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

// func TestExamplesOneTwo(t *testing.T) {
// 	tests := []struct {
// 		test     string
// 		expected int
// 	}{
// 		{"", 0},
// 	}

// 	for _, test := range tests {
// 		result := run(test.test)
// 		if test.expected != result {
// 			t.Fatalf("Result % d != expected % d", result, test.expected)
// 		}
// 	}
// }

// func TestOneTwo(t *testing.T) {
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
// 		result := run(test.test)
// 		if result[0] != test.expected {
// 			t.Fatalf("Result % d != expected % d", result, test.expected)
// 		}
// 	}
// }
