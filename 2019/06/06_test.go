package main

import (
	"log"
	"os"
	"testing"
)

func TestExamplesSixOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{"COM)B", 1},
		{"COM)B\nB)C\nC)D", 6},
		{"COM)B\nB)C\nB)D\nB)E\nB)F\nB)G", 11},
		{"COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L", 42},
	}

	for _, test := range tests {
		result, _ := run(test.test)
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
		{string(file), 122782},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestFind(t *testing.T) {

	p1 := Planet{
		Name: "COM",
	}

	p2 := Planet{
		Name: "B",
	}

	p3 := Planet{
		Name: "C",
	}

	p4 := Planet{
		Name: "D",
	}

	p5 := Planet{
		Name: "SAN",
	}

	p6 := Planet{
		Name: "YOU",
	}

	orbitalMap := make(map[*Planet]*Orbit)

	orbitalMap[&p1] = &Orbit{
		Planet: &p1,
		Orbits: []*Planet{
			&p2,
		},
	}

	orbitalMap[&p2] = &Orbit{
		Planet: &p2,
		Orbits: []*Planet{
			&p3,
		},
	}

	orbitalMap[&p3] = &Orbit{
		Planet: &p3,
		Orbits: []*Planet{
			&p4, &p5,
		},
	}

	orbitalMap[&p4] = &Orbit{
		Planet: &p4,
		Orbits: []*Planet{
			&p6,
		},
	}

	result := find("YOU", &p1, orbitalMap)

	len_result := len(result)
	if result[len_result-1] != &p6 {
		log.Fatalf("Did not find 'YOU', found %v", result[len_result-1])
	}

	if len_result != 5 {
		log.Fatalf("Incorrect result length, expected 4, got %d on %v", len_result, result)
	}
}

func TestExamplesSixTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{"COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN", 4},
	}

	for _, test := range tests {
		_, result := run(test.test)
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
		{string(file), 271},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

// func TestOneTwo(t *testing.T) {
// 	file, err := os.ReadFile("./input")
// 	if err != nil {
// 		log.Fatalf("could not read file")
// 	}

// 	tests := []struct {
// 		test     int
// 		expected int
// 	}{
// 		{0, 0},
// 	}

// 	for _, test := range tests {
// 		result := run(test.test)
// 		if result[0] != test.expected {
// 			t.Fatalf("Result % d != expected % d", result, test.expected)
// 		}
// 	}
// }
