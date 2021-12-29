package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		test     string
		expected []*Scanner
	}{
		{`--- scanner 1 ---
1,2,3

--- scanner 11 ---
3,2,1

--- scanner 111 ---
-1,-2,-10
`, []*Scanner{
			{N: 1, Beacons: []Beacon{
				{X: 1, Y: 2, Z: 3},
			}},
			{N: 11, Beacons: []Beacon{
				{X: 3, Y: 2, Z: 1},
			}},
			{N: 111, Beacons: []Beacon{
				{X: -1, Y: -2, Z: -10},
			}},
		}},
	}

	for _, test := range tests {
		result := parseInput(test.test)
		if !reflect.DeepEqual(test.expected, result) {
			t.Fatalf("Result %v != expected %v", result, test.expected)
		}
	}
}

// func TestIntersection(t *testing.T) {

// 	a := DistanceSet{
// 		storage:     []float64{1, 2, 3},
// 		contentHash: map[float64]bool{1: true, 2: true, 3: true},
// 	}

// 	b := DistanceSet{
// 		storage:     []float64{3, 4, 5},
// 		contentHash: map[float64]bool{3: true, 4: true, 5: true},
// 	}

// 	tests := []struct {
// 		setA     *DistanceSet
// 		setB     *DistanceSet
// 		expected int
// 	}{
// 		{&a, &b, 1},
// 	}

// 	for _, test := range tests {
// 		result := test.setA.Intersection(test.setB)
// 		if test.expected != result {
// 			t.Fatalf("Set intersection failed, got %d, expected %d matches", result, test.expected)
// 		}
// 	}
// }

// func TestBeaconRotatations(t *testing.T) {
// 	// Test Y Axis
// 	tests := []struct {
// 		test     *Beacon
// 		expected Beacon
// 	}{
// 		{&Beacon{X: 1, Y: 1, Z: 1}, Beacon{X: 1, Y: 1, Z: -1}},
// 		{&Beacon{X: 1, Y: 1, Z: -1}, Beacon{X: -1, Y: 1, Z: -1}},
// 		{&Beacon{X: -1, Y: 1, Z: -1}, Beacon{X: -1, Y: 1, Z: 1}},
// 		{&Beacon{X: -1, Y: 1, Z: 1}, Beacon{X: 1, Y: 1, Z: 1}},
// 	}

// 	for _, test := range tests {
// 		result := test.test.rotate(YAxis)
// 		if !reflect.DeepEqual(test.expected, result) {
// 			t.Fatalf("RotateY failed, %v != expected %v", result, test.expected)
// 		}
// 	}

// 	// Test X Axis
// 	tests = []struct {
// 		test     *Beacon
// 		expected Beacon
// 	}{
// 		{&Beacon{X: 1, Y: 1, Z: 1}, Beacon{X: 1, Y: -1, Z: 1}},
// 		{&Beacon{X: 1, Y: -1, Z: 1}, Beacon{X: 1, Y: -1, Z: -1}},
// 		{&Beacon{X: 1, Y: -1, Z: -1}, Beacon{X: 1, Y: 1, Z: -1}},
// 		{&Beacon{X: 1, Y: 1, Z: -1}, Beacon{X: 1, Y: 1, Z: 1}},
// 	}

// 	for _, test := range tests {
// 		result := test.test.rotate(XAxis)
// 		if !reflect.DeepEqual(test.expected, result) {
// 			t.Fatalf("RotateX failed, %v != expected %v", result, test.expected)
// 		}
// 	}

// 	// Test Z Axis
// 	tests = []struct {
// 		test     *Beacon
// 		expected Beacon
// 	}{
// 		{&Beacon{X: 1, Y: 1, Z: 1}, Beacon{X: 1, Y: -1, Z: 1}},
// 		{&Beacon{X: 1, Y: -1, Z: 1}, Beacon{X: -1, Y: -1, Z: 1}},
// 		{&Beacon{X: -1, Y: -1, Z: 1}, Beacon{X: -1, Y: 1, Z: 1}},
// 		{&Beacon{X: -1, Y: 1, Z: 1}, Beacon{X: 1, Y: 1, Z: 1}},
// 	}

// 	for _, test := range tests {
// 		result := test.test.rotate(ZAxis)
// 		if !reflect.DeepEqual(test.expected, result) {
// 			t.Fatalf("RotateZ failed, %v != expected %v", result, test.expected)
// 		}
// 	}
// }

