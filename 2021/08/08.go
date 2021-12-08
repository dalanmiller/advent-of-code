package main

import (
	"sort"
	"strconv"
	"strings"
)

type Signal struct {
	Segments int
	Raw      string
	A        bool
	B        bool
	C        bool
	D        bool
	E        bool
	F        bool
	G        bool
	N        string
}

type Row struct {
	Signals []Signal
	Outputs []Signal
}

// Segments => Numbers
// 0 => 0
// 1 => 0
// 2 => 1
// 3 => 7
// 4 => 4
// 5 => 5, 3, 2
// 6 => 6, 0, 9
// 7 => 8

var segmentValueMap = map[int][]int{
	0: {0},
	1: {0},
	2: {1},
	3: {7},
	4: {4},
	5: {2, 3, 5},
	6: {0, 6, 9},
	7: {8},
}

// For each wire "A, B, C, ...", is incorrectly lighting up SEGMENT A, B, C.... etc.
// First layer: deduce based on digits with unique segment counts 1, 7, 4, 8
// Second layer: ???

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func parseSignalsString(signal string) Signal {
	runeSlice := []rune(signal)
	sort.Sort(sortRuneString(runeSlice))
	s := Signal{
		Segments: len(signal),
		Raw:      string(runeSlice),
	}
	for _, char := range signal {
		switch char {
		case 'a':
			s.A = true
		case 'b':
			s.B = true
		case 'c':
			s.C = true
		case 'd':
			s.D = true
		case 'e':
			s.E = true
		case 'f':
			s.F = true
		}
	}

	return s
}

func parseInput(input string) []Row {
	lines := strings.Split(input, "\n")

	var rows []Row
	for _, line := range lines {

		var signals, outputs []Signal
		signals_output := strings.Split(line, " | ")
		raw_signals := signals_output[0]
		raw_outputs := signals_output[1]

		signals_slice := strings.Split(raw_signals, " ")
		outputs_slice := strings.Split(raw_outputs, " ")

		for _, signal := range signals_slice {
			s := parseSignalsString(signal)
			signals = append(signals, s)
		}

		for _, output := range outputs_slice {
			o := parseSignalsString(output)
			outputs = append(outputs, o)
		}

		rows = append(rows, Row{
			Signals: signals,
			Outputs: outputs,
		})
	}

	return rows
}

func intersect(a Signal, b Signal) ([]rune, []rune) {
	// segments := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	seen_a := make(map[rune]bool)
	for _, chr := range a.Raw {
		seen_a[chr] = true
	}

	var intersection, difference []rune
	for _, chr := range b.Raw {
		_, ok := seen_a[chr]

		if ok {
			intersection = append(intersection, chr)
		} else {
			difference = append(difference, chr)
		}

		delete(seen_a, chr)
	}

	for k, _ := range seen_a {
		difference = append(difference, k)
	}

	return intersection, difference
}

func run(input string) int {
	rows := parseInput(input)

	var sum int
	for _, row := range rows {
		for _, output := range row.Outputs {
			lo := output.Segments
			if lo == 2 || lo == 3 || lo == 4 || lo == 7 {
				sum++
			}
		}
	}

	return sum
}

func run_two(input string) int {
	rows := parseInput(input)

	var sum int
	for _, row := range rows {
		signals := row.Signals
		outputs := row.Outputs

		var One, Two, Three, Four, Five, Six, Seven, Eight, Nine Signal
		var unknown []Signal
		for _, signal := range signals {
			switch signal.Segments {
			case 2:
				One = signal
				One.N = "1"
			case 3:
				Seven = signal
				Seven.N = "7"
			case 4:
				Four = signal
				Four.N = "4"
			case 7:
				Eight = signal
				Eight.N = "8"
			default:
				unknown = append(unknown, signal)
			}
		}

		for _, unk := range unknown {
			if unk.Segments == 5 {
				// 5, 3, or 2 henceforth
				_, d := intersect(unk, One)
				if len(d) == 5 {
					// 2 or 5
					_, d := intersect(unk, Four)
					if len(d) == 5 {
						//2
						Two = unk
						Two.N = "2"
					} else if len(d) == 2 {
						//5
						Five = unk
						Five.N = "5"
					}
				} else {
					Three = unk
					Three.N = "3"
				}

			} else {
				// 0, 6, or 9 henceforth
				_, d := intersect(unk, Eight)
				if len(d) == 1 {
					// 0 or 9 henceforth
					_, d := intersect(unk, One)
					if len(d) == 4 {
						// 9
						Nine = unk
						Nine.N = "9"
					} else {
						// 5
						Five = unk
						Five.N = "5"
					}

				} else {
					// 6
					Six = unk
					Six.N = "6"
				}
			}
		}

		signalMap := make(map[string]Signal)
		for _, signal := range []Signal{One, Two, Three, Four, Five, Six, Seven, Eight, Nine} {
			signalMap[signal.Raw] = signal
		}

		final_output := make([]string, len(outputs))
		for i, output := range outputs {
			final_output[i] = signalMap[output.Raw].N
		}

		value, _ := strconv.Atoi(strings.Join(final_output, ""))
		sum += value
	}

	return sum
}
