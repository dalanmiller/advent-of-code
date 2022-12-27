package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

const EXAMPLE = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

func TestWithin(t *testing.T) {
	sensor := sensor{
		X:         0,
		Y:         0,
		Manhattan: float64(2),
	}

	tests := []struct {
		Point    point
		Expected bool
	}{
		{point{0, 3}, false},
		{point{0, 2}, true},
		{point{0, 0}, true},
		{point{-1, -1}, true},
		{point{2, 0}, true},
		{point{0, -3}, false},
		{point{3, 3}, false},
	}

	for _, test := range tests {
		result := sensor.Within(test.Point)
		if result != test.Expected {
			log.Fatalf("Incorrectly determined within, got %v and expected %v on point %v", result, test.Expected, test.Point)
		}
	}
}

func TestExamplesFifteenOne(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 26},
	}

	for _, test := range tests {
		result := runPartOne(test.test, 10)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestFifteenOne(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		// 4055738 too low
		{reader, 5108096},
	}

	for _, test := range tests {
		result := runPartOne(test.test, 2000000)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesFifteenTwo(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 56000011},
	}

	for _, test := range tests {
		result := runPartTwo(test.test, 0, 20)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestFifteenTwo(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		{reader, 10553942650264},
	}

	for _, test := range tests {
		result := runPartTwo(test.test, 0, 4000000)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
