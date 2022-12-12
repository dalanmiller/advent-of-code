package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

const EXAMPLE = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

func TestMonkeyOpParsing(t *testing.T) {

	monkeys := readInput(strings.NewReader(EXAMPLE))

	tests := []struct {
		monkeyN       int
		opValue       int
		expectedValue int
		errorMessage  string
	}{
		{0, 100, 1900, "Monkeys[0] op is wrong"},
		{1, 100, 106, "Monkeys[1] op is wrong"},
		{2, 100, 10000, "Monkeys[2] op is wrong"},
		{3, 100, 103, "Monkeys[3] op is wrong"},
	}

	for _, test := range tests {

		if monkeys[test.monkeyN].op(test.opValue) != test.expectedValue {
			t.Fatalf(test.errorMessage)
		}
	}
}

func TestMonkeyTestParsing(t *testing.T) {

	monkeys := readInput(strings.NewReader(EXAMPLE))

	tests := []struct {
		monkeyN       int
		opValue       int
		expectedValue int
		errorMessage  string
	}{
		{0, 23, 2, "Monkeys[0] test is wrong"},
		{0, 24, 3, "Monkeys[0] test is wrong"},
		{1, 19, 2, "Monkeys[1] test is wrong"},
		{1, 20, 0, "Monkeys[1] test is wrong"},
		{2, 13, 1, "Monkeys[2] test is wrong"},
		{2, 14, 3, "Monkeys[2] test is wrong"},
		{3, 17, 0, "Monkeys[3] test is wrong"},
		{3, 18, 1, "Monkeys[3] test is wrong"},
	}

	for _, test := range tests {

		if monkeys[test.monkeyN].test(test.opValue) != test.expectedValue {
			t.Fatalf(test.errorMessage)
		}
	}
}

func TestExamplesElevenOne(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 10605},
	}

	for _, test := range tests {
		result := run(test.test, 20, 1)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestElevenOne(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{
		{reader, 64032},
	}

	for _, test := range tests {
		result := run(test.test, 20, 1)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestExamplesElevenTwo(t *testing.T) {
	tests := []struct {
		test     *strings.Reader
		expected int
	}{
		{strings.NewReader(EXAMPLE), 2713310158},
	}

	for _, test := range tests {
		result := run(test.test, 10000, 2)
		if test.expected != result {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}

func TestElevenTwo(t *testing.T) {
	file, _ := os.Open("./input")
	defer file.Close()
	reader := bufio.NewReader(file)

	tests := []struct {
		test     *bufio.Reader
		expected int
	}{

		// 26669778282, too high
		{reader, 12729522272},
	}

	for _, test := range tests {
		result := run(test.test, 10000, 2)
		if result != test.expected {
			t.Fatalf("Result % d != expected % d", result, test.expected)
		}
	}
}
