package main

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestParseinput(t *testing.T) {
	ul := &Point{
		N:    9,
		Sink: false,
	}
	um := &Point{
		N:    8,
		Sink: false,
	}
	ur := &Point{
		N:    9,
		Sink: false,
	}

	ml := &Point{
		N:    9,
		Sink: false,
	}
	mm := &Point{
		N:    1,
		Sink: true,
	}
	mr := &Point{
		N:    9,
		Sink: false,
	}

	bl := &Point{
		N:    9,
		Sink: false,
	}
	bm := &Point{
		N:    9,
		Sink: false,
	}
	br := &Point{
		N:    9,
		Sink: false,
	}

	ul.Right = um
	ul.Down = ml
	ul.NextLowest = um

	ur.Left = um
	ur.Down = mr
	ur.NextLowest = um

	um.Right = ur
	um.Left = ul
	um.Down = mm
	um.NextLowest = mm
	um.FinalSink = mm

	ml.Up = ul
	ml.Right = mm
	ml.Down = bl
	ml.NextLowest = mm

	mm.Up = um
	mm.Right = mr
	mm.Down = bm
	mm.Left = ml

	mr.Up = ur
	mr.Left = mm
	mr.Down = br
	mr.NextLowest = mm

	bl.Up = ml
	bl.Right = bm

	bm.Left = bl
	bm.Up = mm
	bm.Right = br
	bm.NextLowest = mm

	br.Left = bm
	br.Up = mr

	tests := []struct {
		test     string
		expected [][]*Point
	}{
		{
			`989
919
999`,
			[][]*Point{
				{ul, um, ur},
				{ml, mm, mr},
				{bl, bm, br},
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

func TestExamplesNineOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`2199943210
3987894921
9856789892
8767896789
9899965678`, 15},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestNineOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 554},
	}

	for _, test := range tests {
		result, _ := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesNineTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`2199943210
3987894921
9856789892
8767896789
9899965678`, 1134},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestNineTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 1017792},
	}

	for _, test := range tests {
		_, result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
