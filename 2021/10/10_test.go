package main

import (
	"log"
	"os"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`{([(<{}[<>[]}>{[]{[(<()>`, 1197},
		{`[[<[([]))<([[{}[[()]]]`, 3},
		{`[{[{({}]{}}([{[{{{}}([]`, 57},
		{`[<(<(<(<{}))><([]([]()`, 3},
		{`<{([([[(<>()){}]>(<<{{`, 25137},
	}

	for i, test := range tests {
		result, _ := parseLine(test.test, false)
		if result != test.expected {
			log.Fatalf("%d - %s, incorrect got %d, expected %d", i, test.test, result, test.expected)
		}
	}
}
func TestExamplesTenOne(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`, 26397},
	}

	for _, test := range tests {
		result := run(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestTenOne(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 193275},
	}

	for _, test := range tests {
		result := run(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesTenTwo(t *testing.T) {
	tests := []struct {
		test     string
		expected int
	}{
		{`[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`, 288957},
	}

	for _, test := range tests {
		result := run_two(test.test)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestOneTwo(t *testing.T) {
	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file")
	}

	tests := []struct {
		test     string
		expected int
	}{
		{string(file), 2429644557},
	}

	for _, test := range tests {
		result := run_two(test.test)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