// func TestScannerRotatations(t *testing.T) {
// 	// Test Y Axis
// 	tests := []struct {
// 		test     Scanner
// 		expected Scanner
// 	}{
// 		{
// 			Scanner{1, []Beacon{{1, 1, 1, DistanceSet{}}}},
// 			Scanner{1, []Beacon{{1, -1, 1, DistanceSet{}}}},
// 		},
// 		{
// 			Scanner{1, []Beacon{{1, 1, 1, DistanceSet{}}}},
// 			Scanner{1, []Beacon{{1, -1, 1, DistanceSet{}}}},
// 		},
// 		{
// 			Scanner{1, []Beacon{{1, 1, 1, DistanceSet{}}}},
// 			Scanner{1, []Beacon{{1, -1, 1, DistanceSet{}}}},
// 		},
// 		{
// 			Scanner{1, []Beacon{{1, 1, 1, DistanceSet{}}}},
// 			Scanner{1, []Beacon{{1, -1, 1, DistanceSet{}}}},
// 		},
// 	}

// 	for _, test := range tests {
// 		result := test.test.rotate(XAxis)
// 		if !reflect.DeepEqual(test.expected, result) {
// 			t.Fatalf("Scanner rotation failed, %v != expected %v", result, test.expected)
// 		}
// 	}
// }

// func TestOrientationGenerator(t *testing.T) {
// 	scanners := parseInput(`--- scanner 0 ---
// 1,1,1`)

// 	if !reflect.DeepEqual(scanners[0], &Scanner{N: 0, Beacons: []Beacon{{1, 1, 1}}}) {
// 		t.Fatalf("Scanner parsed incorrectly")
// 	}

// 	orientations := scanners[0].generateOrientations()

// 	if len(orientations) != 24 {
// 		t.Fatalf("Did not generate 24 orientations, got %d", len(orientations))
// 	}
// 	if !reflect.DeepEqual(
// 		orientations,
// 		[]Scanner{
// 			// First four
// 			{0, []Beacon{{1, 1, 1}}},
// 			{0, []Beacon{{1, -1, 1}}},
// 			{0, []Beacon{{1, -1, -1}}},
// 			{0, []Beacon{{1, 1, -1}}},

// 			// Rotate Z Axis, next four
// 			{0, []Beacon{{1, -1, -1}}},
// 			{0, []Beacon{{-1, -1, -1}}},
// 			{0, []Beacon{{-1, -1, 1}}},
// 			{0, []Beacon{{1, -1, 1}}},

// 			// Rotate Z Axis, next four
// 			{0, []Beacon{{-1, -1, 1}}},
// 			{0, []Beacon{{-1, -1, -1}}},
// 			{0, []Beacon{{-1, 1, -1}}},
// 			{0, []Beacon{{-1, 1, 1}}},

// 			// Rotate Z Axis, next four
// 			{0, []Beacon{{1, 1, 1}}},
// 			{0, []Beacon{{1, 1, -1}}},
// 			{0, []Beacon{{-1, 1, -1}}},
// 			{0, []Beacon{{-1, 1, 1}}},

// 			// Rotate X Axis (facing up), next four
// 			{0, []Beacon{{-1, -1, 1}}},
// 			{0, []Beacon{{-1, 1, 1}}},
// 			{0, []Beacon{{1, 1, 1}}},
// 			{0, []Beacon{{1, -1, 1}}},

// 			// Rotate X Axis twice (facing down), final four
// 			{0, []Beacon{{1, 1, -1}}},
// 			{0, []Beacon{{1, -1, -1}}},
// 			{0, []Beacon{{-1, -1, -1}}},
// 			{0, []Beacon{{-1, 1, -1}}},
// 		},
// 	) {
// 		t.Fatalf("Did not generate correct orientations")
// 	}
// }

// func TestIntersectingBeacons(t *testing.T) {
// 	tests := []struct {
// 		a        Scanner
// 		b        Scanner
// 		expected int
// 	}{
// 		{
// 			Scanner{0, []Beacon{
// 				Beacon{X: 1, Y: 1, Z: 1},
// 				Beacon{X: 1, Y: 2, Z: 3},
// 				Beacon{X: 3, Y: 2, Z: 1},
// 			}},
// 			Scanner{1, []Beacon{
// 				{1, 0, 1, DistanceSet{}},
// 				{1, 2, 3, DistanceSet{}},
// 				{3, 1, 1, DistanceSet{}},
// 			}},
// 			1,
// 		},
// 	}

// 	for _, test := range tests {
// 		n := intersectBeacons(test.a, test.b)
// 		if n != test.expected {
// 			t.Fatalf("N of intersecting values %d, did not match expected %d", n, test.expected)
// 		}
// 	}
// }

func TestExamplesNineteenOne(t *testing.T) {
	input, _ := os.ReadFile("test_1")

	tests := []struct {
		test     string
		expected int
	}{
		{string(input), 79},
	}

	for _, test := range tests {
		result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestNineteenOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		t.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 0},
	}

	for _, test := range tests {
		result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

// func TestExamplesNineteenTwo(t *testing.T) {
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

// func TestNineteenTwo(t *testing.T) {
// 	file, err := os.ReadFile("./input")
// 	if err != nil {
// 		t.Fatalf("could not read file")
// 	}

// 	tests := []struct {
// 		test string
// 		expected  int
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
