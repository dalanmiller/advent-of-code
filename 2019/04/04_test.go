package main

import (
	"testing"
)

func TestExamplesFourOne(t *testing.T) {
	tests := []struct {
		r        Range
		expected int
	}{
		{Range{Begin: 111111, End: 111111}, 1},
		{Range{Begin: 223450, End: 223450}, 0},
		{Range{Begin: 123789, End: 123789}, 0},
	}

	for _, test := range tests {
		result := run(test.r, []Rule{
			DoubleAdjacentDigitRule{},
			LengthRule{length: 6},
			IncreasingRule{},
		})
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestFourOne(t *testing.T) {
	tests := []struct {
		r        Range
		expected int
	}{
		{Range{Begin: 145852, End: 616942}, 1767},
	}

	for _, test := range tests {
		result := run(test.r, []Rule{
			DoubleAdjacentDigitRule{},
			LengthRule{length: 6},
			IncreasingRule{},
		})
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesFourTwo(t *testing.T) {
	tests := []struct {
		r        Range
		expected int
	}{
		{Range{Begin: 112233, End: 112233}, 1},
		{Range{Begin: 112233, End: 112233}, 1},
		{Range{Begin: 123444, End: 123444}, 0},
		{Range{Begin: 111122, End: 111122}, 1},
		{Range{Begin: 112222, End: 112222}, 1},
		{Range{Begin: 123455, End: 123455}, 1},
		{Range{Begin: 999999, End: 999999}, 0},
	}

	for _, test := range tests {
		result := run(test.r, []Rule{
			DoubleAdjacentDigitRuleTwo{},
			LengthRule{length: 6},
			IncreasingRule{},
		})
		if result != test.expected {
			t.Fatalf("Result % d != expected % d, rule %d", result, test.expected, test.r.Begin)
		}
	}
}

func TestFourTwo(t *testing.T) {
	tests := []struct {
		r        Range
		expected int
	}{
		{Range{Begin: 145852, End: 616942}, 1192},
	}

	for _, test := range tests {
		result := run(test.r, []Rule{
			DoubleAdjacentDigitRuleTwo{},
			LengthRule{length: 6},
			IncreasingRule{},
		})
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
